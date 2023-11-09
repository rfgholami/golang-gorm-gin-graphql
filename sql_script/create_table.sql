CREATE TABLE IF NOT EXISTS users (
    id serial primary key not null,
    username character varying(45) not null,
    password character varying(45) not null
);

INSERT INTO users(username,password) VALUES('kwa','root'),('admin','admin'),('john','john'),('manager','manager')

