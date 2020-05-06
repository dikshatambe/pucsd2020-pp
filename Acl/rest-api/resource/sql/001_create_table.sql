CREATE DATABASE IF NOT EXISTS acl;
drop database acl;
create database acl;
use acl;

DROP TABLE  IF EXISTS Groups;
DROP TABLE  IF EXISTS Permission;
DROP TABLE  IF EXISTS Files;
DROP TABLE  IF EXISTS Users;
DROP TABLE  IF EXISTS UserGroup;

CREATE TABLE Groups (
    GroupId int NOT NULL,
    GroupName varchar(10) NOT NULL,
    Owner varchar(20),
    password varchar(10),
    PRIMARY KEY (GroupId),
);

CREATE TABLE Permission (
    PId int NOT NULL,
    PName varchar(10) NOT NULL,
);

CREATE TABLE Files (
    FileId int NOT NULL,
    FileName varchar(10) NOT NULL,
    Owner varchar(20),
    FOREIGN KEY (GroupId) REFERENCES Groups(GroupId),
    FOREIGN KEY (FileId) REFERENCES Files(FileId),
    FOREIGN KEY (PId) REFERENCES Permission(PId)
    PRIMARY KEY (FileId),
);


CREATE TABLE Users (
    UserId int NOT NULL,
    userName varchar(10) NOT NULL,
    password varchar(10),
    FileId  INT,
    PId  INT,
    PRIMARY KEY (UserId),
    FOREIGN KEY (FileId) REFERENCES Files(FileId),
    FOREIGN KEY (PId) REFERENCES Permission(PId)
);


CREATE TABLE UserGroup (
        UserId          INT,
        GroupId         INT,
        PRIMARY KEY (UserId, GroupId),
        FOREIGN KEY (UserId) REFERENCES Users(UserId),
        FOREIGN KEY (GroupId) REFERENCES Groups(GroupId)
);

CREATE TABLE IF NOT EXISTS user_detail (
    id                  INT         AUTO_INCREMENT      PRIMARY KEY,
    first_name          CHAR(25)    NOT NULL,
    last_name           CHAR(25)    NOT NULL,
    email               CHAR(64)    NOT NULL UNIQUE,
    password            VARBINARY(128)    NOT NULL,
    contact_number      CHAR(15)    NOT NULL,
    updated_by          INT         NOT NULL DEFAULT 0,
    deleted             TINYINT(1)  NOT NULL DEFAULT 0,
    creation_date       DATETIME    DEFAULT CURRENT_TIMESTAMP,
    last_update         DATETIME    DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
