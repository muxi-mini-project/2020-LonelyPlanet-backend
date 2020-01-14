drop database if exists `mini_project`

create database `mini_project`

use `mini_project`

create table users
(
	sid varchar(50) not null,
	nike_name varchar(50) null,
	college varchar(50) null,
	gender int null,
	grade varchar(20) null,
	requirements int null,
	debunks int null,
	constraint sid_pk
		primary key (sid)
)DEFAULT CHARSET=UTF8MB4;

create table debunks
(
	debunk_id int auto_increment,
	sender_sid varchar(50) null,
	title varchar(100) null,
	content int null,
	post_time varchar(50) null,
	constraint debunks_pk
		primary key (debunk_id)
)DEFAULT CHARSET=UTF8MB4;

create table requirements
(
	requirement_id int auto_increment,
	sender_sid varchar(50) null,
	title varchar(50) null,
	content varchar(200) null,
	post_time varchar(50) null,
	date varchar(50) null,
	time_form tinyint null,
	time_end tinyint null,
	require_people_num int null,
	place tinyint null,
	tag tinyint null,
	type tinyint null,
	contact_way_type tinyint null,
	contact_way varchar(50) null,
	constraint requirements_pk
		primary key (requirement_id)
)DEFAULT CHARSET=UTF8MB4;

create table night_comments
(
	commend_id int auto_increment,
	debunk_id int null,
	comment_time varchar(50) null,
	content varchar(200) null,
	colour varchar(50) null,
	constraint night_comments_pk
		primary key (commend_id)
)DEFAULT CHARSET=UTF8MB4;

create table reminders
(
	remind_id int auto_increment,
	receiver varchar(100) null,
	type tinyint null,
	constraint reminders_pk
		primary key (remind_id)
);

create table day_applications
(
	application_id int not null,
	receiver_id int null,
	sender_id int null,
	requirements_id int null,
	confirm tinyint null,
	constraint day_applications_pk
		primary key (application_id)
);




