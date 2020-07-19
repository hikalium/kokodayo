#!/bin/bash -e
mysql --defaults-extra-file=mysql_client.conf \
  -h ${MYSQL_HOSTNAME} \
  -P ${MYSQL_PORT} << EOS
drop database if exists ${MYSQL_DATABASE};
create database ${MYSQL_DATABASE};

use ${MYSQL_DATABASE}

DROP TABLE IF EXISTS boxes;
CREATE TABLE boxes (
  box_id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  image_url longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
describe boxes;

DROP TABLE IF EXISTS items;
CREATE TABLE items (
  item_id bigint NOT NULL AUTO_INCREMENT PRIMARY KEY,
  box_id bigint NOT NULL,
  image_url longtext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
describe items;
EOS

echo "DB ${MYSQL_DATABASE} created!"
