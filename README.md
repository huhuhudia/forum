# forum
毕业设计



users

```sql
uid ( primary key ) 
username string
password string
avatar string
sex bool
age int
email string
phone string

create table users(
    uid integer auto_increment primary key,
    username varchar(50),
    password varchar(50),
    avatar varchar(50),
    sex boolean,
    age integer,
    email varchar(50),
    phone varchar(50)
)

```



relations

```sql
rid (primary key int) 
uid_star int
uid_fan int

create table relations(
	rid integer auto_increment primary key,
    uid_star integer,
    uid_fan integer
)
```

topics

```sql
tid (primary key int)
author_id  
title string
content string
up int
down int


create table topics(
	tid integer auto_increment primary key,
	author_id integer,
	title varchar(150),
	content varchar(500),
	up integer,
	down integer
)
```



comments

```sql
cid  (primary key int ) 
tid int

author_id int
content string
up int 
down int

create table comments(
	cid integer auto_increment primary key,
    tid integer,
    author_id integer,
    content varchar(500),
    up integer,
    down integer
)
```

