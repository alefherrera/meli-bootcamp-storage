CREATE DATABASE mydb;
USE mydb;
CREATE TABLE users
(
    id         VARCHAR(36)   NOT NULL PRIMARY KEY,
    firstname  VARCHAR(100)  NOT NULL,
    lastname   VARCHAR(100)  NOT NULL,
    username   VARCHAR(100)  NOT NULL,
    password   VARCHAR(100)  NOT NULL,
    email      VARCHAR(100)  NOT NULL,
    ip         VARCHAR(100)  NOT NULL,
    macAddress VARCHAR(100)  NOT NULL,
    website    VARCHAR(1000) NOT NULL,
    image      VARCHAR(1000) NOT NULL
);