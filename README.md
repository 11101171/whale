# whale
a revel framework web,
+ 1.Provide SSH login and custom script by youseft;
+ 2.Online API testing; 

#database
mysql

#data
+ Import Sql folder conf: data.json,  mysql -u username -pxxx -D gorp < data.sql
+ Config app.conf of mysql account
+ After that, there is an account for you to use,  account name is 'super', password '123456'

#deploy
> ##run
> + 模式dev:   sh run-dev-9999.sh start 
> + 模式prod:  sh run-prod-9988.sh start

> ##stop
> sh run-dev-9999.sh stop
