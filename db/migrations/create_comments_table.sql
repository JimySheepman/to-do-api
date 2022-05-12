CREATE TABLE comments (
    id serial PRIMARY KEY,
    task_id INTEGER NOT NULL,
    user_name VARCHAR (50) NOT NULL,
    user_comment VARCHAR (160) NOT NULL,
    statu  VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_task FOREIGN KEY(task_id) REFERENCES tasks(id)
);


INSERT into comments (task_id,user_name ,user_comment ,statu ,created_at) 
VALUES (1,'testUser','testComment','approved', current_timestamp);

INSERT into comments (task_id,user_name ,user_comment ,statu ,created_at) 
VALUES (1,'testUser','testComment','rejection', current_timestamp);

INSERT into comments (task_id,user_name ,user_comment ,statu ,created_at) 
VALUES (2,'testUser','testComment','approved', current_timestamp);

INSERT into comments (task_id,user_name ,user_comment ,statu ,created_at) 
VALUES (2,'testUser','testComment','rejection', current_timestamp);