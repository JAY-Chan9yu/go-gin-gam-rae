CREATE TABLE User(
    id VARCHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    name VARCHAR(100),
) DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

INSERT INTO User(name) VALUES ("제이지");
INSERT INTO User(name) VALUES ("요요");
INSERT INTO User(name) VALUES ("귀도반로섬");
