### 创建社区表
CREATE TABLE `community`
(
    `id`             BIGINT AUTO_INCREMENT PRIMARY KEY,
    `community_id`   INT UNSIGNED NOT NULL,
    `community_name` VARCHAR(128) NOT NULL,
    `introduction`   VARCHAR(256) NOT NULL,
    `create_time`    TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`    TIMESTAMP    NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    constraint `idx_community_id` unique (community_id),
    constraint `idx_community_name` unique (community_name)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

### 插入数据
INSERT INTO `community`(community_id, community_name, introduction)
VALUES (
        3,"android","android study"
       );

### 社区表结构设计
### id 数据类型 bigint(20) 不能为空 自增 主健
### 社区id 数据类型 int 非负数 不能为空
### 社区名称 数据类型 varchar(128) collate utf8mb4_general_ci 不能为空
### 社区介绍（ 数据类型 varchar(256) collate utf8mb4_general_ci 不能为空
### 创建时间 数据类型 时间戳 默认值 当前时间戳
### 更新时间 数据类型 时间戳 默认值 当前时间戳 只要字段更新就更新为当前时间