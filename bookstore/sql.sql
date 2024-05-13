-- 创建users表
create table users(
    id int primary key not null auto_increment,
    username varchar(100) not null,
    password varchar(100) not null,
    email varchar(100) not null
);
-- 创建books表
create table books(
    id int primary key not null auto_increment,
    title varchar(100) not null,
    author varchar(100) not null,
    price double not null,
    sales int not null,
    stock int not null,
    img_path varchar(100) not null,
);
-- 书的数据
INSERT INTO `books` VALUES (1, '解忧杂货店', '东野圭吾', 27, 100, 100, 'static/img/解忧杂货店.jpg');
INSERT INTO `books` VALUES (2, '边城', '沈从文', 23, 101, 99, 'static/img/边城.jpg');
INSERT INTO `books` VALUES (3, '中国哲学史', '冯友兰', 44, 101, 99, 'static/img/中国哲学史.jpg');
INSERT INTO `books` VALUES (4, '忽然七日', ' 劳伦', 19, 104, 96, 'static/img/忽然七日.jpg');
INSERT INTO `books` VALUES (5, '苏东坡传', '林语堂', 19, 100, 100, 'static/img/苏东坡传.jpg');
INSERT INTO `books` VALUES (6, '百年孤独', '马尔克斯', 29, 100, 100, 'static/img/百年孤独.jpg');
INSERT INTO `books` VALUES (7, '扶桑', '严歌苓', 20, 102, 98, 'static/img/扶桑.jpg');
INSERT INTO `books` VALUES (8, '给孩子的诗', '北岛', 22, 102, 98, 'static/img/给孩子的诗.jpg');
INSERT INTO `books` VALUES (9, '为奴十二年', '所罗门', 16, 101, 99, 'static/img/为奴十二年.jpg');
INSERT INTO `books` VALUES (10, '平凡的世界', '路遥', 55, 101, 99, 'static/img/平凡的世界.jpg');
INSERT INTO `books` VALUES (11, '悟空传', '今何在', 14, 103, 97, 'static/img/悟空传.jpg');
INSERT INTO `books` VALUES (12, '硬派健身', '斌卡', 31, 101, 99, 'static/img/硬派健身.jpg');
INSERT INTO `books` VALUES (13, '从晚清到民国', '唐德刚', 40, 100, 100, 'static/img/从晚清到民国.jpg');
INSERT INTO `books` VALUES (14, '三体', '刘慈欣', 56, 100, 100, 'static/img/三体.jpg');
INSERT INTO `books` VALUES (15, '看见', '柴静', 19, 102, 98, 'static/img/看见.jpg');
INSERT INTO `books` VALUES (16, '活着', '余华', 11, 100, 100, 'static/img/活着.jpg');
INSERT INTO `books` VALUES (17, '小王子', '安托万', 19, 100, 100, 'static/img/小王子.jpg');
INSERT INTO `books` VALUES (18, '我们仨', '杨绛', 11, 100, 100, 'static/img/我们仨.jpg');
INSERT INTO `books` VALUES (19, '生命不息,折腾不止', '罗永浩', 25, 100, 100, 'static/img/default.jpg');
INSERT INTO `books` VALUES (20, '皮囊', '蔡崇达', 24, 100, 100, 'static/img/皮囊.jpg');
INSERT INTO `books` VALUES (21, '恰到好处的幸福', '毕淑敏', 16, 100, 100, 'static/img/恰到好处的幸福.jpg');
INSERT INTO `books` VALUES (22, '大数据预测', '埃里克', 37, 101, 99, 'static/img/大数据预测.jpg');
INSERT INTO `books` VALUES (23, '人月神话', '布鲁克斯', 56, 100, 100, 'static/img/人月神话.jpg');
INSERT INTO `books` VALUES (24, 'C语言入门经典', '霍尔顿', 45, 101, 99, 'static/img/C语言入门经典.jpg');
INSERT INTO `books` VALUES (25, '数学之美', '吴军', 30, 100, 100, 'static/img/数学之美.jpg');
INSERT INTO `books` VALUES (26, 'Java编程思想', '埃史尔', 70, 100, 100, 'static/img/Java编程思想.jpg');
INSERT INTO `books` VALUES (27, '设计模式之禅', '秦小波', 20, 100, 100, 'static/img/设计模式之禅.jpg');
INSERT INTO `books` VALUES (28, '图解机器学习', '杉山将', 34, 100, 100, 'static/img/图解机器学习.jpg');
INSERT INTO `books` VALUES (29, '艾伦图灵传', '安德鲁', 47, 100, 100, 'static/img/艾伦图灵传.jpg');
INSERT INTO `books` VALUES (30, '教父', '马里奥普佐', 29, 100, 100, 'static/img/教父.jpg');
-- 创建sessions表
create table sessions(
    session_id varchar(100) primary key not null,
    username varchar(100) not null,
    user_id int(11) not null,
    foreign key(user_id) references users(id)
);
-- 创建cart表
create table carts(
    id varchar(100) primary key not null,
    total_count int not null,
    total_amount double(11,2) not null,
    user_id int not null,
    foreign key(user_id) references users(id)
);
-- 创建cart_item表
create table cart_items(
    id int primary key not null auto_increment,
    count int not null,
    amount double(11,2) not null,
    book_id int not null,
    cart_id varchar(100) not null,
    foreign key(book_id) references books(id),
    foreign key(cart_id)references carts(id)
);
-- 创建order表
create table order(
    id varchar(100) primary key not null,
    create_time datetime not null default current_timestamp,
    total_count int not null,
    total_amount double(11,2) not null,
    state int
    user_id int not null,
    foreign key(user_id)references users(id)
);
-- 创建order_items表
create table order_items(
    id varchar(100) primary key not null,
    count double(11,2) not null,
    amount int not null,
    title varchar(100) not null,
    price double(11,2) not null,
    img_path varchar(100),
    order_id varchar(100),
    foreign key(order_id) references orders(id)
);