create table tasks (
    id int(11) not null auto_increment,
    name varchar(255) default null,
    created_at datetime default null,
    updated_at datetime default null,
    del_flg smallint default 0,
    primary key(id)
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
), (
    (select max(id) from tasks as t)+1,
    "読書",
    NOW(),
    NOW(),
    0
);
