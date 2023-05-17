-- +goose Up
-- +goose StatementBegin
create table university(
                         id bigserial primary key,
                         univ_name varchar(255) not null default '',
                         facility varchar(255) not null default ''


);


create table student(
    id bigserial primary key,
    name varchar(255) not null default '',
    grades smallint check ( grades between 0 and 315) not null DEFAULT 4,
    univ_apply_id bigint references university (id)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table student;
drop table university;
-- +goose StatementEnd
