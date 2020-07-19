#!/bin/bash -e
: ${MYSQL_HOSTNAME:='127.0.0.1'}
: ${MYSQL_PORT:="`docker port kokodayo-mysql 3306/tcp | cut -d ':' -f 2`"}
: ${MYSQL_USER:='root'}
: ${MYSQL_DATABASE:="kokodayo_test"}
: ${MYSQL_PASSWORD:="kokodayo"}

echo "MYSQL_HOSTNAME = ${MYSQL_HOSTNAME}"
echo "MYSQL_PORT     = '${MYSQL_PORT}'"
echo "MYSQL_USER     = ${MYSQL_USER}"
echo "MYSQL_DATABASE = ${MYSQL_DATABASE}"
echo "MYSQL_PASSWORD = ${MYSQL_PASSWORD}"
