create table users(
    id serial not null,
    login text not null,
    password text not null
)

alter table users owner to user;