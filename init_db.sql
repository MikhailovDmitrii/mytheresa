create table products
(
    sku      varchar not null
        constraint promotions_pk
            primary key,
    name     varchar not null,
    price    integer not null,
    category varchar not null
);