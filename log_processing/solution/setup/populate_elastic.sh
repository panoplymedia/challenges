#!/bin/bash

OUTPUT_FILE=${1}

if [[ -z ${OUTPUT_FILE} ]]; then
    echo "Missing parameter OUTPUT_FILE"
    exit 1
fi

curl -H 'Content-Type: application/json' -XPUT 'localhost:9200/logs' --data-binary @mapping.json


curl -H 'Content-Type: application/x-ndjson' -XPOST 'localhost:9200/logs/_bulk' --data-binary @${OUTPUT_FILE}
