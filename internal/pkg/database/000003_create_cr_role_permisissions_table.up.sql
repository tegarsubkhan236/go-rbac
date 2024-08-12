CREATE TABLE `cr_role_permissions`
(
    `id`            BIGINT UNSIGNED AUTO_INCREMENT        NOT NULL,
    `role_id`       BIGINT UNSIGNED                       NOT NULL,
    `permission_id` BIGINT UNSIGNED                       NOT NULL,
    `created_by`    VARCHAR(40)                           NOT NULL,
    `updated_by`    VARCHAR(40)                           NOT NULL,
    `created_at`    DATETIME  DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `updated_at`    TIMESTAMP DEFAULT CURRENT_TIMESTAMP() NOT NULL,
    `deleted_at`    TIMESTAMP                             NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`role_id`) REFERENCES `cr_roles` (`id`),
    FOREIGN KEY (`permission_id`) REFERENCES `cr_permissions` (`id`)
) ENGINE = InnoDB;

CREATE INDEX `idx_role_id` ON cr_role_permissions (`role_id`);
CREATE INDEX `idx_permission_id` ON cr_role_permissions (`permission_id`);