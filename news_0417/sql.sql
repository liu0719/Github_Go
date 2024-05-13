--登录mysql
mysql -u root -p
--查看数据库
show databases;
--创建新的数据库
create database news;
-- 使用新创建的数据库
use news;
--创建class表
create table class (
    id int(9) not null  primary key auto_increment,
    class varchar(100) not null 
);
--创建表
create table news_0417(
    id int(9) not null  primary key auto_increment,
    title varchar(100) not null ,
    author varchar(100) not null ,
    context varchar(100) not null ,
    class_id int(9) not null,
    create_time datetime not null,
    foreign key(class) references class(id)
)
