create database upay;

use upay;

create table account(
    id bigint primary key auto_increment comment '主键id',
    username varchar(32) not null comment '登录用户名',
    password varchar(128) not null comment '登录密码',
    status varchar(16) not null default 'created',
    create_at datetime default now() not null comment '创建时间'
) comment '后台账户表';