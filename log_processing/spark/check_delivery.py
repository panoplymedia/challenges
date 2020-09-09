import argparse
import logging
import os
import hashlib
from pyspark.sql import SparkSession
from pyspark.sql import functions as f

PARQUET_PATH = '/data/parquet'

def parse_args():
  """
  Parse the command line arguments and return them.
  """
  parser = argparse.ArgumentParser(description='Query delivery data')

  parser.add_argument('--parquet_path',
                      type=str,
                      default=PARQUET_PATH,
                      help="The path from which to read the parquet data")

  parser.add_argument('--ip',
                      type=str,
                      help="The ip address to check")

  parser.add_argument('--ua',
                      type=str,
                      help="The user agent string to check")

  parser.add_argument('--asset',
                      type=str,
                      help="The asset to check")

  parser.add_argument('--start_byte_range',
                      type=int,
                      help="The starting byte range to check")

  parser.add_argument('--end_byte_range',
                      type=int,
                      help="The ending byte range to check")

  return parser.parse_args()

def query_delivery_parquet(spark, parquet_path, ip, ua, asset, start_range, end_range):
  """
  Queries the delivery parquet data for a given ip/ua/asset and returns the
  resulting DataFrame

  :param spark: The current active SparkSession
  :type spark: SparkSession
  :param parquet_path: The path from which to read the parquet data.
  :type parquet_path: str
  :param ip: The ip address to check
  :type ip: str
  :param ua: The user_agent string to check
  :type ua: str
  :param asset: The asset to check
  :type asset: str
  :param start_range: The starting byte range
  :type start_range: str
  :param end_range: The ending byte range
  :type end_range: str
  """
  # Convert the user_agent to md5 for lookup
  md5_ua = hashlib.md5(ua.encode('utf-8')).hexdigest()

  # Read in the parquet data
  pq_path = os.path.join(parquet_path, 'delivery.parquet')
  df = spark.read.parquet(pq_path)

  # Filter to the provided critera
  df = df\
    .select(
      f.col('ip'),
      f.col('user_agent_md5'),
      f.col('request'),
      f.col('byte_num')
    )\
    .filter(
      (f.col('ip') == ip) &
      (f.col('user_agent_md5') == md5_ua) &
      (f.col('request') == asset) &
      (f.col('byte_num').between(start_range, end_range))
    )

  return df

def is_delivered(df, start_range, end_range):
  """
  Checks if the results indicate that the asset was delivered successfully

  :param df: The delivery DataFrame
  :type df: DataFrame
  :param start_range: The starting byte range
  :type start_range: str
  :param end_range: The ending byte range
  :type end_range: str
  """
  expected_bytes = (end_range - start_range) + 1

  return df.distinct().count() == expected_bytes

if __name__ == '__main__':
  args = parse_args()

  logging.basicConfig(level=logging.INFO,
                      format='%(asctime)s %(levelname)s - %(message)s')

  # Setup our spark session in local mode
  spark = SparkSession.builder\
      .master("local[*]")\
      .appName("check_delivery")\
      .getOrCreate()

  df = query_delivery_parquet(spark=spark,
                              parquet_path=args.parquet_path,
                              ip=args.ip,
                              ua=args.ua,
                              asset=args.asset,
                              start_range=args.start_byte_range,
                              end_range=args.end_byte_range)

  delivered = is_delivered(df, args.start_byte_range, args.end_byte_range)

  message = [
    "\n[Delivery Report]",
    f"IP: {args.ip}",
    f"UA: {args.ua}",
    f"Asset: {args.asset}",
    f"Byte Range: {args.start_byte_range}-{args.end_byte_range}",
    f"\nDelivered: {delivered}"
  ]

  print("\n".join(message))
