# Go-Database
go 语言链接操作各种类型的数据库。



link1-redis:https://blog.csdn.net/wangshubo1989/article/details/75050024

link2-pgsql:https://blog.csdn.net/wangshubo1989/article/details/77478838
### Go 链接redis  

#### docker 

docker run --name redis-test -p 6379:6379 -d --restart=always redis:latest redis-server --appendonly yes --requirepass "your passwd"

docker exec -it a126ec987cfe redis-cli -a 'your passwd'


### Go pgsql

docker run --name postgres1 -e POSTGRES_PASSWORD=password -p 5432:5432 -d --restart=always postgres:9.4