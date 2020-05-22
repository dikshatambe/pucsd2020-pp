DROP DATABASE IF EXISTS acl;
CREATE DATABASE acl;
use acl;

DROP TABLE  IF EXISTS Groups;
DROP TABLE  IF EXISTS Users;
DROP TABLE  IF EXISTS Files;
DROP TABLE  IF EXISTS Folders;
DROP TABLE  IF EXISTS permissions;
DROP TABLE  IF EXISTS UserGroup;
DROP TABLE  IF EXISTS FileFolder;
DROP TABLE  IF EXISTS FilePermission;
DROP TABLE  IF EXISTS FolderPermission;
DROP TABLE  IF EXISTS GroupPermission;

CREATE TABLE Groups (
    gid   int NOT NULL,
    GroupName varchar(10) NOT NULL,
    PRIMARY KEY (gid)
);

CREATE TABLE Users (
    id   int NOT NULL,
    FirstName varchar(10) NOT NULL,
    LastName varchar(10) NOT NULL,
    UserName varchar(10) NOT NULL,
    password varchar(10) NOT NULL,
    IsUserRoot  int,
    PRIMARY KEY (id)
);


CREATE TABLE Files (
    fid   int NOT NULL,
    FileName varchar(10) NOT NULL,
    FilePath varchar(50),
    PRIMARY KEY (fid)
);

CREATE TABLE Folders (
    foid   int,
    FolderName varchar(20) NOT NULL,
    PRIMARY KEY (foid)
);

CREATE TABLE permissions (
    pid   int NOT NULL,
    PName varchar(10) NOT NULL,
    PRIMARY KEY (pid)
);

CREATE TABLE UserGroup (
        id  int,
        gid int,
        PRIMARY KEY (id, gid),
        FOREIGN KEY(id) REFERENCES Users(id),
        FOREIGN KEY(gid) REFERENCES Groups(gid)
);

CREATE TABLE FileFolder (
	fid   int,
	foid int,
        PRIMARY KEY (fid, foid),
        FOREIGN KEY(fid) REFERENCES Files(fid),
        FOREIGN KEY(foid) REFERENCES Folders(foid)
);

CREATE TABLE FilePermission (
        id int,
        fid int,
        pid    int,
        PRIMARY KEY (id,fid),
        FOREIGN KEY(id) REFERENCES Users(id),
        FOREIGN KEY(fid) REFERENCES Files(fid),
        FOREIGN KEY(pid) REFERENCES permissions(pid)
);


CREATE TABLE FolderPermission (
        id   int,
        foid int,
        pid      int,
        PRIMARY KEY (id,foid),
        FOREIGN KEY(id)    REFERENCES Users(id),
        FOREIGN KEY(foid)  REFERENCES Folders(foid),
        FOREIGN KEY(pid)       REFERENCES permissions(pid)
);


CREATE TABLE IF NOT EXISTS GroupPermission
(
    gp_id       INT     AUTO_INCREMENT      PRIMARY KEY,
    fid INT NOT NULL,
    gid INT  NOT NULL, #group
    pid INT NOT NULL,
    FOREIGN KEY(gid) REFERENCES Groups(gid) ,
    FOREIGN KEY(fid) REFERENCES Files(fid) ,
    FOREIGN KEY(pid) REFERENCES permissions(pid)
);

