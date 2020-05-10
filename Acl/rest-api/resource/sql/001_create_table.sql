DROP DATABASE IF EXISTS acl;
CREATE DATABASE acl;
use acl;

DROP TABLE  IF EXISTS Groups;
DROP TABLE  IF EXISTS Users;
DROP TABLE  IF EXISTS Files;
DROP TABLE  IF EXISTS Folders;
DROP TABLE  IF EXISTS Permissions;
DROP TABLE  IF EXISTS UserGroup;
DROP TABLE  IF EXISTS FileFolder;
DROP TABLE  IF EXISTS FilePermission;
DROP TABLE  IF EXISTS FolderPermission;

CREATE TABLE Groups (
    GroupId   int NOT NULL,
    GroupName varchar(10) NOT NULL,
    Owner     varchar(20),
    Password  varchar(10),
    PRIMARY KEY (GroupId)
);

CREATE TABLE Users (
    UserId   int NOT NULL,
    UserName varchar(10) NOT NULL,
    Password varchar(10),
    PRIMARY KEY (UserId)
);


CREATE TABLE Files (
    FileId   int NOT NULL,
    FileName varchar(10) NOT NULL,
    Owner    varchar(20),
    PRIMARY KEY (FileId)
);

CREATE TABLE Folders (
    FolderId   int,
    FolderName varchar(20) NOT NULL,
    PRIMARY KEY (FolderId)
);

CREATE TABLE Permissions (
    PId   int NOT NULL,
    PName varchar(10) NOT NULL,
    PRIMARY KEY (PId)
);

CREATE TABLE UserGroup (
        UserId  int,
        GroupId int,
        PRIMARY KEY (UserId, GroupId),
        FOREIGN KEY(UserId) REFERENCES Users(UserId),
        FOREIGN KEY(GroupId) REFERENCES Groups(GroupId)
);

CREATE TABLE FileFolder (
	FileId   int,
	FolderId int,
        PRIMARY KEY (FileId, FolderId),
        FOREIGN KEY(FileId) REFERENCES Files(FileId),
        FOREIGN KEY(FolderId) REFERENCES Folders(FolderId)
);

CREATE TABLE FilePermission (
        UserId int,
        FileId int,
        PId    int,
        PRIMARY KEY (UserId,FileId),
        FOREIGN KEY(UserId) REFERENCES Users(UserId),
        FOREIGN KEY(FileId) REFERENCES Files(FileId),
        FOREIGN KEY(PId) REFERENCES Permissions(PId)
);


CREATE TABLE FolderPermission (
        UserId   int,
        FolderId int,
        PId      int,
        PRIMARY KEY (UserId,FolderId),
        FOREIGN KEY(UserId)    REFERENCES Users(UserId),
        FOREIGN KEY(FolderId)  REFERENCES Folders(FolderId),
        FOREIGN KEY(PId)       REFERENCES Permissions(PId)
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
