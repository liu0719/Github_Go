--创建Cart_item表
CREATE TABLE cart_items(
id INT PRIMARY KEY AUTO_INREMENT,
COUNT INT NOT NULL,
amount DOUBLE(11,2) NOT NULL,
boo_id INT NOT NULL
)

--创建购物车表
CREATE TABLE carts(
id VARCHAR(100) PRIMARY KEY,
total_count DOUBLE(11,2) NOT NULL,
total_amount INT NOT NULL,
user_id INT NOT NULL,
FOREIGN KEY (user_id) REFERENCES users(id)
)

--创建orders表
CREATE TABLE orders(
id VARCHAR(100) PRIMARY KEY,
create_time datetime NOT NULL,
total_count INT NOT NULL,
total_amount DOUBLE(11,2)NOT NULL,
state INT NOT NULL,
user_id INT ,
FOREIGN KEY(user_id)REFERENCES users(id)
)