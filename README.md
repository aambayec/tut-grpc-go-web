# tut-grpc-go-web

- mysql
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [go-playground-validator](github.com/go-playground/validator)

## Dependency Installation

### MYSQL

```shell
sudo apt install mysql-server
mysql -u root -p

# need to run this due to error 'Access denied for user 'root'@'localhost''
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'Ulyanin123';

create database tut_grpc;
```

### [golang-migrate](https://github.com/golang-migrate/migrate)

Linux Install

```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
sudo mv migrate.linux-amd64 /usr/bin/
which migrate
```

Mac Install

```shell
brew install golang-migrate

```

Migration

```shell
# generate up and down file
migrate create -dir migrations -ext sql create_users_table

# modify the generated file to your liking

# run migrate up/down
migrate -path ./migrations -database "mysql://root:Ulyanin123@/tut_grpc" -verbose up
migrate -path ./migrations -database "mysql://root:Ulyanin123@/tut_grpc" -verbose down
```
