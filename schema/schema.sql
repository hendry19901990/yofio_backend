-- MySQL dump 10.13  Distrib 8.0.22, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: test_db
-- ------------------------------------------------------
-- Server version	5.7.32

--
-- Table structure for table `credit_Type`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `credit_Type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `credit_type_300` int(11) NOT NULL,
  `credit_type_500` int(11) NOT NULL,
  `credit_type_700` int(11) NOT NULL,
  `investment` int(11) NOT NULL,
  `success` tinyint(1) NOT NULL,
  `date_created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;


