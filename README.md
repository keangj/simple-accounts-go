# go

## 运行 docker 数据库

``` sh
docker run -d --name pg-go-simple-accounts -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=simple_accounts_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-simple-accounts-data:/var/lib/postgresql/data postgres:14
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
docker exec -it pg-go-simple-accounts bash # 进入 docker psql 数据库
psql -U <user name> -d <database name> # 连接数据库
\l # 查看全部数据库
\c <database name> # 连接数据库
\d # 查看全部表
\d <table name> # 查看表
\q # 退出数据库，或者 control + d
```
