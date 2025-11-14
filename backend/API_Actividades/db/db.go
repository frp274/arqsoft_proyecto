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

	// 3) Seed inicial con todas las actividades (idempotente)
	if err := seedActividades(ctx, db, col); err != nil {
		return err
	}
	
	return nil
}

func seedActividades(ctx context.Context, db *mongo.Database, col string) error {
	// Verificar si ya existen actividades
	count, err := db.Collection(col).CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	
	if count > 0 {
		log.Info("Database already has activities, skipping seed data")
		return nil
	}
	
	log.Info("Seeding initial activity data...")
	
	actividades := []interface{}{
		bson.M{
			"nombre":      "Yoga Integral",
			"descripcion": "Clases de yoga para todos los niveles. Mejora tu flexibilidad, fuerza y equilibrio mental.",
			"profesor":    "Laura Martinez",
			"tags":        bson.A{"yoga", "relajacion", "flexibilidad"},
			"horarios": bson.A{
				bson.M{"dia": "Lunes", "horaInicio": "08:00", "horaFin": "09:00", "cupo": 15},
				bson.M{"dia": "Miércoles", "horaInicio": "08:00", "horaFin": "09:00", "cupo": 15},
				bson.M{"dia": "Viernes", "horaInicio": "18:00", "horaFin": "19:00", "cupo": 15},
			},
		},
		bson.M{
			"nombre":      "CrossFit Avanzado",
			"descripcion": "Entrenamiento de alta intensidad combinando levantamiento de pesas, gimnasia y cardio.",
			"profesor":    "Roberto Sanchez",
			"tags":        bson.A{"crossfit", "fuerza", "intensidad"},
			"horarios": bson.A{
				bson.M{"dia": "Lunes", "horaInicio": "19:00", "horaFin": "20:00", "cupo": 20},
				bson.M{"dia": "Miércoles", "horaInicio": "19:00", "horaFin": "20:00", "cupo": 20},
				bson.M{"dia": "Viernes", "horaInicio": "19:00", "horaFin": "20:00", "cupo": 20},
			},
		},
		bson.M{
			"nombre":      "Spinning Indoor",
			"descripcion": "Entrenamiento cardiovascular sobre bicicleta estática con música motivadora.",
			"profesor":    "Ana Rodriguez",
			"tags":        bson.A{"spinning", "cardio", "resistencia"},
			"horarios": bson.A{
				bson.M{"dia": "Martes", "horaInicio": "07:00", "horaFin": "08:00", "cupo": 25},
				bson.M{"dia": "Jueves", "horaInicio": "07:00", "horaFin": "08:00", "cupo": 25},
				bson.M{"dia": "Sábado", "horaInicio": "09:00", "horaFin": "10:00", "cupo": 25},
			},
		},
		bson.M{
			"nombre":      "Pilates Mat",
			"descripcion": "Método de ejercicio que enfatiza el equilibrio, la postura y la respiración para fortalecer el core.",
			"profesor":    "Sofia Fernandez",
			"tags":        bson.A{"pilates", "core", "postura"},
			"horarios": bson.A{
				bson.M{"dia": "Martes", "horaInicio": "10:00", "horaFin": "11:00", "cupo": 12},
				bson.M{"dia": "Jueves", "horaInicio": "10:00", "horaFin": "11:00", "cupo": 12},
			},
		},
		bson.M{
			"nombre":      "Entrenamiento Funcional",
			"descripcion": "Ejercicios que imitan movimientos cotidianos para mejorar la funcionalidad y prevenir lesiones.",
			"profesor":    "Diego Torres",
			"tags":        bson.A{"funcional", "movimiento", "prevencion"},
			"horarios": bson.A{
				bson.M{"dia": "Lunes", "horaInicio": "18:00", "horaFin": "19:00", "cupo": 18},
				bson.M{"dia": "Miércoles", "horaInicio": "18:00", "horaFin": "19:00", "cupo": 18},
				bson.M{"dia": "Viernes", "horaInicio": "07:00", "horaFin": "08:00", "cupo": 18},
			},
		},
		bson.M{
			"nombre":      "Zumba Fitness",
			"descripcion": "Baile fitness con ritmos latinos que combina cardio y tonificación muscular.",
			"profesor":    "Valentina Castro",
			"tags":        bson.A{"zumba", "baile", "diversion"},
			"horarios": bson.A{
				bson.M{"dia": "Martes", "horaInicio": "19:00", "horaFin": "20:00", "cupo": 30},
				bson.M{"dia": "Jueves", "horaInicio": "19:00", "horaFin": "20:00", "cupo": 30},
				bson.M{"dia": "Sábado", "horaInicio": "10:00", "horaFin": "11:00", "cupo": 30},
			},
		},
		bson.M{
			"nombre":      "Boxeo Fitness",
			"descripcion": "Entrenamiento de boxeo para mejorar resistencia cardiovascular, coordinación y fuerza.",
			"profesor":    "Martin Ruiz",
			"tags":        bson.A{"boxeo", "fuerza", "coordinacion"},
			"horarios": bson.A{
				bson.M{"dia": "Lunes", "horaInicio": "20:00", "horaFin": "21:00", "cupo": 16},
				bson.M{"dia": "Miércoles", "horaInicio": "20:00", "horaFin": "21:00", "cupo": 16},
				bson.M{"dia": "Viernes", "horaInicio": "20:00", "horaFin": "21:00", "cupo": 16},
			},
		},
		bson.M{
			"nombre":      "Natación Intermedia",
			"descripcion": "Clases de natación para mejorar técnica y resistencia en los diferentes estilos.",
			"profesor":    "Paula Medina",
			"tags":        bson.A{"natacion", "tecnica", "resistencia"},
			"horarios": bson.A{
				bson.M{"dia": "Martes", "horaInicio": "06:00", "horaFin": "07:00", "cupo": 10},
				bson.M{"dia": "Jueves", "horaInicio": "06:00", "horaFin": "07:00", "cupo": 10},
				bson.M{"dia": "Sábado", "horaInicio": "08:00", "horaFin": "09:00", "cupo": 10},
			},
		},
		bson.M{
			"nombre":      "Stretching y Flexibilidad",
			"descripcion": "Clases enfocadas en mejorar la flexibilidad, reducir tensión muscular y prevenir lesiones.",
			"profesor":    "Lucia Vargas",
			"tags":        bson.A{"stretching", "flexibilidad", "recuperacion"},
			"horarios": bson.A{
				bson.M{"dia": "Lunes", "horaInicio": "12:00", "horaFin": "13:00", "cupo": 20},
				bson.M{"dia": "Miércoles", "horaInicio": "12:00", "horaFin": "13:00", "cupo": 20},
				bson.M{"dia": "Viernes", "horaInicio": "12:00", "horaFin": "13:00", "cupo": 20},
			},
		},
		bson.M{
			"nombre":      "GAP (Glúteos-Abdomen-Piernas)",
			"descripcion": "Entrenamiento localizado para tonificar y fortalecer el tren inferior y abdomen.",
			"profesor":    "Carolina Paz",
			"tags":        bson.A{"gap", "tonificacion", "fuerza"},
			"horarios": bson.A{
				bson.M{"dia": "Martes", "horaInicio": "18:00", "horaFin": "19:00", "cupo": 25},
				bson.M{"dia": "Jueves", "horaInicio": "18:00", "horaFin": "19:00", "cupo": 25},
			},
		},
	}
	
	result, err := db.Collection(col).InsertMany(ctx, actividades)
	if err != nil {
		return err
	}
	
	log.Infof("✅ Seeded %d activities successfully", len(result.InsertedIDs))
	return nil
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
