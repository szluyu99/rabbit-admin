-- MySQL dump 10.13  Distrib 8.0.19, for Win64 (x86_64)
--
-- Host: localhost    Database: rabbit_admin
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article_tags`
--

DROP TABLE IF EXISTS `article_tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_tags` (
  `tag_id` bigint unsigned NOT NULL,
  `article_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`tag_id`,`article_id`),
  KEY `fk_article_tags_article` (`article_id`),
  CONSTRAINT `fk_article_tags_article` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`),
  CONSTRAINT `fk_article_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_tags`
--

LOCK TABLES `article_tags` WRITE;
/*!40000 ALTER TABLE `article_tags` DISABLE KEYS */;
INSERT INTO `article_tags` VALUES (1,1),(2,1),(2,2);
/*!40000 ALTER TABLE `article_tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `articles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `title` varchar(100) NOT NULL,
  `desc` varchar(200) DEFAULT NULL,
  `content` longtext,
  `cover` varchar(100) DEFAULT NULL,
  `type` varchar(20) DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  `is_top` tinyint(1) DEFAULT NULL,
  `is_delete` tinyint(1) DEFAULT NULL,
  `original_url` varchar(100) DEFAULT NULL,
  `category_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_articles_category` (`category_id`),
  CONSTRAINT `fk_articles_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'2023-06-05 23:31:00.587','2023-06-05 23:32:56.195','Congratulations on successfully running the project','','# Congratulations\n\nCongratulations on successfully running the project!\n\n```js\nconsole.log(\"Congratulations!\")\n```','http://cdn.hahacode.cn/167946151088e991a33534911ad2ab1bc5cb75f730.png','original','public',1,0,'',1),(2,'2023-06-05 23:32:50.193','2023-06-05 23:32:50.193','The methods of learning','','# 学习有捷径\n\n![cover](http://cdn.hahacode.cn/16794608978464e5d4342e4e11e0967b000afd679d.png)\n\n学习有捷径。学习的捷径之一就是多看看别人是怎么理解这些知识的。\n\n举两个例子。\n\n如果你喜欢《水浒》，千万不要只读原著当故事看，一定要读一读各代名家的批注和点评，看他们是如何理解的。之前学 C# 时，看《CLR via C#》晦涩难懂，但是我又通过看《你必须知道的.net》而更加了解了。因为后者就是中国一个 80 后写的，我通过他对 C# 的了解，作为桥梁和阶梯，再去窥探比较高达上的书籍和知识。\n\n最后，真诚的希望你能借助别人的力量来提高自己。我也一直在这样要求我自己。','http://cdn.hahacode.cn/16794608978464e5d4342e4e11e0967b000afd679d.png','reprint','public',0,0,'',2);
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'2023-06-05 23:28:08.221','2023-06-05 23:28:08.221','TECHNOLOGY'),(2,'2023-06-05 23:28:13.967','2023-06-05 23:28:13.967','LIFE');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `configs`
--

DROP TABLE IF EXISTS `configs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `configs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `key` longtext,
  `value` longtext,
  `desc` longtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `configs`
--

LOCK TABLES `configs` WRITE;
/*!40000 ALTER TABLE `configs` DISABLE KEYS */;
INSERT INTO `configs` VALUES (1,'USER_NEED_ACTIVATE','false',''),(2,'API_NEED_AUTH','true','if false, anyone can access any api'),(3,'CREATE_DEFAULT_ROLE','user','the default role when creating a user');
/*!40000 ALTER TABLE `configs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_members`
--

DROP TABLE IF EXISTS `group_members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group_members` (
  `group_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`group_id`,`user_id`),
  KEY `fk_group_members_user` (`user_id`),
  CONSTRAINT `fk_group_members_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`),
  CONSTRAINT `fk_group_members_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_members`
--

LOCK TABLES `group_members` WRITE;
/*!40000 ALTER TABLE `group_members` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `group_permissions`
--

DROP TABLE IF EXISTS `group_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `group_permissions` (
  `permission_id` bigint unsigned NOT NULL,
  `group_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`permission_id`,`group_id`),
  KEY `fk_group_permissions_group` (`group_id`),
  CONSTRAINT `fk_group_permissions_group` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`),
  CONSTRAINT `fk_group_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `group_permissions`
--

LOCK TABLES `group_permissions` WRITE;
/*!40000 ALTER TABLE `group_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `groups`
--

DROP TABLE IF EXISTS `groups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `groups` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_groups_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `groups`
--

LOCK TABLES `groups` WRITE;
/*!40000 ALTER TABLE `groups` DISABLE KEYS */;
/*!40000 ALTER TABLE `groups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `permissions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `parent_id` bigint unsigned DEFAULT NULL,
  `name` varchar(200) DEFAULT NULL,
  `p1` varchar(200) DEFAULT NULL,
  `p2` varchar(200) DEFAULT NULL,
  `p3` varchar(200) DEFAULT NULL,
  `anonymous` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_permissions_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,'2023-06-05 23:13:56.818','2023-06-05 23:13:56.818',0,'role','','','',0),(2,'2023-06-05 23:13:56.825','2023-06-05 23:13:56.825',1,'role.query','/role','POST','',1),(3,'2023-06-05 23:13:56.831','2023-06-05 23:13:56.831',1,'role.create','/role','PUT','',0),(4,'2023-06-05 23:13:56.836','2023-06-05 23:13:56.836',1,'role.delete','/role/:key','DELETE','',0),(5,'2023-06-05 23:13:56.841','2023-06-05 23:13:56.841',1,'role.update','/role/:key','PATCH','',0),(6,'2023-06-05 23:13:56.845','2023-06-05 23:13:56.845',0,'permission','','','',0),(7,'2023-06-05 23:13:56.851','2023-06-05 23:13:56.851',6,'permission.query','/permission','POST','',1),(8,'2023-06-05 23:13:56.856','2023-06-05 23:13:56.856',6,'permission.create','/permission','PUT','',0),(9,'2023-06-05 23:13:56.862','2023-06-05 23:13:56.862',6,'permission.delete','/permission/:key','DELETE','',0),(10,'2023-06-05 23:13:56.867','2023-06-05 23:13:56.867',6,'permission.update','/permission/:key','PATCH','',0),(11,'2023-06-05 23:13:56.872','2023-06-05 23:13:56.872',0,'user','','','',0),(12,'2023-06-05 23:13:56.877','2023-06-05 23:13:56.877',11,'user.query','/user','POST','',1),(13,'2023-06-05 23:13:56.883','2023-06-05 23:13:56.883',11,'user.role','/user/role','PATCH','',0),(14,'2023-06-05 23:13:56.888','2023-06-05 23:13:56.888',0,'config','','','',0),(15,'2023-06-05 23:13:56.892','2023-06-05 23:13:56.892',14,'config.query','/config','POST','',1),(16,'2023-06-05 23:13:56.897','2023-06-05 23:13:56.897',14,'config.create','/config','PUT','',0),(17,'2023-06-05 23:13:56.902','2023-06-05 23:13:56.902',14,'config.delete','/config/:key','DELETE','',0),(18,'2023-06-05 23:13:56.907','2023-06-05 23:13:56.907',14,'config.update','/config/:key','PATCH','',0),(19,'2023-06-05 23:13:56.911','2023-06-05 23:13:56.911',0,'article','','','',0),(20,'2023-06-05 23:13:56.914','2023-06-05 23:13:56.914',19,'article.get','/article/:key','GET','',1),(21,'2023-06-05 23:13:56.920','2023-06-05 23:13:56.920',19,'article.query','/article','POST','',1),(22,'2023-06-05 23:13:56.924','2023-06-05 23:13:56.924',19,'article.create','/article','PUT','',0),(23,'2023-06-05 23:13:56.930','2023-06-05 23:13:56.930',19,'article.save_or_update','/article/save_or_update','POST','',0),(24,'2023-06-05 23:13:56.935','2023-06-05 23:13:56.935',0,'tag','','','',0),(25,'2023-06-05 23:13:56.941','2023-06-05 23:13:56.941',24,'tag.get','/tag/:key','GET','',1),(26,'2023-06-05 23:13:56.946','2023-06-05 23:13:56.946',24,'tag.query','/tag','POST','',1),(27,'2023-06-05 23:13:56.952','2023-06-05 23:13:56.952',24,'tag.create','/tag','PUT','',0),(28,'2023-06-05 23:13:56.956','2023-06-05 23:13:56.956',24,'tag.update','/tag','PATCH','',0),(29,'2023-06-05 23:13:56.960','2023-06-05 23:13:56.960',24,'tag.all','/tag/all','GET','',1),(30,'2023-06-05 23:13:56.965','2023-06-05 23:13:56.965',0,'category','','','',0),(31,'2023-06-05 23:13:56.970','2023-06-05 23:13:56.970',30,'category.get','/category/:key','GET','',1),(32,'2023-06-05 23:13:56.976','2023-06-05 23:13:56.976',30,'category.query','/category','POST','',1),(33,'2023-06-05 23:13:56.981','2023-06-05 23:13:56.981',30,'category.create','/category','PUT','',0),(34,'2023-06-05 23:13:56.987','2023-06-05 23:13:56.987',30,'category.update','/category','PATCH','',0),(35,'2023-06-05 23:13:56.991','2023-06-05 23:13:56.991',30,'category.all','/category/all','GET','',1);
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_permissions`
--

DROP TABLE IF EXISTS `role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `role_permissions` (
  `role_id` bigint unsigned NOT NULL,
  `permission_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`),
  KEY `fk_role_permissions_permission` (`permission_id`),
  CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_permissions`
--

LOCK TABLES `role_permissions` WRITE;
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;
/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` varchar(50) DEFAULT NULL,
  `label` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_roles_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'2023-06-05 23:13:54.806','2023-06-05 23:13:54.806','admin','Admin'),(2,'2023-06-05 23:13:54.811','2023-06-05 23:13:54.811','user','User'),(3,'2023-06-05 23:13:54.817','2023-06-05 23:13:54.817','test','Test');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tags`
--

DROP TABLE IF EXISTS `tags`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tags` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tags_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tags`
--

LOCK TABLES `tags` WRITE;
/*!40000 ALTER TABLE `tags` DISABLE KEYS */;
INSERT INTO `tags` VALUES (1,'2023-06-05 23:31:00.576','2023-06-05 23:31:00.576','PROJECT'),(2,'2023-06-05 23:31:00.582','2023-06-05 23:31:00.582','STUDY');
/*!40000 ALTER TABLE `tags` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_roles`
--

DROP TABLE IF EXISTS `user_roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_roles` (
  `role_id` bigint unsigned NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`user_id`),
  KEY `fk_user_roles_user` (`user_id`),
  CONSTRAINT `fk_user_roles_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `fk_user_roles_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_roles`
--

LOCK TABLES `user_roles` WRITE;
/*!40000 ALTER TABLE `user_roles` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `group_id` bigint unsigned DEFAULT NULL,
  `email` varchar(128) DEFAULT NULL,
  `password` varchar(128) DEFAULT NULL,
  `first_name` varchar(128) DEFAULT NULL,
  `last_name` varchar(128) DEFAULT NULL,
  `display_name` varchar(128) DEFAULT NULL,
  `is_super_user` tinyint(1) DEFAULT NULL,
  `enabled` tinyint(1) DEFAULT NULL,
  `activated` tinyint(1) DEFAULT NULL,
  `last_login` datetime(3) DEFAULT NULL,
  `last_login_ip` varchar(128) DEFAULT NULL,
  `source` varchar(64) DEFAULT NULL,
  `locale` varchar(20) DEFAULT NULL,
  `timezone` varchar(200) DEFAULT NULL,
  `profile` longblob,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_users_email` (`email`),
  KEY `idx_users_group_id` (`group_id`),
  KEY `idx_users_source` (`source`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'2023-06-05 23:13:39.101','2023-06-05 23:39:50.461',0,'admin@test.com','sha256$8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92','','','',1,1,1,'2023-06-05 23:39:50.000','127.0.0.1','','','',NULL);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'rabbit_admin'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-06-05 23:40:53
