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

INSERT INTO `community`(community_id, community_name, introduction)
VALUES (
        3,"android","android study"
       );