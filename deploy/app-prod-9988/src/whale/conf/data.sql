-- MySQL dump 10.13  Distrib 5.6.24, for osx10.8 (x86_64)
--
-- Host: localhost    Database: gorp
-- ------------------------------------------------------
-- Server version	5.6.25

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Dumping data for table `o_agent`
--

LOCK TABLES `o_agent` WRITE;
/*!40000 ALTER TABLE `o_agent` DISABLE KEYS */;
INSERT INTO `o_agent` VALUES ('2016-03-19 20:28:39','2016-03-19 20:28:39','190338319319031908jebiymagent','120.26.78.245','22','souche','','cd /home','好机器111','190227219219021907-user-bzqmlj'),('2016-03-20 18:06:26','2016-03-20 18:06:26','200325320320032006bcdebmagent','42.96.134.15','21','souche','','cd /home','test','190227219219021907-user-bzqmlj'),('2016-03-20 20:01:57','2016-03-20 20:01:57','200357320320032008pjczsfagent','112.124.18.98','22','souche','','cd /home','weijin','190227219219021907-user-bzqmlj');
/*!40000 ALTER TABLE `o_agent` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `o_cmd`
--

LOCK TABLES `o_cmd` WRITE;
/*!40000 ALTER TABLE `o_cmd` DISABLE KEYS */;
INSERT INTO `o_cmd` VALUES ('2016-03-19 21:59:24','2016-03-19 21:59:24','190324319319031909smhkixcmd','csdcdscsdc sddfdsf','190338319319031908jebiymagent'),('2016-03-19 21:36:08','2016-03-19 21:36:08','19038319319031909oxitebcmd','cd /home','190338319319031908jebiymagent'),('2016-03-20 20:10:10','2016-03-20 20:10:10','200310320320032008vkuckccmd','hehhe','200325320320032006bcdebmagent'),('2016-03-20 18:01:15','2016-03-20 18:01:15','200315320320032006nkqeqycmd','df -hl','190338319319031908jebiymagent'),('2016-03-20 20:10:15','2016-03-20 20:10:15','200315320320032008pxsxhvcmd','aa','200325320320032006bcdebmagent'),('2016-03-20 18:01:21','2016-03-20 18:01:21','200321320320032006wljcstcmd','look','190338319319031908jebiymagent'),('2016-03-20 18:01:25','2016-03-20 18:01:25','200325320320032006nabgaacmd','sssss','190338319319031908jebiymagent'),('2016-03-20 18:01:29','2016-03-20 18:01:29','200329320320032006dudbdbcmd','dcdcdc','190338319319031908jebiymagent');
/*!40000 ALTER TABLE `o_cmd` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `p_server`
--

LOCK TABLES `p_server` WRITE;
/*!40000 ALTER TABLE `p_server` DISABLE KEYS */;
INSERT INTO `p_server` VALUES ('2016-03-11 17:27:26','2016-03-11 17:27:26','110325311311031105-server-bfhizo','190227219219021907-user-bzqmlj','1231QQQ','{\"host\":\"http://xx.xx.sqaproxy.souche.com\",\r\n   \"port\":\"80\",     \"name\":\"rrrapi\",     \"lists\":[\r\n          {\r\n              \"folder\":\"css\",\"sort\":\"1\",\"api_params\":[\r\n                  {\r\n                      \"id\":\"11\",\r\n                      \"path\":\"/admin/server/encode\",\r\n                      \"name\":\"测试\",\r\n                      \"method\":\"GET\",\r\n                      \"fileds\":[\r\n                              {\"name\":\"username\", \"ftype\":\"text\", \"value\":\"zhangsan\",\"lable\":\"用户名(不能小于3个字符)\",\"placeholder\":\"如:张三\"},\r\n                               {\"name\":\"uang\", \"ftype\":\"text\", \"value\":\"\",\"lable\":\"xx用户名(不能小于3个字符)\",\"placeholder\":\"如:你妈妈的\"},\r\n                              {\"name\":\"passwrod\", \"ftype\":\"text\", \"value\":\"\",\"lable\":\"密码\",\"placeholder\":\"如:abc123\"},\r\n                              {\"name\":\"sign\", \"ftype\":\"text\", \"value\":\"\",\"lable\":\"秘钥\",\"placeholder\":\"按加密按钮即可\", \"salt\":\"xxs\", \"ptype\":\"MD5\",\"pway\":\"1\"}\r\n                       ]\r\n                  },\r\n                   {\r\n                      \"id\":\"12\",\r\n                      \"path\":\"/v\",\r\n                      \"name\":\"xx认证接口\",\r\n                      \"method\":\"POST\",\r\n                      \"fileds\":[\r\n                              {\"name\":\"bankCard\", \"ftype\":\"text\", \"value\":\"6226197900918805\",\"lable\":\"银行卡\",\"placeholder\":\"如:6226197900918805\"},  {\"name\":\"name\", \"ftype\":\"text\", \"value\":\"邢希红\",\"lable\":\"用户名(不能小于1个字符)\",\"placeholder\":\"如:张三\"}, {\"name\":\"phone\", \"ftype\":\"text\", \"value\":\"18805391696\",\"lable\":\"预存手机号\",\"placeholder\":\"如:153000000\"}\r\n                       ]\r\n                  },\r\n                  {\r\n                      \"id\":\"13\",\r\n                      \"path\":\"/pwd/v1/api\",\r\n                      \"name\":\"验证xxxx接口\",\r\n                      \"method\":\"POST\",\r\n                      \"fileds\":[\r\n                              {\"name\":\"Authorization\", \"ftype\":\"text\", \"value\":\"SOUCHE=6c73b86978f95d9551f0fdb2fc55ee1f\",\"lable\":\"服务名\",\"placeholder\":\"\",\"location\":\"header\"},\r\n                              {\"name\":\"service\", \"ftype\":\"text\", \"value\":\"verify\",\"lable\":\"服务名\",\"placeholder\":\"\"},\r\n                               {\"name\":\"password\", \"ftype\":\"text\", \"value\":\"123456\",\"lable\":\"密码\",\"placeholder\":\"如:张三\"},\r\n                              {\"name\":\"sign\", \"ftype\":\"text\", \"value\":\"\",\"lable\":\"秘钥\",\"placeholder\":\"按加密按钮即可\", \"salt\":\"keyvalue\", \"ptype\":\"MD5\",\"pway\":\"1\"}\r\n                       ]\r\n                  }\r\n              ]\r\n          },\r\n          {\r\n              \"folder\":\"xxxxx\",\"sort\":\"1\",\"api_params\":[]\r\n          },\r\n          {\r\n              \"folder\":\"yyyyyyy\",\"sort\":\"1\",\"api_params\":[]\r\n          }\r\n      ]\r\n  }');
/*!40000 ALTER TABLE `p_server` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping data for table `s_user`
--

LOCK TABLES `s_user` WRITE;
/*!40000 ALTER TABLE `s_user` DISABLE KEYS */;
INSERT INTO `s_user` VALUES ('2016-02-19 19:03:27','2016-02-19 19:03:27','190227219219021907-user-bzqmlj','super','47a215710ca7b648251742453b1149db','本人','0');
/*!40000 ALTER TABLE `s_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-03-20 20:17:44
