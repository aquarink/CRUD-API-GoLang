-- Adminer 4.8.1 MySQL 5.5.68-MariaDB dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

DROP TABLE IF EXISTS `playing`;
CREATE TABLE `playing` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `msisdn` varchar(20) NOT NULL,
  `id_ws` int(11) NOT NULL,
  `playing_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


DROP TABLE IF EXISTS `prizes`;
CREATE TABLE `prizes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `index_prize` int(11) NOT NULL,
  `prize_name` varchar(250) NOT NULL,
  `qouta` int(11) NOT NULL,
  `day_limit` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `index_prize` (`index_prize`),
  KEY `qouta` (`qouta`),
  KEY `day_limit` (`day_limit`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `prizes` (`id`, `index_prize`, `prize_name`, `qouta`, `day_limit`, `created_at`) VALUES
(1,	8,	'Kuota 1 Giga',	10,	1,	'2020-07-27 16:32:28'),
(2,	4,	'Kuota 500 megabyte',	10,	1,	'2020-07-27 16:32:28'),
(3,	150,	'Ganti dari Postman',	115,	125,	'2020-07-27 16:32:28'),
(4,	2,	'Pulsa Rp.10.000',	10,	1,	'2020-07-27 16:32:28'),
(5,	1,	'Voucher Game 1 Hari',	0,	0,	'2020-07-27 16:32:28'),
(6,	7,	'Voucher Game 2 Hari',	0,	0,	'2020-07-27 16:32:28'),
(7,	3,	'Voucher Game 2 Hari',	0,	0,	'2020-07-27 16:32:28'),
(8,	5,	'Voucher Game 1 Hari',	0,	0,	'2020-07-27 16:32:28'),
(10,	12,	'Percobaan Service',	12,	3,	'2021-10-08 16:42:51'),
(11,	15,	'Coba dari Postman',	15,	15,	'2021-10-08 16:56:47'),
(12,	15,	'Coba dari Postman',	15,	15,	'2021-10-11 15:06:14');

DROP TABLE IF EXISTS `take_prize`;
CREATE TABLE `take_prize` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `prize_id` int(11) NOT NULL,
  `prize_index` int(11) NOT NULL,
  `msisdn` varchar(20) NOT NULL,
  `take_datetime` datetime NOT NULL,
  PRIMARY KEY (`id`),
  KEY `prize_id` (`prize_id`),
  KEY `msisdn` (`msisdn`),
  KEY `prize_index` (`prize_index`),
  KEY `take_datetime` (`take_datetime`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `take_prize` (`id`, `prize_id`, `prize_index`, `msisdn`, `take_datetime`) VALUES
(1,	2,	4,	'cpmvcr',	'2020-07-29 22:41:18'),
(2,	2,	4,	'CPMVCR',	'2020-07-30 15:53:04');

-- 2021-10-11 09:22:54
