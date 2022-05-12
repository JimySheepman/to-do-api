CREATE TABLE tasks (
    id serial PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    content VARCHAR (160) NOT NULL,
    category VARCHAR (50) NOT NULL,
    statu  VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);


INSERT into tasks (title,"content",category,statu,created_at) 
VALUES ('Test','Test','Test','Test', current_timestamp);

INSERT into tasks (title,"content",category,statu,created_at) 
VALUES ('Test1','Test1','Test1','Test1', current_timestamp);


INSERT into tasks (title,"content",category,statu,created_at) 
VALUES ('Test2','Test2','Test2','Test2', current_timestamp);

INSERT into tasks (title,"content",category,statu,created_at) 
VALUES ('Test3','Test3','Test3','Test3', current_timestamp);