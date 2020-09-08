import pytest

@pytest.fixture('session', autouse=True)
def quiet_python_logs():
  """
  Prevents python logs from spamming the output during tests by setting the
  level to ERROR.
  This fixture is automatically executed for every test session.
  """
  import logging
  noisy_loggers = ('py4j',)

  for logger in noisy_loggers:
      logging.getLogger(logger).setLevel(logging.ERROR)

@pytest.fixture('session')
def spark_session():
  from pyspark.sql import SparkSession
  return SparkSession.builder\
      .config("spark.sql.session.timeZone", "UTC")\
      .master("local[*]")\
      .appName("pytest")\
      .getOrCreate()

@pytest.fixture(scope="function")
def staging_path(tmpdir):
  """
  Creates a temporary staging path to used during testing
  """
  return tmpdir.mkdir("staging")

@pytest.fixture(scope="function")
def parquet_path(tmpdir):
  """
  Creates a temporary parquet path to used during testing
  """
  return tmpdir.mkdir("parquet")
