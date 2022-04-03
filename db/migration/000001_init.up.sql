create table realms
(
    id   serial
        constraint realms_pk
            primary key,
    name text not null
);

alter table realms
    owner to postgres;

create unique index realms_id_uindex
    on realms (id);

create unique index realms_name_uindex
    on realms (name);

create table users
(
    id    serial
        constraint users_pk
            primary key,
    email text not null
);

alter table users
    owner to postgres;

create unique index users_id_uindex
    on users (id);

create table realm_organizations
(
    id              serial
        constraint realm_organizations_pk
            primary key,
    realm_id        integer not null
        constraint realm_organizations_realms_id_fk
            references realms
            on update restrict on delete restrict,
    organization_id integer not null
        constraint realm_organizations_organizations_id_fk
            references organizations
            on update restrict on delete restrict
);

alter table realm_organizations
    owner to postgres;

create unique index realm_organizations_id_uindex
    on realm_organizations (id);

create table organizations
(
    id   serial
        constraint organizations_pk
            primary key,
    name text not null
);

alter table organizations
    owner to postgres;

create unique index organizations_id_uindex
    on organizations (id);

create table organization_users
(
    id              serial
        constraint organization_users_pk
            primary key,
    organization_id integer not null
        constraint organization_users_organizations_id_fk
            references organizations
            on update restrict on delete restrict,
    user_id         integer not null
        constraint organization_users_users_id_fk
            references users
            on update restrict on delete restrict
);

alter table organization_users
    owner to postgres;

create unique index organization_users_id_uindex
    on organization_users (id);

create table clients
(
    id            bigserial
        constraint clients_pk
            primary key,
    realm_id      integer                         not null
        constraint clients_realms_id_fk
            references realms
            on update restrict on delete restrict,
    client_secret text                            not null,
    client_id     text                            not null,
    redirect_urls text[] default ARRAY []::text[] not null
);

alter table clients
    owner to postgres;

create unique index clients_client_id_uindex
    on clients (client_id);

create unique index clients_client_secret_uindex
    on clients (client_secret);



