数据库设计过程：
---
用户表设计：

1. 整体登录表分为两个部分，个人表和用户表，现在暂时预留个人表作为登录用户。
2. 数据库为容器启动，将容器中的数据暂存服务器端。

Docker数据库远程连接踩坑：

1. 阿里云服务器有自己的防火墙，导致远程不能连接服务器。
2. 给远程连接增加权限 

```
GRANT ALL PRIVILEGES ON *.* TO 'your name'@'%' IDENTIFIED BY 'your password' WITH GRANT OPTION;

FLUSH PRIVILEGES;
```