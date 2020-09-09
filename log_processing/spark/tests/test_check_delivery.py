import pytest
from pyspark.sql import Row
from .. import check_delivery as cd

@pytest.fixture(scope="function")
def delivery_parquet(spark_session, parquet_path):
  """
  Writes out a static delivery parquet file for use in testing
  """
  data = [
    Row(date='2020-07-31', epoch=1596213967, ip='63.110.194.22', user_agent_md5='f7f4c471f44c4ed9070606fee66c8f58', request='file1', byte_num=0),
    Row(date='2020-07-31', epoch=1596213967, ip='63.110.194.22', user_agent_md5='f7f4c471f44c4ed9070606fee66c8f58', request='file1', byte_num=1),
    Row(date='2020-07-31', epoch=1596213967, ip='63.110.194.22', user_agent_md5='f7f4c471f44c4ed9070606fee66c8f58', request='file1', byte_num=2),
    Row(date='2020-07-31', epoch=1596213968, ip='63.110.194.22', user_agent_md5='f7f4c471f44c4ed9070606fee66c8f58', request='file1', byte_num=2),
    Row(date='2020-07-31', epoch=1596213968, ip='63.110.194.22', user_agent_md5='f7f4c471f44c4ed9070606fee66c8f58', request='file1', byte_num=3),
    Row(date='2020-07-31', epoch=1596213999, ip='63.110.194.10', user_agent_md5='9d7bee595eafab2cfffc3df4377db814', request='file2', byte_num=0),
  ]
  delivery_df = spark_session.createDataFrame(data)

  delivery_df.write.parquet(parquet_path.join("delivery.parquet").strpath,
                            mode='overwrite')

  return parquet_path.strpath


def test_query_delivery_parquet(spark_session, delivery_parquet):
  """
  Should be able to read the delivery data, and return with a valid dataframe
  """
  # Read the delivery data
  df = cd.query_delivery_parquet(spark=spark_session,
                                 parquet_path=delivery_parquet,
                                 ip='63.110.194.22',
                                 ua='my_user1',
                                 asset='file1',
                                 start_range=1,
                                 end_range=3)

  # Assert that we get the expected results
  assert(df.count() == 4)
  assert(df.columns == ['ip', 'user_agent_md5', 'request', 'byte_num'])


test_is_delivered_cases = {
    'Delivered - Dupe': ('63.110.194.22', 'my_user1', 'file1', 1, 3, True),
    'Delivered - No Dupe': ('63.110.194.22', 'my_user1', 'file1', 0, 1, True),
    'Not Delivered - Dupe': ('63.110.194.22', 'my_user1', 'file1', 1, 5, False),
    'Not Delivered - No Dupe': ('63.110.194.10', 'my_user2', 'file2', 0, 5, False),
}
@pytest.mark.parametrize('ip, ua, asset, start_rng, end_rng, expected',
                         list(test_is_delivered_cases.values()),
                         ids=list(test_is_delivered_cases.keys()))
def test_is_delivered(ip, ua, asset, start_rng, end_rng, expected, delivery_parquet, spark_session):
  """
  Should correctly indicate if an asset was delivered
  """
  # Generate a delivery dataframe for testing
  df = cd.query_delivery_parquet(spark=spark_session,
                                 parquet_path=delivery_parquet,
                                 ip=ip,
                                 ua=ua,
                                 asset=asset,
                                 start_range=start_rng,
                                 end_range=end_rng)

  # Compare to the expected results
  assert(cd.is_delivered(df, start_rng, end_rng) ==  expected)
