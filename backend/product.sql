DROP TABLE IF EXISTS shop;
CREATE TABLE shop (
    id      bigserial       NOT NULL,
    name    varchar(255)    NOT NULL, 
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS product;
CREATE TABLE product (
    id          bigserial       NOT NULL,
    name        varchar(255)    NOT NULL, 
    price       numeric           NOT NULL, 
    category    varchar(255)    NOT NULL,
    shop_id     bigint          NOT NULL,   
    PRIMARY KEY (id)
);
 
INSERT INTO shop VALUES (1, 'Shop Satu');
INSERT INTO shop VALUES (2, 'Shop Dua');
INSERT INTO shop VALUES (3, 'Shop Tiga');
INSERT INTO shop VALUES (4, 'Shop Empat');
 
INSERT INTO product VALUES (1, 'Product A', 1000, 'Gadgets', 1);
INSERT INTO product VALUES (2, 'Product B', 2000, 'Gadgets', 1);
INSERT INTO product VALUES (3, 'Product C', 3000, 'Gadgets', 1);
INSERT INTO product VALUES (4, 'Product D', 4000, 'Food', 2);
INSERT INTO product VALUES (5, 'Product E', 5000, 'Drink', 2);
INSERT INTO product VALUES (6, 'Product F', 6000, 'Sports', 3);
INSERT INTO product VALUES (7, 'Product G', 7000, 'Sports', 3);
INSERT INTO product VALUES (8, 'Product H', 8000, 'Sports', 3);
INSERT INTO product VALUES (9, 'Product I', 9000, 'Sports', 3);
INSERT INTO product VALUES (10, 'Product J', 10000, 'Fashion', 4);