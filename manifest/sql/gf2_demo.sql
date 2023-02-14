/*
$ mysql -uroot -p'123456' < manifest/sql/gf2_demo.sql
*/

-- Create demo database 
CREATE DATABASE IF NOT EXISTS `gf2_demo`;

USE `gf2_demo`;

-- Create demo table
DROP TABLE IF EXISTS `demo`;
CREATE TABLE `demo`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `fielda`  varchar(45) NOT NULL COMMENT 'Field demo',
    `fieldb`  varchar(45) NOT NULL COMMENT 'Private field demo',
    `created_at` datetime DEFAULT NULL COMMENT 'Created Time',
    `updated_at` datetime DEFAULT NULL COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_fielda` (`fielda`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Demo records of table `demo`
BEGIN;
INSERT INTO `demo` (fielda, fieldb, created_at, updated_at) VALUES ('windvalley', 'q0RybNeqtIra', '2008-08-08 08:08:08', '2008-08-08 08:08:08');
INSERT INTO `demo` (fielda, fieldb, created_at, updated_at) VALUES ('noone', 'YVnb95KsUL9j', '2009-08-08 08:08:08', '2009-08-08 08:08:08');
COMMIT;
