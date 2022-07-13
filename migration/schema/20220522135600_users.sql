--Up
/* For Example */
CREATE TABLE `users` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL UNIQUE,
  `password` varchar(100) NOT NULL,
  `gender` enum('male','female') NOT NULL,
  `image` varchar(100) NULL,
  `verified_at` TIMESTAMP NULL,
  `created_at` TIMESTAMP NULL,
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
--EndUp
--Down
/* For Example */
DROP TABLE IF EXISTS `users`;
--EndDown