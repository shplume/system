package db

// CREATE TABLE Students (
// 	Sno INT PRIMARY KEY ,
// 	Sname VARCHAR(8) NOT NULL,
// 	Ssex CHAR(2) NOT NULL CHECK (Ssex IN('男'，'女')),
// 	Sspecialty VARCHAR(2) NOT NULL,
// 	Sclass VARCHAR(8) NOT NULL,
// 	Sphone VARCHAR(64) NOT NULL
// );

// CREATE TABLE Admins (
// 	Ano INT PRIMARY KEY IDENTITY(1,1),
// 	Aname VARCHAR(16) UNIQUE NOT NULL,
// 	Apsw VARCHAR(16) NOT NULL
// );

// CREATE TABLE Users (
// 	Uno INT PRIMARY KEY IDENTITY(1,1),
// 	Uname INT UNIQUE NOT NULL,
// 	Upsw VARCHAR(16) NOT NULL,
// );

// CREATE TABLE SC (
// 	Sno INT,
// 	Cno INT,
// 	Cname VARCHAR(16),
//  	PRIMARY KEY (Sno,Cno),
//  	FOREIGN KEY(Sno) REFERENCES Students(Sno)
//     	ON DELETE CASCADE
// 		ON UPDATE CASCADE
// 	FOREIGN KEY(Cno) REFERENCES Courses(Cno)
// 		ON DELETE CASCADE
// 		ON UPDATE CASCADE
// );

// CREATE TABLE Courses (
// 	Cno INT PRIMARY KEY IDENTITY(1,1),
// 	Cname CHAR(40) NOT NULL,
// 	Ctime INT CHECK (Ctime>=16 AND Ctime <=64),
// );
