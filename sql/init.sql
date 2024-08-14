create database upay;

use upay;

create table account(
    id bigint primary key auto_increment comment '主键id',
    username varchar(32) not null comment '登录用户名',
    password varchar(128) not null comment '登录密码',
    email varchar(255) not null comment '用户邮箱',
    status varchar(16) not null default 'created' comment '用户状态：active, suspended, closed, pending, locked, deleted',
    secret varchar(8) not null comment '用户密码加密盐值(8位任意值)',
    create_at datetime default now() not null comment '创建时间'
) comment '后台账户表';

create table token(
    id bigint primary key auto_increment comment '主键id',
    target_id bigint not null comment '目标id，如果',
    type varchar(16) not null comment 'token类型 account: 管理员账户, 商户: merchants',
    access_token text not null comment '访问token令牌',
    refresh_token text not null comment '刷新token令牌',
    expires_at datetime comment '刷新token令牌过期时间 如果为空则不会过期',
    create_at datetime default now() not null comment '创建时间'
) comment '用户Token令牌表';

# 初始管理员账号密码 admin 123456
insert into account (id, username, password, email, status, secret, create_at) values (null, 'admin', '39ae1deda52c5e399b5c2697689af504', 'admin@upay.com', 'created', 'GnYchJd4', now());

create table settings(
    id bigint primary key auto_increment comment '主键id',
    config varchar(64) not null comment '配置key',
    name varchar(64) comment '配置显示名称',
    value text default null comment '配置值',
    required tinyint(1) default false not null comment '值是否必填',
    type varchar(16) not null comment '数据类型',
    description text default null comment '描述',
    modified_by bigint not null comment '-1为系统创建',
    create_at datetime default now() not null comment '创建时间'
) comment '系统设置';