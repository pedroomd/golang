CREATE TABLE `tables` (
  `id` INT NOT NULL auto_increment,
  `capacity` INT NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `guests` (
  `id` INT NOT NULL auto_increment,
  `name` varchar(40) NOT NULL UNIQUE,
  `accompanying_guests` INT NOT NULL,
  `table_id` INT NOT NULL UNIQUE,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`table_id`) REFERENCES `tables`(`id`)
);

CREATE TABLE `party_guests` (
  `id` INT NOT NULL auto_increment,
  `time_arrived` TIMESTAMP NOT NULL,
  `accompanying_guests` INT NOT NULL,
  `guest_id` INT NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`guest_id`) REFERENCES `tables`(`id`)
);