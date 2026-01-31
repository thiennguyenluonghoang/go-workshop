docker run --name postgre16 -p 5450:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=daithuvien -d postgres:16-alpine


docker run: Tạo mới & khởi động 
--name postgre16: Đặt tên cho container này là postgre16

-p 5432:5432: Port mapping -p [Cổng máy thật(host)]:[Cổng trong container] => Nối cổng 5432 trên lap của mình với 5432 của cái container đó => Giúp DBeaver, code Go/python có thể kết nối vô db thông qua địa chỉ localhost:5432

-e POSTGRES_USER=root: Biến môi trường để tạo tài khoản user mặc định 

-e POSTGRES_PASSWORD=password: Mật khẩu cho user trên 

-d (Detached mode): Chạy ngầm, sau khi chạy lên xong thì terminal sẽ trả lại quyền điều khiển cho mình, chứ nếu ko có cờ này thì terminal sẽ bị kẹt ở màn log của DB, tắt terminal là tắt DB

postgres:16-alpine (Image): Là bản thiết kế để tạo container, gồm tên image -> postgres, phiên bản postgreSQL: 16, alpine là bản siêu nhẹ 

----------------------------------------
Cấu hình để kết nối vô db
Host: localhost
Port: 5432
User: root
Password: password
Database name: root

----------------------------------------

docker ps: Xem danh sách các container đang hoạt động

Truy cập vào chế độ tương tác với db: docker exec -it postgre16 /bin/bash

Truy cập psql: psql

root=# show dbs
root=# \l

root=# create database testdb
root=# drop database testdb;
root=# exit


Lệnh tạo db từ ngoài: docker exec -it postgre16 createdb --username=root --owner=root testdb

Lệnh xóa db từ ngoài: docker exec -it postgre16 dropdb testdb

---------------------------------


create table users
(
    id                serial primary key,
    user_name          character(50),
    email           character(50),
    encrypted_password character(50),
    create_at timestamptz,
    updated_at timestamptz,
    create_by character(50),
    updated_by character(50),
    delete_flag bool 
);
