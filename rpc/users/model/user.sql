create table user_tab
(
    id       bigint auto_increment,
    username varchar(128) not null,
    password varchar(64)  not null,
    ctime    int          not null,
    mtime    int          not null,
    constraint user_tab_pk
        primary key (id)
);