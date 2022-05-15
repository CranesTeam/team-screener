create table user_roles (
	id serial not null unique,
	name varchar(10) not null,
	description varchar(50)	
);

create table users (
	id serial not null unique,
	external_uuid varchar(55) not null unique default uuid_generate_v4(),
	username varchar(25) not null unique,
	password_hash varchar(255) not null,
	role_id int references user_roles(id) on delete cascade not null
);
create unique index user_extrernal_uuid_idx on users(external_uuid);

create table user_info(
	id serial not null unique,
	external_uuid varchar(55) not null unique default uuid_generate_v4(),
	user_id int references users(id) on delete cascade not null,
	name varchar(55) not null,
	email varchar(55) unique not null
);

create table skills (
	id serial not null unique,
	external_uuid varchar(55) not null unique default uuid_generate_v4(),
	name varchar(55) not null,
	title varchar(55) not null,
	description varchar(255)	
);

create unique index skill_extrernal_uuid_idx on skills(external_uuid);
create unique index skill_name_idx on skills(name);

create table user_skills (
	id serial not null unique,
	external_uuid varchar(55) not null unique default uuid_generate_v4(),	
	user_uuid varchar(55) references users(external_uuid) on delete cascade not null,
	skill_uuid varchar(55) references skills(external_uuid) on delete cascade not null,
	points int not null
);

create unique index userSkill_extrernal_uuid_idx on user_skills(external_uuid);
