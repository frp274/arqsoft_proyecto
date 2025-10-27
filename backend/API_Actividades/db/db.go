package initdb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	URI        string
	Database   string
	TimeoutSec int
}

func Connect(cfg MongoConfig) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.TimeoutSec)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}
	// ⚠️ AÑADIR ESTE LOG ⚠️
	log.Printf("INFO: Conexión a MongoDB exitosa a DB: %s", cfg.Database)
	return client, client.Database(cfg.Database), nil
}

// RunMigrations corre todo lo necesario para que el servicio arranque con la BD lista.
func RunMigrations(ctx context.Context, db *mongo.Database) error {
	if err := ensureActividadesCollection(ctx, db); err != nil {
		return err
	}
	return nil
}

// =============================================================================================
// ACTIVIDADES
// =============================================================================================
func ensureActividadesCollection(ctx context.Context, db *mongo.Database) error {
	const col = "actividades"

	// 1) Si no existe la colección, crearla con validator
	has, err := collectionExists(ctx, db, col)
	if err != nil {
		return err
	}
	if !has {
		validator := bson.M{
			"$jsonSchema": bson.M{
				"bsonType": "object",
				"required": bson.A{"nombre", "profesor", "horarios"},
				"properties": bson.M{
					"nombre":      bson.M{"bsonType": "string", "minLength": 3},
					"descripcion": bson.M{"bsonType": "string"},
					"profesor":    bson.M{"bsonType": "string", "minLength": 3},
					"horarios": bson.M{
						"bsonType": "array",
						"items": bson.M{
							"bsonType": "object",
							"required": bson.A{"dia", "horaInicio", "horaFin", "cupo"},
							"properties": bson.M{
								"cupo":       bson.M{"bsonType": "int", "minimum": 1, "maximum": 200},
								"dia":        bson.M{"bsonType": "string"},
								"horaInicio": bson.M{"bsonType": "string", "pattern": "^[0-2][0-9]:[0-5][0-9]$"},
								"horaFin":    bson.M{"bsonType": "string", "pattern": "^[0-2][0-9]:[0-5][0-9]$"},
							},
						},
					},
					"tags": bson.M{"bsonType": "array", "items": bson.M{"bsonType": "string"}},
				},
			},
		}
		opts := options.CreateCollection().SetValidator(validator)
		if err := db.CreateCollection(ctx, col, opts); err != nil {
			return err
		}
	}

	// 2) Índices (idempotentes)
	idxView := db.Collection(col).Indexes()
	_, err = idxView.CreateMany(ctx, []mongo.IndexModel{
		{Keys: bson.D{{Key: "nombre", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "horarios.dia", Value: 1}}},
	})
	if err != nil {
		return err
	}

	// 3) Seed opcional (idempotente)
	_, err = db.Collection(col).UpdateOne(
		ctx,
		bson.M{"nombre": "Funcional"},
		bson.M{"$setOnInsert": bson.M{
			"profesor":    "Carlos Perez",
			"descripcion": "Entrenamiento funcional",
			"horarios": bson.A{
				bson.M{"dia": "lunes", "horaInicio": "07:30", "horaFin": "08:30", "cupo": 20},
				bson.M{"dia": "martes", "horaInicio": "07:30", "horaFin": "08:30", "cupo": 20},
				bson.M{"dia": "jueves", "horaInicio": "07:30", "horaFin": "08:30", "cupo": 20},
			},
			//"createdAt":         time.Now(),
			//"updatedAt":         time.Now(),
		}},
		options.Update().SetUpsert(true),
	)
	return err
}

func collectionExists(ctx context.Context, db *mongo.Database, name string) (bool, error) {
	cur, err := db.ListCollections(ctx, bson.D{{Key: "name", Value: name}})
	if err != nil {
		return false, err
	}
	defer cur.Close(ctx)
	return cur.Next(ctx), cur.Err()
}

//package db
//
//import (
//	actividadClient "api_actividades/clients/actividades"
//	_"os"
//
//	model "api_actividades/model"
//	"fmt"
//
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/schema"
//)
//
//func InitConnection() *gorm.DB {
//	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
//	dsn := "root:genagena1@tcp(mysql:3306)/arquisoftware?charset=utf8mb4&parseTime=True&loc=Local"
//	// dsn := fmt.Sprintf(
//	// 	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
//	// 	os.Getenv("DB_USER"),
//	// 	os.Getenv("DB_PASSWORD"),
//	// 	os.Getenv("DB_HOST"),
//	// 	os.Getenv("DB_PORT"),
//	// 	os.Getenv("DB_NAME"),
//	// )
//
//	fmt.Println("DSN generado:", dsn)
//
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//	})
//	if err != nil {
//		fmt.Printf("failed to connect database %v", err)
//		panic("failed to connect database")
//	}
//
//	db.AutoMigrate(&model.Actividad{}, &model.Horario{})
//	actividadClient.Db = db
//
//	return db
//}
//
//func Close(db *gorm.DB) {
//	sqlDB, err := db.DB()
//	if err != nil {
//		panic("failed to close database")
//	}
//	sqlDB.Close()
//}
