# tut-grpc-go-web

- mysql
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [go-playground/validator](github.com/go-playground/validator)
- [go-xorm/xorm](https://github.com/go-xorm/xorm)
- [onsi/ginkgo](github.com/onsi/ginkgo/ginkgo)
- [DATA-DOG/go-sqlmock](https://github.com/DATA-DOG/go-sqlmock)
- [Mockgen](https://github.com/golang/mock)

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

## Test

ginkgo

```shell
# cd to the folder you want to have testing
cd repos

# to init test file
ginkgo bootstrap

# to test one
ginkgo -v -failFast --focus="UsersRepo"

# to test all
ginkgo -r -failFast
```

mockgen

```shell
# install mockgen
go get github.com/golang/mock/mockgen
mockgen --version

# create mock file
touch repos/generateMock.sh
# change sh file restrictions
chmod +x repos/generateMock.sh
# edit the sh file
# run the sh file to generate mock files of the repo
repos/generateMock.sh
```
