## 简介


#### 工具
1、界面框架layUI2.4.5
2、makedown.md
3、beego1.8 4、Ztree


#### 部署
```
1. 准备环境
git clone https://github.com/Ohimma/beegoSre.git
go get github.com/astaxie/beego
go get github.com/beego/bee
go get github.com/astaxie/beego 加入到 ~/.bashrc或~/.bash_profile文件中

2. 数据库准备
CREATE DATABASE beego DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci
set global validate_password.policy=0;
set global validate_password.length=4;
create user 'beego_sre'@'%' identified by 'beego_sre@xxx';

mysql -ubeego_sre -p"beego_sre@xxx" -P3306 -h 140.xxxxx beego  < beego_sre.sql


3. 启动
bee run

4. 默认用户名密码：
admin george518

```

