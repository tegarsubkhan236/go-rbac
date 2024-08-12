CREATE TABLE `cr_users`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT        NOT NULL,
    `username`   VARCHAR(255)                          NOT NULL,
    `name`       VARCHAR(40)                           NOT NULL,
    `email`      VARCHAR(255)                          NOT NULL,
    `password`   VARCHAR(500)                          NOT NULL,
    `status`     INT       DEFAULT 2                   NOT NULL,
    `role_id`    BIGINT UNSIGNED                       NOT NULL,
    `created_by` VARCHAR(40)                           NOT NULL,
    `updated_by` VARCHAR(40)                           NOT NULL,
    `created_at` DATETIME  DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `deleted_at` TIMESTAMP                             NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`role_id`) REFERENCES `cr_roles` (`id`)
) ENGINE = InnoDB;

CREATE INDEX `idx_role_id` ON cr_users (`role_id`);