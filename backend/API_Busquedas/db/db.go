package db

import (
	actividadClient "api_busquedas/clients/actividades"
	inscripcionClient "api_busquedas/clients/inscripciones"
	usuarioClient "api_busquedas/clients/usuarios"
	_"os"

	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitConnection() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(mysql_usuarios:3306)/usuarios_db?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := fmt.Sprintf(
	// 	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	fmt.Println("DSN generado:", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("failed to connect database %v", err)
		panic("failed to connect database")
	}

	// API_Busquedas NO crea tablas, solo consulta
	// Las tablas están en:
	// - MongoDB (actividades) - API_Actividades
	// - MySQL (usuario, inscripcion) - API_Usuarios
	log.Info("API_Busquedas: DB connection established (read-only access)")
	
	// Solo inicializar clientes para consultas READ-ONLY
	actividadClient.Db = db
	inscripcionClient.Db = db
	usuarioClient.Db = db

	return db
}

func Close(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to close database")
	}
	sqlDB.Close()
}
