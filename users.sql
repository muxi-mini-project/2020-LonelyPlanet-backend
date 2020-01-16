drop database if exists `mini_project`;

create database `mini_project`;

use `mini_project`;

create table users
(
	sid varchar(50) not null,
	nike_name varchar(50) null,
	college varchar(50) null,
	gender varchar(20) null,
	grade varchar(20) null,
	requirements int null,	debunks int null,
    night_portrait int null,
	constraint sid_pk
		primary key (sid)
)DEFAULT CHARSET=UTF8MB4;



create table debunks
(
	debunk_id int auto_increment,
	sender_sid varchar(50) null comment"发送者id",
	content int null,
	post_time varchar(50) null,
	constraint debunks_pk
		primary key (debunk_id)
)DEFAULT CHARSET=UTF8MB4;

create index sender_sid_index on debunks(sender_sid);

create table requirements
(
	requirement_id int auto_increment,
	sender_sid varchar(50) null,
	title varchar(50) null,
	content varchar(200) null,
	post_time varchar(50) null,
	date varchar(50) null comment"需求时间区间",
	time_from tinyint null,
	time_end tinyint null,
	require_people_num int null,
	place tinyint null,
	tag tinyint null comment"第二级标签",
	type tinyint null comment"第一级标签",
	contact_way_type tinyint null,
	contact_way varchar(50) null,
	constraint requirements_pk
		primary key (requirement_id)
)DEFAULT CHARSET=UTF8MB4;

create index sender_sid_index on requirements(sender_sid);
create index date_index on requirements(date);
create index time_from_index on requirements(time_from);
create index time_end_index on requirements(time_end);
create index place_index on requirements(place);
create index tag_index on requirements(tag);
create index type_index on requirements(type);

create table night_comments
(
	commend_id int auto_increment,
	debunk_id int null,
	comment_time varchar(50) null,
	content varchar(200) null,
	constraint night_comments_pk
		primary key (commend_id)
)DEFAULT CHARSET=UTF8MB4;

create index debunk_id_index on night_comments(debunk_id);

create table reminders
(
	remind_sid int auto_increment,
	receiver_sid varchar(100) null,
	type tinyint null comment"黑天提醒或者白天提醒",
	constraint reminders_pk
		primary key (remind_sid)
);

create index receiver_sid_index on reminders(receiver_sid); 
create index type_index on reminders(type);

create table day_applications
(
	application_id int not null,
	receiver_sid int null,
	sender_sid int null,
	requirements_id int null,
	confirm tinyint null,
	constraint day_applications_pk
		primary key (application_id)
);

create index application_id_index on day_applications(application_id);
create index receiver_sid_index on day_applications(receiver_sid);
create index sender_sid_index on day_applications(sender_sid);
create index requirements_id_index on day_applications(requirements_id);



