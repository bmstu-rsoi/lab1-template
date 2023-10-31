-- file: 10-create-user-and-db.sql
CREATE DATABASE persons;
CREATE ROLE program WITH PASSWORD 'test';
GRANT ALL PRIVILEGES ON DATABASE persons TO program;
ALTER ROLE program WITH LOGIN;
\c persons;
create table if not exists persons(
    id serial NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    work text NOT NULL,
    age int NOT NULL
);
alter table persons
add primary key(id);