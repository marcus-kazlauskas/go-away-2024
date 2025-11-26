-- +goose Up
-- +goose StatementBegin
create type result_status_enum as enum ('CREATED', 'STARTED', 'COMPLITED', 'ERROR');

create table if not exists result(
    id bigint primary key generated always as identity,
    status result_status_enum not null
        default 'CREATED',
    result text,
    started_at timestamptz,
    completed_at timestamptz,
    request_id bigint not null 
        constraint fk_request_id
            references request
            on delete cascade
);

comment on table result is 'Task solution result';
comment on column result.id is 'Result ID';
comment on column result.status is 'Result status';
comment on column result.result is 'Result value';
comment on column result.started_at is 'Start date of the task solving';
comment on column result.completed_at is 'Completion date of the task solving';
comment on column result.request_id is 'Link to incoming task request';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists result;
drop type result_status_enum;
-- +goose StatementEnd