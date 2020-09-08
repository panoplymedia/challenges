import pytest
import os
from pyspark.sql import functions as f
from pyspark.sql import Row
from pyspark.sql.types import (DateType, IntegerType, LongType, StringType,
                               StructField, StructType)
from .. import process_log as pl


@pytest.fixture(scope="function")
def raw_delivery_log_file(staging_path):
  """
  Writes out a static delivery log file for use in testing.
  """
  content = '\n'.join((
    "#\tTHIS\tIS\tA\tHEADER,\tIGNORE\tME!",
    "2020-07-31\t16:46:07\t63.110.194.22\tMy Silly User Agent String\t/6d8a9754-e8c7-4193-8491-58b2122c1c10\t206\t0-5",
    "2020-07-31\t12:46:09\t95.142.72.8\tAnother User Agent String\t/734f46f2-4e54-4f1a-80b4-15abb8220eaf\t206\t0-5",
    "2020-08-01\t12:00:00\t95.142.72.9\tInvalid Status\t/0bf0f036-f1f9-11ea-adc1-0242ac120002\t500\t0-5",
    "2020-08-01\t12:00:00\t95.142.72.10\tAnother Invalid Status\t/034dc480-f1fa-11ea-adc1-0242ac120002\t999\t0-5"
  ))

  file_path = staging_path.join("delivery.log")
  file_path.write(content)

  return file_path


def test_read_delivery_log(spark_session, raw_delivery_log_file):
  """
  Should be able to read the log file, and return with a valid dataframe
  """
  # Read the raw delivery log
  df = pl.read_delivery_log(spark=spark_session,
                            logfile_path=raw_delivery_log_file.strpath)

  # Assert that we get the expected results
  assert(df.count() == 4)
  assert(df.columns == ['date', 'timestamp', 'ip', 'user_agent', 'request', 'status', 'byte_range'])


test_filter_invalid_status_cases = {
    'Internal Server Error': (500, 0),
    'OK': (200, 1),
    'Partial Content': (206, 1)
}
@pytest.mark.parametrize('status, expected_rows',
                         list(test_filter_invalid_status_cases.values()),
                         ids=list(test_filter_invalid_status_cases.keys()))
def test_filter_invalid_status(spark_session, status, expected_rows):
  """
  Should filter out any records with a status that is not valid
  """
  # Create a DataFrame with our parameterized test data
  df = spark_session.createDataFrame([Row(status=status)])

  # Filter the DF
  filtered_df = pl.filter_invalid_status(df)

  # Verify we got the expected results
  assert(filtered_df.count() == expected_rows)


test_build_epoch_cases = {
    'Past': ('2020-07-31', '16:46:07', 1596213967),
    'Future': ('3000-10-25', '10:00:00', 32529376800)
}
@pytest.mark.parametrize('date, timestamp, expected',
                         list(test_build_epoch_cases.values()),
                         ids=list(test_build_epoch_cases.keys()))
def test_build_epoch(spark_session, date, timestamp, expected):
  """
  Should add an 'epoch' column to the dataframe with the expected value
  """
  # Create a DataFrame with our parameterized test data
  df = spark_session.createDataFrame([Row(date=date, timestamp=timestamp)])

  # Add the epoch column
  epoch_df = pl.build_epoch(df)

  # Verify new column exists
  assert(epoch_df.columns == ['date', 'timestamp', 'epoch'])

  # Verify we got the expected results
  actual_epoch = epoch_df\
    .select(f.col('epoch'))\
    .collect()[0]['epoch']
  assert(actual_epoch == expected)


test_hash_useragent_cases = {
    'iPhone': (
      'Mozilla/5.0 (Apple-iPhone7C2/1202.466; U; CPU like Mac OS X; en) AppleWebKit/420+ (KHTML, like Gecko) Version/3.0 Mobile/1A543 Safari/419.3',
      'be6febd99e8116e1dc3f0c95611017a4'
    ),
    'Linux': (
      'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:15.0) Gecko/20100101 Firefox/15.0.1',
      '6aa47caae0c78fd4c712e72d2fcd6d7f'
    )
}
@pytest.mark.parametrize('user_agent, expected',
                         list(test_hash_useragent_cases.values()),
                         ids=list(test_hash_useragent_cases.keys()))
def test_hash_useragent(spark_session, user_agent, expected):
  """
  Should add a 'user_agent_md5' column to the dataframe with the expected value
  """
   # Create a DataFrame with our parameterized test data
  df = spark_session.createDataFrame([Row(user_agent=user_agent)])

  # Add the user_agent_md5 column
  uahash_df = pl.hash_useragent(df)
  actual_uahash = uahash_df\
    .select(f.col('user_agent_md5'))\
    .collect()[0]['user_agent_md5']

  assert(actual_uahash == expected)


test_explode_bytes_cases = {
    'Small Range': ('0-1', 2),
    'Large Range': ('0-1000', 1001),
    'Range Starting In Middle': ('300-310', 11)
}
@pytest.mark.parametrize('byte_range, expected_rows',
                         list(test_explode_bytes_cases.values()),
                         ids=list(test_explode_bytes_cases.keys()))
def test_explode_bytes(spark_session, byte_range, expected_rows):
  """
  Should explode the byte_range so each byte becomes its own row
  """
  # Create a DataFrame with our parameterized test data
  df = spark_session.createDataFrame([Row(byte_range=byte_range)])

  # Explode the byte range
  exploded_df = pl.explode_bytes(df)

  # Verify new column exists
  assert(exploded_df.columns == ['byte_range', 'byte_num'])

  # Verify we got the expected results
  exploded_df.show(10, truncate=False)
  assert(exploded_df.count() == expected_rows)


def test_write_delivery(spark_session, parquet_path):
  """
  Should write the input dataframe to the expected parquet path.
  """
  # Create a DataFrame with our test data
  data = [
    Row(date='2020-07-31', epoch=1596213967, ip='63.110.194.22', user_agent_md5='my_user1', request='file1', byte_num=1),
    Row(date='2020-07-31', epoch=1596213967, ip='63.110.194.22', user_agent_md5='my_user1', request='file1', byte_num=2),
    Row(date='2020-07-31', epoch=1596213999, ip='63.110.194.10', user_agent_md5='my_user2', request='file2', byte_num=0),
  ]
  input_df = spark_session.createDataFrame(data)

  # Write the clean pagecounts to parquet
  pl.write_delivery_parquet(spark=spark_session,
                            parquet_path=parquet_path,
                            df=input_df)

  # Read the written parquet
  pq_file = os.path.join(parquet_path, 'delivery.parquet')
  written_df = spark_session.read.parquet(pq_file)

  # The data written to parquet should match what we submitted
  assert(written_df.count() == input_df.count())
  assert(set(written_df.columns) == set(input_df.columns))
