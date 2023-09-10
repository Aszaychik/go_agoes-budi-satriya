CREATE DATABASE IF NOT EXISTS dbms_online_shop DEFAULT CHARACTER SET = 'utf8mb4';

-- Use the database.
USE dbms_online_shop;

CREATE TABLE IF NOT EXISTS operators (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_types (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS product_description (
    id INT NOT NULL AUTO_INCREMENT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS products (
    id INT NOT NULL AUTO_INCREMENT,
    product_type_id INT NOT NULL,
    operator_id INT NOT NULL,
    code VARCHAR(50),
    name VARCHAR(100),
    status SMALLINT,
    descriptions INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (product_type_id) REFERENCES product_types (id),
    FOREIGN KEY (operator_id) REFERENCES operators (id),
    FOREIGN KEY (descriptions) REFERENCES product_description(id)
);

CREATE TABLE IF NOT EXISTS payment_methods (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255),
    status SMALLINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    status SMALLINT,
    dob DATE,
    gender CHAR(1),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS transactions (
    id INT NOT NULL AUTO_INCREMENT,
    user_id INT NOT NULL,
    payment_method_id INT NOT NULL,
    status VARCHAR(10),
    total_qty INT NOT NULL,
    total_price DECIMAL(25,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id)
);

CREATE TABLE IF NOT EXISTS transactions_detail (
    transaction_id INT NOT NULL,
    product_id INT NOT NULL,
    status VARCHAR(10),
    qty INT NOT NULL,
    price DECIMAL(25,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (transaction_id, product_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);

# 1A Insert 5 operators pada table operators.
INSERT INTO operators (name) VALUES 
  ('OP 1'),
  ('OP 2'),
  ('OP 3'),
  ('OP 4'),
  ('OP 5');


# 1B Insert 3 product type.
INSERT INTO product_types (name) VALUES 
  ('Storage'),
  ('Memory'),
  ('Processor');

# 1C Insert 2 product dengan product type id = 1, dan operators id = 3.
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
  (1, 3, 'SSD-SEAGATE-FIRECUDA530Lightsaber-2TB', 'Seagate FireCuda 530 Lightsaber 2TB M.2 2280 NVMe Gen 4', 1),
  (1, 5, 'SSD-SAMSUNG-990PRO-2TB', 'Samsung SSD 990 PRO Heatsink 2TB M.2 NVME PCIe Gen 4 M2NVMe 2.0', 1);

# 1D Insert 3 product dengan product type id = 2, dan operators id = 1
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES 
  (2, 1,'RAM-CORSAIR-DOMINATOR-DDR5-64GB', 'CORSAIR DOMINATOR PLATINUM RGB 64GB (2x32GB) DDR5 6600MHz C32', 1),
  (2, 1, 'RAM-GIGABYTE-AORUS-DDR5-32GB', 'GIGABYTE AORUS MEMORY DDR5 32GB (2x16GB) 5200MHz', 1);

# 1E Insert 3 product dengan product type id = 3, dan operators id = 4.
INSERT INTO products (product_type_id, operator_id, code, name, status) VALUES
  (3, 4, 'CPU-INTEL-CORE-I7-10700K', 'Intel Core i7 10700K', 1),
  (3, 4, 'CPU-INTEL-CORE-I5-10400K', 'Intel Core i5 10400K', 1);

# 1F Insert product description pada setiap product.
INSERT INTO product_description (description) VALUES
  ('SSD Seagate FireCuda 530 Lightsaber Edition With 2TB Storage M.2 2280 NVMe Gen 4'),
  ('SSD Samsung 990 PRO Edition With 2TB Storage M.2 NVMe PCIe Gen 4 M2NVMe 2.0'),
  ('RAM Corsair Dominator Platinum Edition RGB 64GB (2x32GB) DDR5 6600MHz C32'),
  ('RAM GigaByte Aorus Memory DDR5 32GB (2x16GB) 5200MHz'),
  ('CPU Intel Core i7 10700K'),
  ('CPU Intel Core i5 10400K');

UPDATE products
SET descriptions = 
  CASE
      WHEN id = 1 THEN 1
      WHEN id = 2 THEN 2
      WHEN id = 3 THEN 3
      WHEN id = 4 THEN 4
      WHEN id = 5 THEN 5
      WHEN id = 6 THEN 6
      ELSE descriptions
  END;

SELECT * FROM products
JOIN product_description ON (products.descriptions = product_description.id);

# 1G Insert 3 payment methods.
INSERT INTO payment_methods (name, status) VALUES
  ('DEBIT', 1),
  ('CREDIT', 1),
  ('CASH', 1),
  ('E-WALLET', 1),
  ('CRYPTO', 1);

# 1H Insert 5 user pada tabel user.
INSERT INTO users (status, name, dob, gender) VALUES
  (1, 'AGOES BUDI SATRIYA', '2001-10-11', 'M'),
  (1, 'KAYOKO ONIKATA', '2004-03-17', 'F'),
  (1, 'RIKUHACHIMA ARU', '2006-03-12', 'F'),
  (1, 'IGUSA HARUKA', '2007-05-13', 'F'),
  (1, 'ASAGI MUTSUKI', '2006-07-29', 'F');

# 1I Insert 3 transaksi di masing-masing user. (soal berlanjut ke soal 1.j)

-- User 1
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (1, 1, 'success', 3, 100000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (1, 2, 'pending', 2, 50000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (1, 3, 'failed', 1, 25000);

-- User 2
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (2, 4, 'success', 2, 75000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (2, 5, 'pending', 1, 35000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (2, 1, 'failed', 3, 105000);

-- User 3
INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (3, 2, 'success', 1, 20000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (3, 3, 'pending', 2, 50000);

INSERT INTO transactions (user_id, payment_method_id, status, total_qty, total_price)
VALUES (3, 4, 'failed', 3, 75000);


# 1J Insert 3 product di masing-masing transaksi.

-- Transaction 1 for User 1
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (1, 1, 1, 30000);

INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (1, 2, 2, 50000);

INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (1, 3, 1, 25000);

-- Transaction 2 for User 1
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (2, 4, 1, 35000);

INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (2, 5, 1, 50000);

-- Transaction 3 for User 1
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (3, 6, 1, 25000);

-- Transaction 1 for User 2
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (4, 1, 1, 30000);

INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (4, 2, 2, 50000);

-- Transaction 2 for User 2
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (5, 3, 1, 25000);

-- Transaction 3 for User 2
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (6, 5, 1, 50000);

-- Transaction 1 for User 3
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (7, 4, 1, 35000);

-- Transaction 2 for User 3
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (8, 5, 2, 50000);

-- Transaction 3 for User 3
INSERT INTO transactions_detail (transaction_id, product_id, qty, price)
VALUES (9, 6, 1, 25000);

# 2 Select
# 2A Tampilkan nama user / pelanggan dengan gender Laki-laki / M.
SELECT NAME FROM users WHERE GENDER = 'M';

# 2B Tampilkan product dengan id = 3.
SELECT * FROM products WHERE id = 3;

# 2C Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata ‘a’.
SELECT * 
FROM users 
WHERE created_at >= CURDATE() - INTERVAL 7 DAY 
  AND name LIKE '%a%';

# 2D Hitung jumlah user / pelanggan dengan status gender Perempuan.
SELECT COUNT(*) AS count_female_user
FROM users
WHERE gender = 'F';

# 2E Tampilkan data pelanggan dengan urutan sesuai nama abjad
SELECT *
FROM users
ORDER BY name ASC;

# 2F Tampilkan 5 data pada data product
SELECT * 
FROM products
LIMIT 5;

# 3 Update
# 3A Ubah data product id 1 dengan nama ‘product dummy’.
UPDATE products
SET name = 'product dummy'
WHERE id = 1;

# 3B Update qty = 3 pada transaction detail dengan product id = 1.
UPDATE transactions_detail
SET qty = 3
WHERE product_id = 1;

# 4 Delete
# 4A Delete data pada tabel product dengan id = 1.
DELETE FROM products
WHERE id = 1;

# 4B Delete pada pada tabel product dengan product type id 1.
DELETE FROM products
WHERE product_type_id = 1;

# Join Union Subquery Function

# 1 Gabungkan data transaksi dari user id 1 dan user id 2.
SELECT *
FROM transactions
WHERE user_id = 1
UNION ALL
SELECT *
FROM transactions
WHERE user_id = 2;

# 2 Tampilkan jumlah harga transaksi user id 1.
SELECT user_id, SUM(total_price) AS total_price
FROM transactions
WHERE user_id = 1;

# 3 Tampilkan total transaksi dengan product type 2.
SELECT 
  sum(transactions_detail.qty),
  sum(transactions_detail.price)
FROM transactions
JOIN transactions_detail ON(transactions.id=transactions_detail.transaction_id)
JOIN products ON(transactions_detail.product_id=products.id)
WHERE products.product_type_id = 2;

#4 Tampilkan semua field table product dan field name table product type yang saling berhubungan.
SELECT *,product_types.name
FROM products
JOIN product_types ON (products.product_type_id=product_types.id);

# 5 Tampilkan semua field table transaction, field name table product dan field name table user.
SELECT
  transactions.*,
  products.name AS product_name,
  users.name AS user_name,
  transactions_detail.product_id
FROM transactions
JOIN users ON transactions.user_id = users.id
JOIN transactions_detail ON transactions.id = transactions_detail.transaction_id
JOIN products ON transactions_detail.product_id = products.id;

# 6 Buat function setelah data transaksi dihapus maka transaction detail terhapus juga dengan transaction id yang dimaksud.
DELIMITER //
CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transactions_detail
    WHERE transaction_id = OLD.id;
END;
//
DELIMITER ;

# 7 Buat function setelah data transaksi detail dihapus maka data total_qty terupdate berdasarkan qty data transaction id yang dihapus.
DELIMITER //
CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transactions_detail
FOR EACH ROW
BEGIN
    DECLARE trans_id INT(11);
    SET trans_id = OLD.transaction_id;
    UPDATE transactions
    SET total_qty = total_qty - OLD.qty
    WHERE id = trans_id;
END;
//
DELIMITER ;

# 8 Tampilkan data products yang tidak pernah ada di tabel transaction_details dengan sub-query.
SELECT *
FROM products
WHERE id NOT IN (
    SELECT DISTINCT product_id
    FROM transactions_detail
);