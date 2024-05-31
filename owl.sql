--- OWL System

/*
CREATE TABLE IF NOT EXISTS `name` (
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/


CREATE TABLE IF NOT EXISTS `users`(
  `userid` varchar(255) NOT NULL,
  `role` INT(255) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `validated` BOOLEAN NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `notifications`(
  `notif_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `read` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `apikey` (
  `apikey` varchar(255) NOT NULL,
  `owner_id` varchar(255) NOT NULL,
  `active` BOOLEAN NOT NULL,
  `custom` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `appointment` (
  `user_id` varchar(255) NOT NULL,
  `appointment_id` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` TEXT NOT NULL,
  `done` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `apts` (
  `name` VARCHAR(255) NOT NULL,
  `code` INT NOT NULL,
  `description` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `asset` (
  `name` VARCHAR(255) NOT NULL,
  `asset_id` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `describer` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `hard` BOOLEAN  NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `ioc` (
  `ioc_id` INT NOT NULL AUTO_INCREMENT,
  `virus_id` VARCHAR(255) NOT NULL,
  `virus_type`  VARCHAR(255) NOT NULL,
  `value` VARCHAR(255) NOT NULL,
  `source` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`ioc_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `minion` (
  `minion_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `uname` VARCHAR(255) NOT NULL,
  `user_id` VARCHAR(255)  NOT NULL,
  `group_id` VARCHAR(255) NOT NULL,
  `homedir` TEXT  NOT NULL,
  `os` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `installed` BOOLEAN NOT NULL,
  `mothership_id` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `motherships` TEXT  NOT NULL,
  `tunnel_address` TEXT  NOT NULL,
  `tls` BOOLEAN  NOT NULL,
  `owner_id` VARCHAR(255) NOT NULL,
  `last_seen` VARCHAR(255) NOT NULL,
  `is_dropper` BOOLEAN  NOT NULL,
  `generate_command` TEXT  NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `motherships` (
  `owner_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255)  NOT NULL,
  `password` VARCHAR(255)  NOT NULL,
  `mothership_id` VARCHAR(255) NOT NULL,
  `address` VARCHAR(255) NOT NULL,
  `implant_tunnel` TEXT NOT NULL,
  `admin_tunnel` TEXT NOT NULL,
  `motherships` TEXT NOT NULL,
  `description` TEXT NOT NULL,
  `tls` BOOLEAN NOT NULL,
  `certpem` TEXT NOT NULL,
  `keypem` TEXT NOT NULL,
  `active` BOOLEAN NOT NULL,
  `generate_command` TEXT NOT NULL,
  `machine_data` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `plugin` (
  `owner_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `hash` VARCHAR(255) NOT NULL,
  `ptype` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `verified` BOOLEAN NOT NULL,
  `active` BOOLEAN NOT NULL,
  `signed` BOOLEAN NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `ratings` (
  `plugin_id` VARCHAR(255) NOT NULL,
  `plugin_name` VARCHAR(255) NOT NULL,
  `hash` VARCHAR(255) NOT NULL,
  `previous_hashes` TEXT NOT NULL,
  `rates` FLOAT NOT NULL,
  `average` INT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `rater` (
  `user_id` VARCHAR(255) NOT NULL,
  `plugin_id` VARCHAR(255) NOT NULL,
  `comment` VARCHAR(255) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS `virus` (
  `apt_id` VARCHAR(255) NOT NULL,
  `virus_id` VARCHAR(255) NOT NULL,
  `hash` TEXT NOT NULL,
  `virus_type` VARCHAR(25) NOT NULL,
  `file_type` VARCHAR(25) NOT NULL,
  `comm_mode` VARCHAR(25) NOT NULL,
  `os_type` VARCHAR(25) NOT NULL,
  `description` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;


CREATE TABLE IF NOT EXISTS `yara` (
  `yara_id` VARCHAR(255) NOT NULL,
  `ioc_id` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `meta` TEXT NOT NULL,
  `condition` VARCHAR(25) NOT NULL,
  `actions` TEXT NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
