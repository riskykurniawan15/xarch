--Up
CREATE TABLE `user_tokens` (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(11) UNSIGNED NOT NULL,
  `method` enum('verified','forgot') NOT NULL,
  `token` varchar(100) UNIQUE NOT NULL,
  `expired` TIMESTAMP NULL,
  `created_at` TIMESTAMP default now(),
  `updated_at` TIMESTAMP NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES users(id)
    ON UPDATE CASCADE ON DELETE CASCADE,
  INDEX (method, token),
  INDEX (user_id, method, token)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
--EndUp
--Down
DROP TABLE IF EXISTS `user_tokens`;
--EndDown