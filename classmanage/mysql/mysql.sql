create table overall (
    id int unsigned PRIMARY key auto_increment comment "ID",
    classid int unsigned DEFAULT 0 comment "课程ID",
    classroomid int unsigned DEFAULT 0 comment "教室ID",
    teacherid int unsigned DEFAULT 0 comment "教师ID",
    update_at TIMESTAMP DEFAULT current_timestamp comment "修改时间"
);

create table class (
    id int unsigned PRIMARY key auto_increment comment "ID",
    className char(64) DEFAULT "" comment "课程名字",
    classDesc VARCHAR(255) DEFAULT "" comment "课程描述",
    create_at TIMESTAMP DEFAULT "0000-00-00 00:00:00" comment "创建日期",
    update_at TIMESTAMP DEFAULT current_timestamp comment "修改时间"
);

create table classroom (
    id int unsigned PRIMARY key auto_increment comment "ID",
    classroomName char(64) DEFAULT "" comment "教室名字",
    classAdress VARCHAR(255) DEFAULT "" comment "教室地址",
    create_at TIMESTAMP DEFAULT "0000-00-00 00:00:00" comment "创建日期",
    update_at TIMESTAMP DEFAULT current_timestamp comment "修改时间"
);

create table teacher (
    id int unsigned PRIMARY key auto_increment comment "ID",
    teacherName char(64) DEFAULT "" comment "教师名字",
    gender enum("select","male","female","secret") DEFAULT "secret" comment "性别",
    create_at TIMESTAMP DEFAULT "0000-00-00 00:00:00" comment "创建日期",
    update_at TIMESTAMP DEFAULT current_timestamp comment "修改时间"
);


create table admin (
    id int unsigned PRIMARY key auto_increment comment "ID",
    user varchar(255) DEFAULT "" comment "账号",
    password varchar(255) DEFAULT "" comment "密码",
    name varchar(255) DEFAULT "" comment "昵称",
    create_at TIMESTAMP DEFAULT "0000-00-00 00:00:00" comment "创建日期",
    update_at TIMESTAMP DEFAULT current_timestamp comment "修改时间"
);


INSERT into overall  VALUES (null,1,1,1,now());

INSERT into class VALUES (null,"百步飞剑班","学习百步飞剑",now(),now());

INSERT into classroom VALUES (null,"剑术教室","大秦",now(),now());

INSERT into teacher VALUES (null,"大叔","male",now(),now());

insert into admin values (null,"777105543",MD5("ld123456"),"赤城赤城我是吃撑",now(),now());