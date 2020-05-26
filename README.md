# MySQL实时同步

#### 场景:

##### 开启MySQL binlog日志, 解决主从复制中MySQL版本不一致需要同步数据的要求, 可以高版本MySQL同步低版本MySQL的数据, 也可以低版本MySQL同步高版本MySQL的数据, 但是不能实时同步表结构, 表结构修改需要在被同步方执行DDL语句, 支持位点重置或者跳过位点;



#### 使用说明

##### 1.将源MySQL做一次全量备份, 记住备份时的binlog文件和位点, 导入目标数据库;

##### 2.修改配置文件 config/config.properties 相关配置项, 需要修改MySQL源地址和目标地址及用户密码, 以及源数据库的binlog文件和位点;

##### 3.运行main.go文件, 实现源MySQL到目标MySQL的数据同步;

##### 4.执行windows部署文件deploy.bat, 在项目目录下生产main文件;

##### 5.将main文件传到Linux下, 依赖的文件夹, config/, data/, doc/, logs/ 一并放在main目录下, 这里需要编辑 config/config.properties 为真实环境配置, 授予main执行权限, 然后 ./main 运行项目.

##### 目的: 为DBA开发一套跨不同版本MySQL实时同步的系统, 后期还会不断扩充功能, 有问题请加微信沟通, 或者直接留言.

![image](https://github.com/xionglihdfs/BusinessMailTimedTask/blob/master/doc/%E7%86%8A%E7%86%8A.png)