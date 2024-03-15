package models

/**
create table user
(
    username varchar(128)             not null
        primary key,
    password varchar(512)  default '' not null,
    email    varchar(256)  default '' not null,
    image    varchar(1024) default '' not null,
    bio      varchar(1024) default '' not null,
    constraint email
        unique (email)
)
    charset = utf8mb4;
*/

type User struct {
	Username string
	Password string
	Email    string
	Image    string
	Bio      string
}
