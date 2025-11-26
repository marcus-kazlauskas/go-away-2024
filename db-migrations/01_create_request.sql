-- +goose Up
-- +goose StatementBegin
create table if not exists request(
    id bigint primary key generated always as identity,
    year int not null,
    day int not null,
    part int not null,
    created_at timestamptz not null,
    s3_link text
);

comment on table request is 'Incoming task request';
comment on column request.id is 'Task ID';
comment on column request.year is 'Year of AoC challenge';
comment on column request.day is 'Puzzle day';
comment on column request.part is 'Puzzle part';
comment on column request.created_at is 'Task creation date';
comment on column request.s3_link is 'Task file link in S3';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists request;
-- +goose StatementEnd