
CREATE DATABASE IF NOT EXISTS blog_test;
USE blog_test;

--
-- Table structure for table `post_tags`
--

DROP TABLE IF EXISTS `post_tags`;

CREATE TABLE `post_tags` (
  `idPost` int(11) NOT NULL,
  `tag` varchar(128) NOT NULL,
  PRIMARY KEY (`idPost`,`tag`),
  UNIQUE INDEX (`idPost`, `tag`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `post_tags` WRITE;

INSERT INTO `post_tags` VALUES

(11, "Tag-A"),
(12, "Tag-A"),

(13, "Tag-A"),
(13, "Tag-B"),
(13, "Tag-C"),

(14, "Tag-A"),
(15, "Tag-D"),

(16, "Tag-E");


UNLOCK TABLES;

--
-- Table structure for table `posts`
--

DROP TABLE IF EXISTS `posts`;

CREATE TABLE `posts` (
  `idPost` int(11) NOT NULL AUTO_INCREMENT,
  `slug` varchar(255) NOT NULL,
  `title` varchar(255) DEFAULT NULL,
  `abstract` blob NOT NULL,
  `body` blob NOT NULL,
  `idUser` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `draft` tinyint(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`idPost`),
  KEY `slug` (`slug`),
  KEY `users` (`idUser`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `posts`
--

LOCK TABLES `posts` WRITE;

INSERT INTO `posts` VALUES 
(11,'slug-one','Title One','This post is about','Body of the post <strong>one</strong>',1,'2014-08-11 21:12:35',0),
(12,'slug-two','Title Two','This post is about','Body of the post <strong>two</strong>',1,'2014-08-11 21:13:00',0),
(13,'slug-three','Title Three','This post is about','Body of the post <strong>three</strong>',1,'2014-08-11 21:13:16',0),
(14,'draft-one','Draft One','This post is about','Body of the post <strong>draft-one</strong>',1,'2014-08-11 21:13:54',1),
(15,'draft-two','Draft Two','This post is about','Body of the post <strong>draft-two</strong>',1,'2014-08-11 21:14:08',1),
(17,'mypost-one','My Post One','This post is about','My post <em>one\'s</em> body',2,'2014-08-11 21:15:51',0),
(18,'mypost-two','My Post Two','This post is about','My post <em>two\'s</em> body',2,'2014-08-11 21:15:51',0),
(19,'mypost-three','My Post Three','This post is about','My post <em>three\'s</em> body',2,'2014-08-11 21:15:51',0),
(20,'mypost-four','My Post Four','This post is about','My post <em>four\'s</em> body',2,'2014-08-11 21:15:51',0),
(21,'mypost-draft','My Draft Post','This post is about','This is still a draft',2,'2014-08-11 21:15:52',1);

UNLOCK TABLES;


DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `idUser` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` char(32) NOT NULL,
  PRIMARY KEY (`idUser`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
-- password1, password

LOCK TABLES `users` WRITE;

INSERT INTO `users` VALUES 
(1,'Jeremy','jeremy@email.com','5f4dcc3b5aa765d61d8327deb882cf99'),
(2,'Mathias','mathias@company.it','7c6a180b36896a0a8c02787eeafb0e4c');

UNLOCK TABLES;
