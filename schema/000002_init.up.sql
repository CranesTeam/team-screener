create table users (
	id serial not null unique,
	external_uuid string not null unique,
	name varchar(55) not null,
	username varchar(25) not null unique,
	password_hash vachar(255) not null
);

create unique index user_extrernal_uuid_idx on users(external_uuid);

create table skills (
	id serial not null unique,
	external_uuid string not null unique,
	name varchar(55) not null,
	title varchar(55) not null,
	description vachar(255)	
);

create unique index skill_extrernal_uuid_idx on skill(external_uuid);
create unique index skill_name_idx on skill(name);

insert into skill (external_uuid, name, title, description) values 
(uuid_generate_v4(), 'Java', 'Java', 'Java 8-16 version'),
(uuid_generate_v4(), 'Kotlin', 'Kotlin', 'Kotlin as main language'),
(uuid_generate_v4(), 'Kotlin Coroutines', 'Kotlin Coroutines', 'Multithreading with Kotlin'),
(uuid_generate_v4(), 'Spring boot', 'Spring boot', 'Spring boot version more 2.5.*'),
(uuid_generate_v4(), 'Spring', 'Spring', 'Spring Framework'),
(uuid_generate_v4(), 'Spring MVC', 'Spring MVC', 'Spring MVC'),
(uuid_generate_v4(), 'Spring Secutiry', 'Spring Secutiry', 'Spring Security implementation'),
(uuid_generate_v4(), 'Spring Actuator', 'Spring Actuator', 'Spring health check'),
(uuid_generate_v4(), 'Spring JPA', 'Spring JPA', 'Spring JPA'),
(uuid_generate_v4(), 'Hibernate', 'Hibernate', 'Spring JPA over Hibernate'),
(uuid_generate_v4(), 'JOOQ', 'JOOQ', 'JOOQ data layer'),
(uuid_generate_v4(), 'Prometheus', 'Prometheus', 'Prometheus metrics'),
(uuid_generate_v4(), 'Micronaut', 'Micronaut', 'Micronaut framework'),
(uuid_generate_v4(), 'ELK', 'ELK', 'Logging with ELK stack'),
(uuid_generate_v4(), 'Grafana', 'Grafana', 'Grafana metrics and alerts'),
(uuid_generate_v4(), 'Golang', 'Golang', 'Golang computer language'),
(uuid_generate_v4(), 'PHP', 'PHP', 'PHP computer language'),
(uuid_generate_v4(), 'Ktor', 'Ktor', 'Ktor Framework'),
(uuid_generate_v4(), 'Docker', 'Docker', 'Docker engine and UI'),
(uuid_generate_v4(), 'Kubernates', 'Kubernates', 'k8s'),
(uuid_generate_v4(), 'Kubernates dashboard', 'Kubernates dashboard', 'k8s ui');

create table userSkills (
	id serial not null unique,
	external_uuid string not null unique,	
	user_id int references users(id) on delete cascade not null,
	skill_id int references skills(id) on delete cascade not null,
	points int not null
);

create unique index userSkill_extrernal_uuid_idx on userSkills(external_uuid);
