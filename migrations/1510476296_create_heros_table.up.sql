CREATE TYPE hero_class as ENUM ("attack", "defense", "tank", "support");

CREATE TABLE heros (
    id      integer unique,
    name    varchar(40),
    class   hero_class
);