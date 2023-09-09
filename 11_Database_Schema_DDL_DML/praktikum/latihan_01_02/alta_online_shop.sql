# 1.Create database alta_online_shop
-- Check if the database exists; if not, create it.
CREATE DATABASE IF NOT EXISTS alta_online_shop DEFAULT CHARACTER SET = 'utf8mb4';

-- Use the database.
USE alta_online_shop;


# 2.Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.

# 2a.Create table user.
-- Check if the table exists before creating it.
CREATE TABLE IF NOT EXISTS user (
  user_id INT NOT NULL AUTO_INCREMENT,
  nama VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  date_of_birth DATE NOT NULL,
  status_user VARCHAR(255) NOT NULL DEFAULT 'Normal',
  gender ENUM('m', 'f') NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id)
);

# 2b.Create table product, product type, operators, product description, payment_method.
CREATE TABLE IF NOT EXISTS product_type (
  product_type_id INT NOT NULL AUTO_INCREMENT,
  product_type VARCHAR(255),
  PRIMARY KEY (product_type_id)
);

CREATE TABLE IF NOT EXISTS operator (
  operator_id INT NOT NULL AUTO_INCREMENT,
  operator_name VARCHAR(255),
  PRIMARY KEY (operator_id)
);

CREATE TABLE IF NOT EXISTS product (
  product_id INT NOT NULL AUTO_INCREMENT,
  product_name VARCHAR(255) NOT NULL,
  product_type_id INT NOT NULL,
  operator_id INT NOT NULL,
  product_price INT NOT NULL,
  product_description TEXT,
  PRIMARY KEY (product_id),
  FOREIGN KEY (product_type_id) REFERENCES product_type (product_type_id),
  FOREIGN KEY (operator_id) REFERENCES operator (operator_id)
);

CREATE TABLE IF NOT EXISTS payment_method (
  payment_method_id INT NOT NULL AUTO_INCREMENT,
  payment_method_name VARCHAR(255),
  payment_method_fee INT DEFAULT 0,
  PRIMARY KEY (payment_method_id)
);

# 2c.Create table transaction, transaction detail.
CREATE TABLE IF NOT EXISTS transaction (
  transaction_id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  transaction_date DATE DEFAULT CURRENT_DATE,
  transaction_status ENUM('p', 's', 'f') DEFAULT 'p',
  total_amount INT NOT NULL,
  PRIMARY KEY (transaction_id),
  FOREIGN KEY (user_id) REFERENCES user (user_id)
);

CREATE TABLE IF NOT EXISTS transaction_detail (
  transaction_detail_id INT NOT NULL AUTO_INCREMENT,
  transaction_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  subtotal INT NOT NULL,
  PRIMARY KEY (transaction_detail_id),
  FOREIGN KEY (transaction_id) REFERENCES transaction (transaction_id),
  FOREIGN KEY (product_id) REFERENCES product (product_id)
);

# 3.Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE IF NOT EXISTS kurir (
  kurir_id INT NOT NULL AUTO_INCREMENT,
  kurir_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (kurir_id)
)

# 4.Tambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurir ADD ongkos_dasar INT NOT NULL;

# 5.Rename tabel kurir menjadi shipping.
RENAME TABLE kurir TO shipping;

# 6.Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;

# 7.Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:

# 7a.1-to-1: payment method description.
CREATE TABLE IF NOT EXISTS payment_method_description (
  payment_method_description_id INT NOT NULL AUTO_INCREMENT,
  payment_method_id INT NOT NULL,
  description TEXT,
  PRIMARY KEY (payment_method_description_id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method (payment_method_id)
);

# 7b.1-to-many: user dengan alamat.
CREATE TABLE IF NOT EXISTS user_address (
  user_address_id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  address VARCHAR(255) NOT NULL,
  PRIMARY KEY (user_address_id),
  FOREIGN KEY (user_id) REFERENCES user (user_id)
);

# 7c.many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE IF NOT EXISTS user_payment_method (
  user_payment_method_id INT NOT NULL AUTO_INCREMENT,
  user_id INT NOT NULL,
  payment_method_id INT NOT NULL,
  PRIMARY KEY (user_payment_method_id),
  FOREIGN KEY (user_id) REFERENCES user (user_id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_method (payment_method_id)
);



# TEST QUERY
-- Active: 1692859170237@@127.0.0.1@3306@alta_online_shop
# Insert data into user table
INSERT INTO user (nama, address, date_of_birth, status_user, gender, created_at, updated_at)
VALUES ('John Doe', 'Jl. Jendral Sudirman No. 1', '1990-01-01', 'Normal', 'm', NOW(), NOW());

INSERT INTO user (nama, address, date_of_birth, status_user, gender, created_at, updated_at)
VALUES ('Jane Doe', 'Jl. Gatot Subroto No. 2', '1991-02-02', 'Normal', 'f', NOW(), NOW());

# Insert data into product_type table
INSERT INTO product_type (product_type)
VALUES ('Smartphone');

INSERT INTO product_type (product_type)
VALUES ('Laptop');

# Insert data into operator table
INSERT INTO operator (operator_name)
VALUES ('Apple');

INSERT INTO operator (operator_name)
VALUES ('Samsung');

# Insert data into product table
INSERT INTO product (product_name, product_type_id, operator_id, product_price, product_description)
VALUES ('iPhone 13 Pro', 1, 1, 15000000, 'The latest iPhone with a powerful A15 Bionic chip, a stunning Super Retina XDR display, and a triple-lens rear camera system.');

INSERT INTO product (product_name, product_type_id, operator_id, product_price, product_description)
VALUES ('Samsung Galaxy S22 Ultra', 1, 1, 14000000, 'The most advanced Samsung Galaxy phone yet, with a powerful Snapdragon 8 Gen 1 chip, a versatile 108MP camera system, and a long-lasting battery.');

# Insert data into payment_method table
INSERT INTO payment_method (payment_method_name, payment_method_fee)
VALUES ('Bank Transfer', 0);

INSERT INTO payment_method (payment_method_name, payment_method_fee)
VALUES ('Credit Card', 3);

# Insert data into transaction table
INSERT INTO transaction ( user_id, transaction_date, transaction_status, total_amount)
VALUES (1, '2023-03-08', 'p', 29000000);

INSERT INTO transaction ( user_id, transaction_date, transaction_status, total_amount)
VALUES (2, '2023-03-09', 's', 28000000);

# Insert data into transaction_detail table
INSERT INTO transaction_detail (transaction_id, product_id, quantity, subtotal)
VALUES (1, 1, 1, 15000000);

INSERT INTO transaction_detail (transaction_id, product_id, quantity, subtotal)
VALUES (1, 2, 1, 14000000);

INSERT INTO transaction_detail (transaction_id, product_id, quantity, subtotal)
VALUES (2, 2, 2, 28000000);

# Insert data into payment_method_description table
INSERT INTO payment_method_description (payment_method_id, description)
VALUES (1, 'Pay via bank transfer to our account.');

INSERT INTO payment_method_description (payment_method_id, description)
VALUES (2, 'Pay via credit card using your preferred card.');

# Insert data into user_address table
INSERT INTO user_address (user_id, address)
VALUES (1, 'Jl. Jendral Sudirman No. 1');

INSERT INTO user_address (user_id, address)
VALUES (2, 'Jl. Gatot Subroto No. 2');

# Insert data into user_payment_method table
INSERT INTO user_payment_method (user_id, payment_method_id)
VALUES (1, 1);

INSERT INTO user_payment_method (user_id, payment_method_id)
VALUES (2, 2);
