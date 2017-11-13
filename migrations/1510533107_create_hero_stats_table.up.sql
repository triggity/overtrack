CREATE TABLE hero_stats (
    id SERIAL,
    game_id integer NOT NULL,
    hero_id integer NOT NULL,
    eliminations integer,
    objective_kills integer,
    objective_time integer,
    hero_damage integer, 
    healing integer,
    deaths integer,
    custom_stats json,
    UNIQUE (game_id, hero_id)
);