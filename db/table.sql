CREATE TABLE `Users`
(
    `uid`      int(11)     NOT NULL AUTO_INCREMENT,
    `username` varchar(45) NOT NULL,
    `email`    varchar(45) NOT NULL,
    `phone`    varchar(20) NOT NULL,
    PRIMARY KEY (`uid`),
    UNIQUE KEY `username` (`username`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = latin1;

CREATE TABLE `Posts`
(
    `post_id`   int(11)                          NOT NULL AUTO_INCREMENT,
    `message`   varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
    `timestamp` timestamp                        NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `longitude` decimal(11, 6)                   NOT NULL,
    `latitude`  decimal(10, 6)                   NOT NULL,
    `u_id`      int(11)                          NOT NULL,
    PRIMARY KEY (`post_id`),
    KEY `u_id` (`u_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 1
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin;

CREATE TABLE `Likes`
(
    `post_id` int(11) DEFAULT NULL,
    `u_id`    int(11) DEFAULT NULL,
    KEY `post_id` (`post_id`),
    KEY `u_id` (`u_id`),
    CONSTRAINT `Likes_ibfk_1` FOREIGN KEY (`post_id`) REFERENCES `Posts` (`post_id`),
    CONSTRAINT `Likes_ibfk_2` FOREIGN KEY (`u_id`) REFERENCES `Users` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

CREATE TABLE `Followers`
(
    `user_id`   int(11) NOT NULL,
    `friend_id` int(11) NOT NULL,
    `stat`      int(11) NOT NULL,
    PRIMARY KEY (`user_id`, `friend_id`),
    KEY `friend_id` (`friend_id`),
    CONSTRAINT `Followers_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `Users` (`uid`),
    CONSTRAINT `Followers_ibfk_2` FOREIGN KEY (`friend_id`) REFERENCES `Users` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;


INSERT INTO Users (username, email, phone)
VALUES ("keanstudent", "kean@kean.edu", "+2016375555");

INSERT INTO Users (username, email, phone)
    VALUES ("hsstudent", "hs@kean.edu", "+2016375555");

INSERT INTO Users (username, email, phone)
    VALUES ("Arian", "Arian@kean.edu", "+2016375552");


INSERT INTO `Posts` (`message`, `timestamp`, `longitude`, `latitude`, `u_id`)
    VALUES
    ('This is my first facenote post!', '2018-04-16 16:15:21', 43.122000, -45.120000, 2);


INSERT INTO `Posts` (`message`, `timestamp`, `longitude`, `latitude`, `u_id`)
    VALUES
    ('Bored, just got home.', '2018-04-21 15:16:58', -74.229438, 40.664293, 1);

INSERT INTO `Posts` (`message`, `timestamp`, `longitude`, `latitude`, `u_id`)
    VALUES
    ('I had Pizza for breakfast!', '2018-04-21 15:16:58', -74.229438, 40.664293, 1);

INSERT INTO `Likes` (`post_id`, `u_id`)
    VALUES
    (1, 1);

INSERT INTO `Followers` (`user_id`, `friend_id`, `stat`)
    VALUES
    (1, 2, 1);

INSERT INTO `Followers` (`user_id`, `friend_id`, `stat`)
    VALUES
    (2, 1, 1);

INSERT INTO `Followers` (`user_id`, `friend_id`, `stat`)
    VALUES
    (3, 1, 1);