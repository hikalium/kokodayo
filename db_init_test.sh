#!/bin/bash
MYSQL_DATABASE=kokodayo_test
. env.sh
mysql --defaults-extra-file=mysql_client.conf \
  -h ${MYSQL_HOSTNAME} \
  -P ${MYSQL_PORT} << EOS
drop database if exists ${MYSQL_DATABASE};
create database ${MYSQL_DATABASE};

use ${MYSQL_DATABASE}

DROP TABLE IF EXISTS items;
CREATE TABLE items (
  item_id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  image_url longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
describe items;

# init data
INSERT INTO items (image_url) VALUES (
"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+C9QDwADpgGQ/B4OjgAAAABJRU5ErkJggg=="
);
select * from items;

EOS
echo "DONE!"
