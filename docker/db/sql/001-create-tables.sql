
---- drop ----
DROP TABLE IF EXISTS `samples`;

---- create ----
create table IF not exists `samples`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
 `deleted_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);

---- drop ----
DROP TABLE IF EXISTS `users`;

---- create ----
create table IF not exists `users`
(
 `id`               INT(20) AUTO_INCREMENT,
 `userID`           VARCHAR(20) NOT NULL,
 `password`         VARCHAR(20) NOT NULL,
 `name`             VARCHAR(20) NOT NULL,
 `Email`            VARCHAR(50) NOT NULL,
 `Year`             INT(4) NOT NULL,
 `Month`            INT(2) NOT NULL,
 `Day`              INT(2) NOT NULL,
 `Sex`              VARCHAR(3) NOT NULL,
--  `created_at`       Datetime DEFAULT NULL,
--  `updated_at`       Datetime DEFAULT NULL,
--  `deleted_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
);
---- insert ----
INSERT INTO users values (1,'trial_user','trial_password','trial','trial@abc.co.jp',1997,1,1,'男');
