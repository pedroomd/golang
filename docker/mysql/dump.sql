CREATE TABLE `example` (
  `id` INT NOT NULL auto_increment,
  PRIMARY KEY (`id`)
);


CREATE TABLE `table` (
  `id` INT NOT NULL auto_increment,
  `capacity` INT NOT NULL
  PRIMARY KEY (`id`)
);


CREATE TABLE `guest` (
  `id` INT NOT NULL auto_increment,
  `name` varchar(40) NOT NULL
  `accompanying_guests` INT NOT NULL
  `table_id` INT NOT NULL
  PRIMARY KEY (`id`),
  FOREIGN KEY (`table_id`) REFERENCES `table`(`id`)

);