create table users if not exists(
    id serial not null,
    login text not null,
    password text not null
)

alter table users owner to user;