/*!40101 SET NAMES utf8 */;
/*!40014 SET FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/ td /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE td;

DROP TABLE IF EXISTS contacts;
CREATE TABLE `contacts` (
  `account` varchar(45) NOT NULL,
  `contactid` int NOT NULL COMMENT 'id',
  `contactphone` varchar(45) DEFAULT NULL COMMENT '手机号码\n',
  `contactname` varchar(45) DEFAULT NULL COMMENT '联系人账号\n',
  `status` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`account`,`contactid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS groupinfos;
CREATE TABLE `groupinfos` (
  `phone` varchar(45) NOT NULL,
  `groupname` varchar(45) NOT NULL,
  `agent` int DEFAULT '0',
  `uid` int DEFAULT '0',
  PRIMARY KEY (`phone`,`groupname`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS groups;
CREATE TABLE `groups` (
  `uid` int NOT NULL AUTO_INCREMENT,
  `agent` int DEFAULT '0',
  `name` varchar(255) DEFAULT NULL,
  `linkurl` varchar(255) DEFAULT NULL,
  `verified` int DEFAULT '0' COMMENT '1 需要验证',
  PRIMARY KEY (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS phones;
CREATE TABLE `phones` (
  `phone` varchar(45) NOT NULL DEFAULT '手机号',
  `account` varchar(45) DEFAULT '账号',
  `tddata` varchar(255) DEFAULT NULL,
  `tdfile` varchar(255) DEFAULT NULL,
  `registered` int DEFAULT '0',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `agent` int DEFAULT NULL,
  PRIMARY KEY (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS taskinfo;
CREATE TABLE `taskinfo` (
  `tid` int NOT NULL AUTO_INCREMENT,
  `account` varchar(45) DEFAULT NULL,
  `groupid` int DEFAULT NULL,
  `groupname` varchar(45) DEFAULT NULL,
  `counts` int DEFAULT NULL COMMENT '循环次数',
  `cron` int DEFAULT NULL COMMENT '间隔次数秒',
  `cycle` int DEFAULT NULL COMMENT '1 循环 0 不循环',
  `text` varchar(1024) DEFAULT NULL COMMENT '发送内容',
  `createtime` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='任务信息';

DROP TABLE IF EXISTS tasklogs;
CREATE TABLE `tasklogs` (
  `tid` int NOT NULL,
  `counts` int DEFAULT NULL,
  `countsed` int DEFAULT NULL,
  `status` int DEFAULT NULL COMMENT ' 0 等等执行\n\n\nstatus 1 开始执行\n\nstatus 2 执行完毕\n\nstatus 4 取消\n',
  `operid` int DEFAULT NULL,
  `createtime` datetime DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='发送任务执行表';