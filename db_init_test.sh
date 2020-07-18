#!/bin/bash
DBNAME=kokodayo_test
mysql --defaults-extra-file=mysql_client.conf \
  -h localhost \
  -P `docker port kokodayo-mysql 3306/tcp | cut -d ':' -f 2` << EOS
drop database if exists ${DBNAME};
create database ${DBNAME};

use ${DBNAME}

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
