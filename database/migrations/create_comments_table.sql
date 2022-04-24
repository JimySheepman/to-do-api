CREATE TABLE comments (
    id serial PRIMARY KEY,
    task_id INTEGER NOT NULL,
    username VARCHAR (50) NOT NULL,
    comment VARCHAR (160) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_task FOREIGN KEY(task_id) REFERENCES tasks(id)
);