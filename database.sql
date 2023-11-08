/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
CREATE table IF NOT EXISTS account
(
    id           serial PRIMARY KEY,
    full_name    VARCHAR(60) NOT NULL,
    phone_number VARCHAR(13) UNIQUE NOT NULL,
    passhash     VARCHAR(200) NOT NULL
);

CREATE table IF NOT EXISTS login_log
(
    id           serial PRIMARY KEY,
    account_id   BIGINT NOT NULL,
    created_time timestamp NOT NULL
);
