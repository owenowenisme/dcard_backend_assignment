CREATE TABLE ads (
    id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    StartAt timestamp with time zone NOT NULL,
    EndAt timestamp with time zone NOT NULL,
    AgeStart int,
    AgeEnd int,
    Country varchar(100)[],
    Platform varchar(100)[]
);