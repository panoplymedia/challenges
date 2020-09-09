#!/usr/bin/env bash

# Launch check_delivery inside a docker container
docker run \
  --rm \
  -v $PWD:/repo \
  -v $PWD/data:/data \
  --env PYSPARK_DRIVER_PYTHON=/usr/bin/python3.6 \
  --env PYSPARK_PYTHON=/usr/bin/python3.6 \
  cjonesy/docker-spark:2.4.4 \
  python3.6 /repo/spark/check_delivery.py \
    --ip='63.110.194.22' \
    --ua='Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0' \
    --asset='/6d8a9754-e8c7-4193-8491-58b2122c1c10' \
    --start_byte_range=10 \
    --end_byte_range=500
