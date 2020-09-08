#!/usr/bin/env bash

# Launch process_pagecounts inside a docker container
docker run \
  --rm \
  -v $PWD:/repo \
  -v $PWD/data:/data \
  --env PYSPARK_DRIVER_PYTHON=/usr/bin/python3.6 \
  --env PYSPARK_PYTHON=/usr/bin/python3.6 \
  cjonesy/docker-spark:2.4.4 \
  /bin/bash -c 'pip3 install pytest -q && pytest --color=yes --disable-warnings -vvv /repo/spark/'
