-- file: 10-create-user-and-db.sql
CREATE DATABASE persons;
CREATE ROLE program WITH PASSWORD 'test';
GRANT ALL PRIVILEGES ON DATABASE persons TO program;
ALTER ROLE program WITH LOGIN;
CREATE TABLE Persons (
    Id serial PRIMARY KEY,
    Name VARCHAR(40), 
    Address VARCHAR(60), 
    Work VARCHAR(60), 
    Age smallint
    );