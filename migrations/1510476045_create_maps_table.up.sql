CREATE TYPE game_type as ENUM ("escort", "assult", "hybrid", "control");

CREATE TABLE maps (
  id            integer unique,
  name          varchar(40) unique,
  full_name     varchar(40),
  city          varchar(40),
  country       varchar(40)
  game_type     game_type
);