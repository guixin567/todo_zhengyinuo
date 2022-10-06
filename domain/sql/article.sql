### 创建文章表
CREATE TABLE `article`(
    `id` bigint(20) AUTO_INCREMENT PRIMARY KEY ,
    `community_id` int UNSIGNED NOT NULL ,
    `community_name` varchar(128) NOT NULL ,
    `author_id` bigint(20) NOT NULL ,
    `article_id` bigint(20) NOT NULL ,
    `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
    `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL ,
    `status` int UNSIGNED DEFAULT 1,
    `create_time` timestamp DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;





### 文章表结构设计
### ID 数据类型 bigint(20) 自增 主健
### 社区ID 数据类型 int 非负数 不能为空
### 社区名称 数据类型 varchar(128) 不能为空
### 作者ID(用户ID) 数据类型 bigint(20) 不能为空
### 文章ID 数据类型 bigint(20) 不能为空
### 文章标题 数据类型 varchar(128) 不能为空
### 文章内容 数据类型 varchar(256) 不能为空
### 文章状态 数据类型 int 非负数  默认值 1 不能为空
### 创建时间 数据类型 时间戳 默认值 当前时间戳
### 更新时间 数据类型 时间戳 默认值 当前时间戳 字段更新也随着更新