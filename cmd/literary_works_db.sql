create
database literary_works

create table author
(
    id          int primary key NOT NULL AUTO_INCREMENT,
    name        varchar(128),
    birth_date  date,
    death_date  date,
    nationality varchar(64),
    author_code varchar(5)
);
insert into author(name, birth_date, death_date, nationality, author_code)
values ('Marcus Aurelius Antoninus', '0121-04-26',	'0180-03-17', 'Rome', 'MAAR');


create table literary_work
(
    id                 int primary key NOT NULL AUTO_INCREMENT,
    author_id          int,
    name               varchar(128),
    publication_date   date,
    literary_work_code varchar(5),
    FOREIGN KEY (author_id) REFERENCES author (id)
);


insert into literary_work(author_id, name, publication_date, literary_work_code)
values (1,	"Meditations",	"0180-01-01",	"MMAAR");


create table famous_quotes
(
    id               int primary key NOT NULL AUTO_INCREMENT,
    literary_work_id int,
    description      varchar(2048),
    FOREIGN KEY (literary_work_id) REFERENCES literary_work (id)
);

insert into famous_quotes(`literary_work_id`, `description`)
select id, 'You have power over your mind - not outside events. Realize this, and you will find strength.' from literary_work
where literary_work_code = 'MMAAR';
