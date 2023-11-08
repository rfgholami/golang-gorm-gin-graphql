CREATE TABLE IF NOT EXISTS nettasec_go_api_table (
    id serial primary key not null,
    username character varying(45) not null,
    password character varying(45) not null
);

