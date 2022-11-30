DROP SCHEMA public CASCADE;
CREATE SCHEMA public;



CREATE
    TABLE person (
                     "id" serial NOT NULL,
                     "first_name" VARCHAR(255) NOT NULL,
                     "last_name" VARCHAR(255) NOT NULL,
                     "patronymic" VARCHAR(255) NULL,
                     "login" VARCHAR(255) NOT NULL unique ,
                     "role_id" integer NOT NULL,
                     "passwd" VARCHAR(255) NOT NULL,
                     "updated_at" TIMESTAMP NOT NULL default current_timestamp,
                     "deleted" BOOLEAN NOT NULL default false,
                     CONSTRAINT "person_pk" PRIMARY KEY ("id")
);


INSERT
INTO person
(first_name, last_name, patronymic, login, role_id, passwd) VALUES
    ('Ivan', 'Ivanov', 'Ivanovich', 'Ivan777', 3, '12345678');


DROP TABLE person CASCADE;