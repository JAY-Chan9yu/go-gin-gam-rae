CREATE TABLE Cosmetic (
    id VARCHAR(36) DEFAULT (UUID()) PRIMARY KEY,
    name VARCHAR(100),
    description VARCHAR(500),
    price int,
);

INSERT INTO Cosmetic(name,description,price) VALUES ("토리든 앰플","너무 좋아여~", 100000);
INSERT INTO Cosmetic(name,description,price) VALUES ("이니스프리 앰플","너무 좋아여~", 25000);
INSERT INTO Cosmetic(name,description,price) VALUES ("건웅 앰플","센빠이 너무 좋아여~", 1330);
