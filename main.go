package main

import (
	"fiber-poc-api/database/entity"
	internal "fiber-poc-api/routes"
	"fmt"
	"strings"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {

	// ==> Get config from config.yaml
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("fatal error config file: %+v \n", err)
	}

	// ==> Connect database (PostgreSQL)
	db := databaseConnection()

	// ควรเช็ค nil ก่อนใช้งาน db ต่อ
	if db != nil {
		generateTable := true
		log.Infof("config generateTable: %t", generateTable)
		if generateTable {
			// AutoMigrate คืนค่า error ด้วย ควร handle หรือ log ไว้
			if err := db.AutoMigrate(
				&entity.User{},
				&entity.Role{},
				&entity.Privilege{},
				&entity.UserRole{},
				&entity.RolePrivilege{},
				&entity.LoginHistory{},
			); err != nil {
				log.Errorf("Migration failed: %v", err)
			} else {
				log.Infof("Generate Tables success")
			}
		}
	}

	app := fiber.New()

	// ==> cors (รวมเหลืออันเดียว และเพิ่ม Authorization)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// ==> JWT Middleware configuration (Updated for contrib/jwt)
	jwtMiddleware := jwtware.New(jwtware.Config{
		// Syntax ใหม่ของ contrib/jwt ต้องระบุ Key ใน Struct
		SigningKey: jwtware.SigningKey{
			Key: []byte(viper.GetString("jwt.secret")),
		},
		ErrorHandler: jwtError,
	})

	// ==> routes
	// หมายเหตุ: ตรวจสอบว่าใน internal.Router รับ parameter เป็น fiber.Handler หรือไม่
	internal.Router(app, jwtMiddleware, db)

	// ==> server start
	port := viper.GetString("server.port")
	log.Infof("Server is starting on port: %s", port)
	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Errorf("error server: %+v \n", err.Error())
		return
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	// เพิ่มการ Log error เพื่อให้ Debug ง่ายขึ้นว่าทำไม Unauthorized (เช่น Token Expired หรือ Signature ผิด)
	if err != nil {
		log.Warnf("JWT Error: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized: " + err.Error(),
		})
	}
	return c.Next()
}

func databaseConnection() *gorm.DB {

	newLogger := logger.New(
		// ใช้ Standard Log Writer แทน nil (เพื่อให้ log ออกมาที่ console ได้ถ้าต้องการ)
		nil,
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.password"),
		viper.GetString("database.name"),
		viper.GetString("database.port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",
			NoLowerCase: true, // ระวัง: PostgreSQL ชอบ table name ตัวเล็ก (lowercase) การใช้ NoLowerCase อาจมีปัญหากับ Convention ทั่วไปของ PG
		},
	})
	if err != nil {
		log.Errorf("error database: %+v \n", err.Error())
		return nil
	}

	fmt.Println("Successfully connected to the database")
	return db
}
