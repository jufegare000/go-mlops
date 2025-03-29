Para crear la base de datos, es necesario tener instalado alguna herramienta de contenerización, en este caso se usará podman:


```
podman run \
    -itd \
    --name mysql \
    -e MYSQL_ROOT_PASSWORD=root \
    -e MYSQL_USER=user \
    -e MYSQL_PASSWORD=secret \
    -e MYSQL_DATABASE=literary_works \
    -p 3306:3306 \
    docker.io/library/mysql
```

Conexión a base de datos: jdbc:mysql://localhost:3306/literary_works?allowPublicKeyRetrieval=true&useSSL=false

Creación tabla de autor:
```
create table author(
	id int primary key NOT NULL AUTO_INCREMENT,
	author_code varchar(5),
	name varchar(128),
	birth_date date, 
	death_date date,
	nationality varchar(64)
);
```

Algunos registros para la tabla author
```
insert into author(`name`, `birth_date`, `death_date`, `nationality`) values('Marcus Aurelius Antoninus', '121-04-26', '180-03-17', 'Rome');
```

Creación de la tabla literary work
```
create table literary_work(
	id int primary key NOT NULL AUTO_INCREMENT,
	author_id int,
	name varchar(128),
	literary_work_code varchar(5), 
	publication_date date,
	FOREIGN KEY (author_id) REFERENCES author(id)
);
```

Insertar una obra lieraria para marco aurelio

```
insert into literary_work(`author_id`, `name`, `publication_date`) 
select id, 'Meditations','180-01-01' from author
where author_code = 'MAAR';
```

Creación de la tabla quotes

create table famous_quotes(
	id int primary key NOT NULL AUTO_INCREMENT,
	literary_work_id int,
	description varchar(2048),
	FOREIGN KEY (literary_work_id) REFERENCES literary_work(id)
);

Algunas quotes

```
insert into famous_quotes(`literary_work_id`, `description`) 
select id, 'You have power over your mind - not outside events. Realize this, and you will find strength.' from literary_work 
where literary_work_code = 'MMAAR';
```