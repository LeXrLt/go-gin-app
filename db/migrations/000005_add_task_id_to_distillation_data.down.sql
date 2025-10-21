ALTER TABLE distillation_data DROP CONSTRAINT IF EXISTS fk_task;
ALTER TABLE distillation_data DROP COLUMN IF EXISTS task_id;