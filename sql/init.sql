create database upay;

use upay;

create table account(
    id bigint primary key auto_increment comment '主键id',
    username varchar(32) not null comment '登录用户名',
    password varchar(128) not null comment '登录密码',
    status varchar(16) not null default 'created' comment '用户状态：active, suspended, closed, pending, locked, deleted',
    token text default null comment '用户当前Token',
    refresh_token text default null comment '用户刷新Token',
    create_at datetime default now() not null comment '创建时间'
) comment '后台账户表';

create table settings(
    id bigint primary key auto_increment comment '主键id',
    config varchar(64) not null comment '配置key',
    name varchar(64) not null comment '配置显示名称',
    value text default null comment '配置值',
    type varchar(16) not null comment '数据类型',
    description text default null comment '描述',
    modified_by bigint not null comment '-1为系统创建'
) comment '系统设置';