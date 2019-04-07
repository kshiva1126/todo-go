create table user (
    id int(11) not null,
    name varchar(255) default null,
    old smallint default 0
);

insert into user values (
    1,
    'Adam',
    18
), (
    (select max(id) from user as u)+1,
    'Brown',
    20
), (
    (select max(id) from user as u)+1,
    'Choerry',
    17
);

create table tasks (
    id int(11) not null default 0,
    name varchar(255) default null,
    created_at datetime default null,
    updated_at datetime default null,
    del_flg smallint default 0
);

insert into tasks values (
    1,
    "洗濯",
    NOW(),
    NOW(),
    0
), (
    (select max(id) from tasks as t)+1,
    "掃除",
    NOW(),
    NOW(),
    0
), (
    (select max(id) from tasks as t)+1,
    "勉強",
    NOW(),
    NOW(),
    0
);
