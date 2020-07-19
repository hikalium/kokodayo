#!/bin/bash
MYSQL_DATABASE=kokodayo_test
. env.sh
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

# init data
# https://png-pixel.com/
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8z8DwHwAFBQIAX8jx0gAAAABJRU5ErkJggg==" # Red square
);
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPj/HwADBwIAMCbHYQAAAABJRU5ErkJggg==" # Blue square
);
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPj/HwADBwIAMCbHYQAAAABJRU5ErkJggg==" # Blue square
);
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPj/HwADBwIAMCbHYQAAAABJRU5ErkJggg==" # Blue square
);
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPj/HwADBwIAMCbHYQAAAABJRU5ErkJggg==" # Blue square
);
INSERT INTO boxes (image_url) VALUES (
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNkYPj/HwADBwIAMCbHYQAAAABJRU5ErkJggg==" # Blue square
);
INSERT INTO items (box_id, image_url) VALUES (
  1,
  "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M/wHwAEBgIApD5fRAAAAABJRU5ErkJggg==" # Green square
);

select * from boxes;
select * from items;

EOS
echo "DONE!"
