
create table if not exists schools (
    school_id integer unsigned auto_increment primary key,
    name varchar(100) not null,
    created_at datetime
);

create table if not exists users (
    user_id integer unsigned auto_increment primary key,
    name varchar(100) not null,
    email varchar(100) not null unique,
    password varchar(100) not null,
    role varchar(100) not null,
    school_id integer unsigned not null,
    verified boolean,
    created_at datetime,
    foreign key (school_id) references schools(school_id)
);

create table if not exists verification_tokens (
    user_id integer unsigned not null unique,
    token varchar(100),
    expired_at datetime,
    foreign key (user_id) references users(user_id)
);

create table if not exists companys (
    company_id integer unsigned auto_increment not null primary key,
    name varchar(100) not null,
    created_at datetime
);

create table if not exists qrmarks (
    qrmark_id integer unsigned not null primary key,
    user_id integer unsigned not null,
    school_id integer unsigned not null,
    company_id integer unsigned not null,
    points integer unsigned not null,
    created_at datetime,
    foreign key (user_id) references users(user_id),
    foreign key (company_id) references companys(company_id),
    foreign key (school_id) references schools(school_id)
);

create table if not exists school_static_points (
    school_id integer unsigned not null,
    company_id integer unsigned not null,
    points integer unsigned not null,
    created_year_month datetime,
    foreign key (school_id) references schools(school_id),
    foreign key (company_id) references companys(company_id),
    unique key (school_id, company_id, created_year_month)
);


