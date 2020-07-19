#!/bin/bash
MYSQL_DATABASE=kokodayo_prod
. env.sh
. db_create.sh
echo "DONE!"
