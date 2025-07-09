-- create the database
create database upc_tracker;
set search_path to upc_tracker;

-- switch to the database (for psql or pgAdmin usage)
\c myplace;

-- create the accounts table
create table accounts (
    id bigserial primary key,
    email varchar(255) not null unique,
    password text not null,
    name varchar(255) not null
);

-- create the products table
create table products (
    id bigserial primary key,
    upc varchar(32) not null,
    name varchar(255) not null,
    descr text
);

create table places (
    id bigserial primary key,
    name varchar(255) not null,
    address varchar(255) not null,
    city varchar(255) not null,
    state varchar(255) not null
)