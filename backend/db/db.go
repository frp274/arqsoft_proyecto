package db

import (
	actividadClient "arqsoft_proyecto/clients/actividades"
	inscripcionClient "arqsoft_proyecto/clients/inscripciones"
	usuarioClient "arqsoft_proyecto/clients/usuarios"

	model "arqsoft_proyecto/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitConnection() *gorm.DB {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details

	dsn := "root:facurp274@tcp(127.0.0.1:3306)/arquisoftware?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("failed to connect database %v", err)
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Actividad{}, &model.Usuario{}, &model.Inscripcion{})
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
