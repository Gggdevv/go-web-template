CREATE DATABASE IF NOT EXISTS database1;
USE database1;
CREATE TABLE `users`(
    `id` int(11) NOT NULL,
    `username` varchar(255) NOT NULL,
    `password` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL,
    `phone` varchar(255) NOT NULL,
    `address` varchar(255) NOT NULL,
    `role` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL
);