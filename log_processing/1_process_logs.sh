#!/usr/bin/env bash

# Launch update_languages inside a docker container
docker run \
  --rm \
  -v $PWD:/repo \
  -v $PWD/data:/data \
  --env PYSPARK_DRIVER_PYTHON=/usr/bin/python3.6 \
  --env PYSPARK_PYTHON=/usr/bin/python3.6 \
  cjonesy/docker-spark:2.4.4 \
  python3.6 /repo/spark/process_log.py $@
