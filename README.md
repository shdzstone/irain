### 项目简介
项目后台采用golang/gin搭建后台服务，后台管理采用node+vue3+antdui，前端页面面采用eact+umi+dva实现。
功能上简单实现后台的博客标签及博文管理。
本项目布署使用docker-compose进行容器化布署，相关的mysql+golang+node+redis+gogs+mysql+nginx服务
均采用docker-compose+Dockerfile方式进行容器构建。

### 后端架构
整体上采用gin做为http服务器，详细框架略

### 项目布署
1. 安装docker/docker-compose

2. 使用docker-compose相关容器
在irain项目根目录运行如下命令：
```
docker-compose up -d
```
3. mysql8 root密码加密方案去除
```
>docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED              STATUS              PORTS                                            NAMES
c3dc74b06bec        gogs/gogs           "/app/gogs/docker/st…"   About a minute ago   Up About a minute   0.0.0.0:10022->22/tcp, 0.0.0.0:3006->3000/tcp    irain_gogs_1
beb546129605        irain_mysql         "docker-entrypoint.s…"   About a minute ago   Up About a minute   0.0.0.0:3306->3306/tcp, 33060/tcp                irain_mysql_1
>docker exec -it bbdd42acfad1 bash
 root@bbdd42acfad1:/# mysql -uroot -p
 Enter password:
 Welcome to the MySQL monitor.  Commands end with ; or \g.
 Your MySQL connection id is 8
 Server version: 8.0.22 MySQL Community Server - GPL
 
 Copyright (c) 2000, 2020, Oracle and/or its affiliates. All rights reserved.
 
 Oracle is a registered trademark of Oracle Corporation and/or its
 affiliates. Other names may be trademarks of their respective
 owners.
 
 Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
 
 mysql> show databases;
 +--------------------+
 | Database           |
 +--------------------+
 | gogs               |
 | information_schema |
 | mysql              |
 | performance_schema |
 | sys                |
 +--------------------+
 5 rows in set (0.00 sec)
 
 mysql> ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';
 Query OK, 0 rows affected (0.01 sec)

```