CREATE TABLE ads (
    id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    start_at timestamp with time zone NOT NULL,
    end_at timestamp with time zone NOT NULL,
    age_start int,
    age_end int,
    country varchar(100)[],
    platform varchar(100)[],
    gender varchar(100)
);
