create table users
(
    user_id            integer not null
        primary key autoincrement,
    user_name          text    not null,
    count_of_questions integer not null
);

create table questions
(
    question_id   INTEGER not null
        primary key  autoincrement,
    user_id       INTEGER not null
        references users,
    question_text TEXT    not null,
    answer_text   TEXT    not null
);
