create table users
(
    id     serial primary key,
    name   varchar(255) not null,
    status varchar(255) check ( status in ('ACTIVE', 'INACTIVE', 'BANNED') )
);

