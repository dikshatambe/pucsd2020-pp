CREATE DATABASE IF NOT EXIST acl;
drop database acl;
create database acl;
use acl;

DROP TABLE  IF NOT EXIST Users;
DROP TABLE  IF NOT EXIST Groups;
DROP TABLE  IF NOT EXIST Files;
DROP TABLE  IF NOT EXIST UserGroup;
DROP TABLE  IF NOT EXIST Permission;

CREATE TABLE Groups (
    GroupId int NOT NULL,
    GroupName varchar(10) NOT NULL,
    Owner varchar(20),
    password varchar(10),
    PRIMARY KEY (GroupId),
);

CREATE TABLE Permission (
    PId int NOT NULL,
    PName varchar NOT NULL,
);

CREATE TABLE Users (
    UserId int NOT NULL,
    userName varchar(10) NOT NULL,
    password varchar(10),
    FileId  INT,
    PId  INT,
    PRIMARY KEY (UserId),
    FOREIGN KEY (FileId) REFERENCES File(FileId),
    FOREIGN KEY (PId) REFERENCES Permission(PId)
);

CREATE TABLE Files (
    FileId int NOT NULL,
    FileName varchar(10) NOT NULL,
    Owner varchar(20),
    FOREIGN KEY (GroupId) REFERENCES Groups(GroupId),
    FOREIGN KEY (FileId) REFERENCES File(FileId),
    FOREIGN KEY (RoleId) REFERENCES Permission(PId)
    PRIMARY KEY (FileId),
);

CREATE TABLE IF NOT EXISTS UserGroup (
        UserId          INT,
        GroupId         INT,
        PRIMARY KEY (UserId, GroupId),
        FOREIGN KEY (UserId) REFERENCES Users(UserId),
        FOREIGN KEY (GroupId) REFERENCES Groups(GroupId)
);


