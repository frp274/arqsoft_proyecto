package db

import (
	usuarioClient "api_usuarios/clients/usuarios"
	model "api_usuarios/model"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitConnection() *gorm.DB {
	// Usar variables de entorno para configuración
	dbHost := getEnv("DB_HOST", "mysql_usuarios")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "genagena1")
	dbName := getEnv("DB_NAME", "usuarios_db")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	log.Infof("Connecting to database: %s@%s:%s/%s", dbUser, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		panic("failed to connect database")
	}

	log.Info("Database connection established successfully")

	// Solo migrar la tabla de usuarios
	if err := db.AutoMigrate(&model.Usuario{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Info("Database migration completed")

	// Inicializar el cliente de usuarios
	usuarioClient.Db = db

	return db
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}
