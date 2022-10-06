CREATE TABLE `user`(
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NUll,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL ,
    `email` varchar(64) COLLATE utf8mb4_general_ci,
    `gender` tinyint(64) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING  BTREE,
    UNIQUE KEY `idx_user_id` (`user_id`) USING  BTREE
)ENGINE =InnoDB DEFAULT CHARSET =utf8mb4 COLLATE =utf8mb4_general_ci;


### 用户表结构设计
### id 数据类型 bigint(20) 不能为空 自增 主健
### 用户id 数据类型 bigint(20) 不能为空
### 用户名称 数据类型 varchar(64) collate utf8mb4_general_ci 不能为空
### 用户密码（加密的）数据类型 varchar(64) collate utf8mb4_general_ci 不能为空
### 邮箱 数据类型 varchar(64) collate utf8mb4_general_ci 可空
### 创建时间 数据类型 时间戳 默认值 当前时间戳
### 更新时间 数据类型 时间戳 默认值 当前时间戳 只要字段更新就更新为当前时间
