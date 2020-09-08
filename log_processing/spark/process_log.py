import argparse
import logging
import os
from pyspark.sql import SparkSession
from pyspark.sql import functions as f
from pyspark.sql.types import (ArrayType, DateType, IntegerType, LongType,
                               StringType, StructType, StructField)

PARQUET_PATH = '/data/parquet'

VALID_STATUSES = [200, 206]

def parse_args():
  """
  Parse the command line arguments and return them.
  """
  parser = argparse.ArgumentParser(description='Process delivery logs')

  parser.add_argument('log_file',
                      type=str,
                      metavar='log_file',
                      help="The delivery log file to process")

  parser.add_argument('--parquet_path',
                      type=str,
                      default=PARQUET_PATH,
                      help="The path in which to write the parquet data")

  return parser.parse_args()

def read_delivery_log(spark, logfile_path):
  """
  Reads the delivery log data from the provided path, splits the data into
  columns, applies a schema, and returns a DataFrame

  :param spark: The current active SparkSession
  :type spark: SparkSession
  :param logfile_path: The path from which to read log data
  :type logfile_path: str
  """
  logging.info('Reading raw log data from: %s', logfile_path)

  # The schema to enforce for incoming raw log files
  raw_logfile_schema = StructType([
    StructField('date', StringType(), True),
    StructField('timestamp', StringType(), True),
    StructField('ip', StringType(), True),
    StructField('user_agent', StringType(), True),
    StructField('request', StringType(), True),
    StructField('status', IntegerType(), True),
    StructField('byte_range', StringType(), True)
  ])

  df = spark.read.csv(path=logfile_path,
                      schema=raw_logfile_schema,
                      sep='\t', # Assuming the file will always be tab delemeted
                      comment='#') # Removes header line if it exists

  return df

def filter_invalid_status(df):
  """
  Filters a raw logfile DataFrame to only rows with valid status

  :param df: A logfile dataframe
  :type df: Dataframe
  """
  return df.where(df['status'].isin(VALID_STATUSES))

def build_epoch(df):
  """
  Adds a 'epoch' column by combining the 'date' and 'timestamp' columns and
  converting them into a unix epoch.

  :param df: A logfile dataframe
  :type df: Dataframe
  """
  df = df.withColumn(
    'epoch',
    f.unix_timestamp(f.concat_ws(' ', f.col('date'), f.col('timestamp')))
  )

  return df

def hash_useragent(df):
  """
  Adds a 'user_agent_md5' column containing the md5 hash of the useragent

  :param df: A logfile dataframe
  :type df: Dataframe
  """
  return df.withColumn('user_agent_md5', f.md5(f.col('user_agent')))

def explode_bytes(df):
  """
  Adds a 'byte_num' column, explodes the range of bytes into rows

  :param df: A logfile dataframe
  :type df: Dataframe
  """
  # UDF to generate array from a string range
  # Note the '+1' on the end range is required so we get the full range of bytes
  str_to_rng = f.udf(lambda srng: list(range(srng[0], srng[1]+1)),
                     ArrayType(IntegerType()))

  return df.withColumn(
    'byte_num',
    # Explodes the bytes so we end up with 1 row per byte
    f.explode(
      str_to_rng(
        f.array(
          f.split(f.col('byte_range'), '-')[0].cast('integer'),
          f.split(f.col('byte_range'), '-')[1].cast('integer')
        )
      )
    ),
  )

def write_delivery_parquet(spark, parquet_path, df):
  """
  Writes a delivery dataframe out to a parquet data store

  :param spark: The current active SparkSession
  :type spark: SparkSession
  :param df: A delivery dataframe
  :type df: Dataframe
  :param parquet_path: The path in which to write the parquet data.
  :type parquet_path: str
  """
  dest_file = os.path.join(parquet_path, 'delivery.parquet')

  logging.info('Writing delivery data to parquet: %s', dest_file)

  # Only write out the columns we actually need
  df = df.select(
    f.col('date').cast(DateType()),
    f.col('epoch').cast(LongType()),
    f.col('ip').cast(StringType()),
    f.col('user_agent_md5').cast(StringType()),
    f.col('request').cast(StringType()),
    f.col('byte_num').cast(IntegerType())
  )

  df.write.parquet(dest_file,
                   mode='append',
                   partitionBy='date',
                   compression='snappy')

if __name__ == '__main__':
  args = parse_args()

  logging.basicConfig(level=logging.INFO,
                      format='%(asctime)s %(levelname)s - %(message)s')

  # Setup out spark session in local mode
  spark = SparkSession.builder\
      .master("local[*]")\
      .appName("process_logs")\
      .getOrCreate()

  # Read in the raw log file
  delivery_df = read_delivery_log(spark, args.log_file)

  # Filter out invalid states
  delivery_df = filter_invalid_status(delivery_df)

  # Convert the user_agent to a hash to cut down on storage
  delivery_df = hash_useragent(delivery_df)

  # Add an epoc column to the dataframe
  delivery_df = build_epoch(delivery_df)

  # Explode the dataframe on the bytes_range of each row
  delivery_df = explode_bytes(delivery_df)

  # Persist the transformed data
  write_delivery_parquet(spark, args.parquet_path, delivery_df)
