CREATE TABLE `books` (
  `id` int AUTO_INCREMENT,
  `title` varchar(100) NOT NULL,
  `author_id` int NOT NULL,
  `year` int NOT NULL,
  `publisher` varchar(100) NOT NULL,
  `description` text,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
);


CREATE TABLE `authors` (
  `id` int AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `gender` char(1) NOT NULL,
  `email` varchar(100) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `users` (
  `id` int AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `hash_password` varchar(100) NOT NULL,
  `role` varchar(50) NOT NULL,
  `is_deleted` boolean NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
);