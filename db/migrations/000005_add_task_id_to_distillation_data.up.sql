ALTER TABLE `distillation_data`
    ADD COLUMN `task_id`  INT NOT NULL AFTER `user_id`,
    ADD INDEX `idx_task_id` (`task_id`);
