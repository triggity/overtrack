CREATE TABLE game_stats (
    id SERIAL,
    game_id integer NOT NULL,
    user_id integer NOT NULL,
    eliminations integer,
    objective_kills integer,
    objective_time integer,
    hero_damage integer, 
    healing integer,
    deaths integer
);