CREATE DATABASE  IF NOT EXISTS `arquisoftware` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `arquisoftware`;
-- MySQL dump 10.13  Distrib 8.0.42, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: arquisoftware
-- ------------------------------------------------------
-- Server version	9.3.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actividad`
--

DROP TABLE IF EXISTS `actividad`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `actividad` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` longtext NOT NULL,
  `descripcion` varchar(250) NOT NULL,
  `profesor` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actividad`
--

LOCK TABLES `actividad` WRITE;
/*!40000 ALTER TABLE `actividad` DISABLE KEYS */;
INSERT INTO `actividad` VALUES (2,'Pilates','Actividad enfocada en fortalecer el core, mejorar la postura y la flexibilidad.','Laura Fernández'),(3,'Spinning','Clase aeróbica sobre bicicleta fija ideal para quemar calorías y mejorar la capacidad cardiovascular.','Martín Pérez'),(4,'Musculacion','Entrenamiento con pesas y máquinas para aumentar la fuerza y la masa muscular.','Ana Torres');
/*!40000 ALTER TABLE `actividad` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `horario`
--

DROP TABLE IF EXISTS `horario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `horario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `actividad_id` bigint DEFAULT NULL,
  `dia` longtext NOT NULL,
  `hora_inicio` longtext,
  `hora_fin` longtext NOT NULL,
  `cupo` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_actividad_horarios` (`actividad_id`),
  CONSTRAINT `fk_actividad_horarios` FOREIGN KEY (`actividad_id`) REFERENCES `actividad` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `horario`
--

LOCK TABLES `horario` WRITE;
/*!40000 ALTER TABLE `horario` DISABLE KEYS */;
INSERT INTO `horario` VALUES (4,2,'Martes','09:00','10:00',20),(5,2,'Jueves','09:00','10:00',20),(6,2,'Sábado','10:00','11:00',20),(7,3,'Lunes','19:00','20:00',20),(8,3,'Miércoles','19:00','20:00',20),(9,4,'Lunes','08:00','09:00',20),(10,4,'Martes','08:00','09:00',20),(11,4,'Miércoles','08:00','09:00',20),(12,4,'Jueves','08:00','09:00',20),(13,4,'Viernes','08:00','09:00',20);
/*!40000 ALTER TABLE `horario` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `inscripcion`
--

DROP TABLE IF EXISTS `inscripcion`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `inscripcion` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `horario_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `actividad_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_inscripcion_horario` (`horario_id`),
  KEY `fk_inscripcion_usuario` (`usuario_id`),
  KEY `fk_inscripcion_actividad` (`actividad_id`),
  CONSTRAINT `fk_inscripcion_actividad` FOREIGN KEY (`actividad_id`) REFERENCES `actividad` (`id`),
  CONSTRAINT `fk_inscripcion_horario` FOREIGN KEY (`horario_id`) REFERENCES `horario` (`id`),
  CONSTRAINT `fk_inscripcion_usuario` FOREIGN KEY (`usuario_id`) REFERENCES `usuario` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `inscripcion`
--

LOCK TABLES `inscripcion` WRITE;
/*!40000 ALTER TABLE `inscripcion` DISABLE KEYS */;
/*!40000 ALTER TABLE `inscripcion` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuario`
--

DROP TABLE IF EXISTS `usuario`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre_apellido` longtext NOT NULL,
  `Username` longtext,
  `es_admin` tinyint(1) NOT NULL,
  `password_hash` longtext NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuario`
--

LOCK TABLES `usuario` WRITE;
/*!40000 ALTER TABLE `usuario` DISABLE KEYS */;
INSERT INTO `usuario` VALUES (1,'Genaro','genacanas',1,'66e49ff234ab9d28adf8662bc56bf31c0e5fec55f0ff2212731b094a1586e961');
/*!40000 ALTER TABLE `usuario` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-07-22 21:01:38
