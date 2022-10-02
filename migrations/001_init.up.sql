CREATE TABLE Person (
    person_id    serial primary key,
    first_name   VARCHAR(40) not null,
    last_name   VARCHAR(40),
    sur_name   VARCHAR(40)
);