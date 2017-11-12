CREATE TYPE game_result as ENUM ('win', 'loss', 'draw');

CREATE TABLE games (
    id integer NOT NULL unique,
    user_id integer NOT NULL,
    map_id integer NOT NULL,
    result game_result,
    start_time timestamp NOT NULL,
    group_size integer,
    is_placement BOOLEAN,
    season integer,
    end_sr integer,
    begin_sr integer,
    leavers integer,
    disconnect BOOLEAN
);