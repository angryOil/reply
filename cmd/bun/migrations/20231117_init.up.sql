SET statement_timeout = 0;

--bun:split

CREATE TABLE "public"."reply"
(
    id              SERIAL PRIMARY KEY,
    cafe_id         bigint        not null,
    board_id        bigint        not null,
    writer          bigint        not null,
    content         varchar(2000) not null,
    created_at      timestamptz,
    last_updated_at timestamptz
);
