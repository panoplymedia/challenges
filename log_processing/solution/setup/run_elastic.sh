#!/bin/bash

if [[ ! $(docker image ls | grep -o "docker.elastic.co/elasticsearch/elasticsearch") ]]; then
    docker pull docker.elastic.co/elasticsearch/elasticsearch:7.8.1
fi

docker run -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.8.1