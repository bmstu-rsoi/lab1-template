-- +goose Up

CREATE SCHEMA IF NOT EXISTS person;

CREATE TABLE IF NOT EXISTS person.persons(
    id bigint generated always as identity primary key
    ,name text not null
    ,age int
    ,address text
    ,work text
);
 -- check(length(name) < 1000),

-- +goose Down
DROP SCHEMA IF EXISTS person CASCADE;