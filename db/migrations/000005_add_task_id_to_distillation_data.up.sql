ALTER TABLE distillation_data
ADD COLUMN task_id INT NOT NULL;

ALTER TABLE distillation_data
ADD CONSTRAINT fk_task
FOREIGN KEY (task_id) REFERENCES task(id);