CREATE TABLE `cr_permissions`
(
    `id`         BIGINT UNSIGNED AUTO_INCREMENT        NOT NULL,
    `name`       VARCHAR(40)                           NOT NULL,
    `created_by` VARCHAR(40)                           NOT NULL,
    `updated_by` VARCHAR(40)                           NOT NULL,
    `created_at` DATETIME  DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `deleted_at` TIMESTAMP                             NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;