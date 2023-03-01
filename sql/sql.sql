CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;
DROP TABLE IF EXISTS user;
CREATE TABLE user (
                      id int auto_increment primary key,
                      name varchar(50) not null,
                      nick varchar(50) not null unique,
                      email varchar(50) not null unique,
                      password varchar(200) not null,
                      created_at timestamp default current_timestamp()
) ENGINE=INNODB;
CREATE USER IF NOT EXISTS 'socialmedia'@'localhost' IDENTIFIED BY 'socialmedia';
GRANT ALL PRIVILEGES ON devbook.* TO 'socialmedia'@'localhost';