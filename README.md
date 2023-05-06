# go

## 运行 docker 数据库

``` sh
# postgres
docker run -d --name pg-go-simple-accounts -p 5432:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=simple_accounts_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-simple-accounts-data:/var/lib/postgresql/data postgres:14
# mysql
docker run -d --name mysql-go-simple-accounts -p 3306:3306 - -e MYSQL_DATABASE=simple_accounts_dev -e MYSQL_USER=jay -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v mysql-go-simple-accounts-data:/var/lib/mysql mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

--name 容器名称。可以作为 ip 使用，需加 --network
-e env 环境变量
  POSTGRES_USER 用户名
  POSTGRES_PASSWORD 密码
  POSTGRES_DB 数据库名
  PGDATA=/var/lib/postgresql/data/pgdata 根据 pg docker 文档配置
-v 数据持久化映射

## 进入数据库

``` sh
# pgsql
docker exec -it pg-go-simple-accounts bash # 进入 docker psql 数据库
psql -U <user name> -d <database name> # 连接数据库
\l # 查看全部数据库
\c <database name> # 连接数据库
\d # 查看全部表
\d <table name> # 查看表
\q # 退出数据库，或者 control + d
# mysql
docker exec -it mysql-go-simple-accounts bash # 进入 docker mysql 数据库
psql -u <user name> -p <database name> # 连接数据库
show databases; # 查看全部数据库
use <database name> # 连接数据库
show tables; # 查看全部表
describe <table name> # 查看表
exit # 退出数据库，或者 control + d
```

## [sqlc](https://docs.sqlc.dev/)

- 安装

  ``` sh
  # mac
  brew install sqlc
  ```

### 用 [golang-migrate](https://github.com/golang-migrate/migrate) 数据迁移

- 安装 golang-migrate

  ``` sh
  brew install golang-migrate
  migrate --version # 4.x.x
  ```

- 创建迁移文件

``` sh
migrate create -ext sql -dir config/migrations -seq create_users_table
```

- 运行迁移文件

``` sh
# 升级
migrate -database "postgres://admin:123456@localhost:5432/simple_accounts_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up
# 降级
migrate -database "postgres://admin:123456@localhost:5432/simple_accounts_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" down 1
```
