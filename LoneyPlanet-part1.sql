drop database if exists `mini_project`;

create database `mini_project`;

use `mini_project`;

create table user_info
(
	sid varchar(10) not null,
	nike_name varchar(20) null,
	college varchar(20) null,
	gender varchar(2) null,
	grade varchar(5) null,
    portrait int null,
	constraint sid_pk
		primary key (sid)
)DEFAULT CHARSET=UTF8MB4;

create table requirements
(
	requirement_id int auto_increment,
	sender_sid varchar(10) null,
	title varchar(50) null,
	content varchar(200) null,
	post_time varchar(10) null,
	date int null comment"需求时间区间",
	time_from tinyint null,
	time_end tinyint null,
	require_people_num int null,
	place tinyint null,
	tag tinyint null comment"第二级标签",
	type tinyint null comment"第一级标签",
	contact_way_type varchar(10) null,
	contact_way varchar(20) null,
    status tinyint default 1 not null,
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

create table application
(
	application_id int auto_increment,
	receiver_sid varchar(10) null,
	sender_sid varchar(10) null,
	requirements_id int null,
	confirm tinyint default 1 not null comment"1:未确认,2:接受,3:拒绝",
    contact_way_type varchar(10) null,
    contact_way varchar(20) null,
    receiver_read_status tinyint default 1 not null comment"1:未读,2:已读",
    sender_read_status tinyint default 1 not null,
    send_time varchar(10) null comment"发送申请的时间",
    confirm_time varchar(10) null comment"确认申请的时间",
    title varchar(50),
	constraint applications_pk
		primary key (application_id)
)DEFAULT CHARSET=UTF8MB4;

create index application_id_index on applications(application_id);
create index receiver_sid_index on applications(receiver_sid);
create index sender_sid_index on applications(sender_sid);
create index requirements_id_index on applications(requirements_id);

create table latest_action
(
	sid varchar(10) not null,
	latest_time varchar(10) null,
	rand_num tinyint null,
	constraint latest_action_pk
		primary key (sid)
)DEFAULT CHARSET=UTF8MB4;



