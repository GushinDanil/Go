# Go
Программа запускается из файла  main.go
Приложение для магазина, в котором пользователь сможет:

1)Зарегистрироваться 

3)Редактировать его аккаунт

4)Удалить его аккаунт

5)Создать продукт и выставить цену на него

6)Редактировать свою запись(продукта)

7)Просмотреть все предложения на продукты которые есть в базе данных

8)Просмотр определенного продукта(по id)

9)Просмотр других предложений в магазине, опубликованных другими пользователями 
  
10)Удалить созданный им продукт


Примечание: для методов типа Get наличие ключа токена не обязательно.
Удалить или изменить продукт  и аккаунт пользователя может только пользователь который его создал.



Перво наперво нужно создать пользователя (/createUser) 
отправив запрос такого формата:

{

    "nickname":"Nick",

    "email":"danilgusin17@gmail.com",

    "password": "password"

}

далее  пройти по пути (login/) для авторизации и получения токена  отправив запрос такого формата:

{

    "email":"danilgusin17@gmail.com",

    "password": "password"

}


В дальнейших операциях использовать
его в параметре Authorization(Bearer Token)

Таблицы из базы данных:


    create table users

    (

    id       serial not null

        constraint users_pkey

            primary key,

    nickname text,

    email    text,

    password text
    );

    alter table users

    owner to postgres;

    create unique index users_email_uindex

    on users (email);

    create unique index users_nickname_uindex

    on users (nickname)

    create table products

    (

     id        serial not null

        constraint products_pkey1

            primary key,

    name      text,

    price     integer,

    user_id   integer,

    createdat time,

    updatedat time
    );
