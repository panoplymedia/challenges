#!/bin/sh
# wait-for-postgres.sh
# Script lifted from
# Using Docker Compose Entrypoint To Check if Postgres is Running
# https://bit.ly/2KCdFxh 
# Author: Kelly Andrews
set -e
cmd="$@"
# service/container name in the docker-compose file
host="db"
# PostgreSQL port
port="5432"
# Database user making commection
user="postgres"
# pg_isready is a postgreSQL client tool for checking the connection
# status of PostgreSQL server
while ! pg_isready -h ${host} -p ${port} -U ${user} > /dev/null 2> /dev/null; do
   echo "Connecting to postgres Failed"
   sleep 1
done
>&2 echo "Postgres is up - executing command"
exec $cmd

