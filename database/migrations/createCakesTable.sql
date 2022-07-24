-- +migrate Up
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

DROP TABLE IF EXISTS `cakes`;

CREATE TABLE `cakes` (
                         `id` int(11) NOT NULL AUTO_INCREMENT ,
                         `title` varchar(256) NOT NULL,
                         `description` varchar(256) NOT NULL,
                         `rating` float NOT NULL,
                         `image` varchar(2048) NOT NULL,
                         `created_at` datetime NOT NULL,
                         `updated_at` datetime NOT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

COMMIT;

-- +migrate Down
DROP TABLE `cakes`;