-- processing systems table
create table if not exists processing_systems
(
    id uuid not null
        constraint processing_systems_pk
            primary key DEFAULT (uuid_generate_v4()),
    code varchar(50) not null,
    name varchar(200),
    is_pay_allowed boolean   default false             not null,
    created_at timestamp with time zone DEFAULT (now()),
    updated_at     timestamp,
    deleted_at     timestamp
);

comment on table processing_systems is 'Справочник ПЦ';

create unique index processing_systems_code_uindex
    on processing_systems (code);

-- card bins table
create table card_bins
(
    id uuid not null
        constraint card_bins_pk
            primary key DEFAULT (uuid_generate_v4()),
    name           varchar(150)                        not null,
    processing_id  uuid  not null constraint processing_systems_fk
        references processing_systems(id),
    number         varchar(50)                         not null,
    currency_code  varchar(10)                         not null,
    is_bank_card   boolean   default false             not null,
    is_add_allowed boolean   default false             not null,
    created_at timestamp with time zone DEFAULT (now()),
    updated_at     timestamp,
    deleted_at     timestamp
);

alter table card_bins
    owner to postgres;

create unique index card_bins_bin_uindex
    on card_bins (number);