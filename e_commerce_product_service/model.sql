

create table product (
    id uuid primary key,
    name varchar,
    image varchar,
    price numeric,
    discount_type varchar,
    discount numeric,
    barcode varchar
);
