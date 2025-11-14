package db

import (
	inscripcionesClient "api_usuarios/clients/inscripciones"
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

	// Migrar tablas de usuarios e inscripciones
	if err := db.AutoMigrate(&model.Usuario{}, &model.Inscripcion{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Info("Database migration completed (Usuario, Inscripcion)")

	// Inicializar los clientes
	usuarioClient.Db = db
	inscripcionesClient.Db = db

	// Seed inicial de usuarios si la base está vacía
	seedInitialData(db)

	return db
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func seedInitialData(db *gorm.DB) {
	// Verificar si ya existen usuarios
	var count int64
	db.Model(&model.Usuario{}).Count(&count)

	if count > 0 {
		log.Info("Database already has users, skipping seed data")
		return
	}

	log.Info("Seeding initial user data...")

	// Hash SHA256 de las passwords
	// admin -> 8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918
	// password123 -> ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f

	usuarios := []model.Usuario{
		{
			Username:     "admin",
			Email:        "admin@gym.com",
			Nombre:       "Admin",
			Apellido:     "Sistema",
			PasswordHash: "8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918",
			EsAdmin:      true,
		},
		{
			Username:     "juan_gomez",
			Email:        "juan.gomez@email.com",
			Nombre:       "Juan",
			Apellido:     "Gomez",
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			EsAdmin:      false,
		},
		{
			Username:     "maria_lopez",
			Email:        "maria.lopez@email.com",
			Nombre:       "Maria",
			Apellido:     "Lopez",
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			EsAdmin:      false,
		},
		{
			Username:     "carlos_diaz",
			Email:        "carlos.diaz@email.com",
			Nombre:       "Carlos",
			Apellido:     "Diaz",
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			EsAdmin:      false,
		},
	}

	for _, usuario := range usuarios {
		if err := db.Create(&usuario).Error; err != nil {
			log.Errorf("Failed to seed user %s: %v", usuario.Username, err)
		} else {
			log.Infof("✅ Seeded user: %s", usuario.Username)
		}
	}

	log.Info("Initial user data seeded successfully")
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}
