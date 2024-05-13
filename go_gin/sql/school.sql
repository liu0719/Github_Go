--  创建学校数据库
create database school;
-- 切换数据库
use school;
-- 创建班级表
create table class(
    --不为空 主键 自增长
    id int(11) not null primary key auto_increment,
    name varchar(255),
    email varchar(255),
    score tinyint(4)
)
