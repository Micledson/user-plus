CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `email` VARCHAR(200) NOT NULL,
    `password` VARCHAR(200) NOT NULL,
    PRIMARY KEY(`id`)
);

