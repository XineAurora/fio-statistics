BEGIN;
CREATE TABLE IF NOT EXISTS fios(
    id SERIAL NOT NULL,
    name text NOT NULL,
    surname text NOT NULL,
    patronymic text,
    age bigint NOT NULL,
    gender text NOT NULL,
    nationality text NOT NULL,
    PRIMARY KEY(id)
);
COMMIT;