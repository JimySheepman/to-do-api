CREATE TABLE tasks (
    id serial PRIMARY KEY,
    title VARCHAR (50) NOT NULL,
    content VARCHAR (160) NOT NULL,
    category VARCHAR (50) NOT NULL,
    statu  VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL
);