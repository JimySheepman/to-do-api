CREATE TABLE comments (
    id serial PRIMARY KEY,
    task_id INTEGER NOT NULL,
    user_name VARCHAR (50) NOT NULL,
    user_comment VARCHAR (160) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_task FOREIGN KEY(task_id) REFERENCES tasks(id)
);