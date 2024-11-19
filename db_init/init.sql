SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;


create schema if not exists public;

alter schema public owner to pg_database_owner;

create table if not exists public.users(
    id serial not null primary key,
    login varchar(24) not null,
    password text not null,
    refresh_token text
);

create table if not exists public.chats(
    id serial not null primary key,
    first_user integer not null references public.users(id),
    second_user integer not null references public.users(id)
);

create table if not exists public.messages(
    id serial not null primary key,
    chat_id integer not null references public.chats(id),
    sender integer not null references public.users(id),
    reciever integer not null references public.users(id),
    send_time timestamp not null,
    message_text text
);