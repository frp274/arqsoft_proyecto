-- Script de inicialización para API_Usuarios
-- Base de datos: usuarios_db

USE `usuarios_db`;

-- Tabla de usuarios
DROP TABLE IF EXISTS `usuario`;
CREATE TABLE `usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre_apellido` varchar(255) NOT NULL,
  `Username` varchar(100) NOT NULL UNIQUE,
  `es_admin` tinyint(1) NOT NULL DEFAULT 0,
  `password_hash` varchar(64) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_username` (`Username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Datos iniciales de usuarios
-- Contraseña de 'genacanas' es: genaro123 (hash SHA256: 3bd517332b9d96f9fbc0de89b613dc07b3101292fd54fd7cb52da0d8846303e2)
-- Contraseña de 'facubuffaz' es: facu123 (hash SHA256: 293bb6d0e7e4c2ee8761e60be2169d09d42156f4167fa58f3e2a0e39e78773d4)
INSERT INTO `usuario` (`id`, `nombre_apellido`, `Username`, `es_admin`, `password_hash`) VALUES 
(1, 'Genaro Canas', 'genacanas', 1, '3bd517332b9d96f9fbc0de89b613dc07b3101292fd54fd7cb52da0d8846303e2'),
(2, 'Facundo Buffaz', 'facubuffaz', 0, '293bb6d0e7e4c2ee8761e60be2169d09d42156f4167fa58f3e2a0e39e78773d4');

-- Usuario de prueba adicional
-- Contraseña de 'testuser' es: test123 (hash SHA256: ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae)
INSERT INTO `usuario` (`nombre_apellido`, `Username`, `es_admin`, `password_hash`) VALUES 
('Usuario de Prueba', 'testuser', 0, 'ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae');
