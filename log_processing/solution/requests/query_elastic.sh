#!/bin/bash

curl -XGET -H 'Content-type: application/json' --data-binary @request.json 'localhost:9200/logs/_search?pretty'