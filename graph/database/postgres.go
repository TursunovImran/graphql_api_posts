package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TursunovImran/graphql_api_posts/graph/dbmodel"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Переменная для хранения данных о подключении к базе данных
var DBInstance *gorm.DB

// Переменная для обработки ошибок
var err error


func ConnectDB() {
	var POSTGRES_CONNECTION_STRING string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
 													os.Getenv("PG_USERNAME"),
													os.Getenv("PG_PASSWORD"),
													os.Getenv("DB_HOST"),
													os.Getenv("PG_PORT"),
													os.Getenv("PG_BNAME"),
													os.Getenv("PG_SSLMODE"))

	DBInstance, err = gorm.Open("postgres", POSTGRES_CONNECTION_STRING)
	if err != nil {
		log.Fatal(err)
   	} else {
		log.Println("Database connected successfully.....")
    }

    DBInstance.LogMode(true)
}


func MigrateDB() {
	// Создание таблиц, ВАЖНО: Соблюдать последовательность
	DBInstance.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{})

	// Отношение "Один ко многим" между Post и Comment
	DBInstance.Model(&dbmodel.Post{}).Association("Comments").Append(&dbmodel.Comment{})
	// Отношение "Один ко многим" между Author и Post
	DBInstance.Model(&dbmodel.User{}).Association("Posts").Append(&dbmodel.Post{})
	// Отношение "Один ко многим" между Author и Comment
	DBInstance.Model(&dbmodel.User{}).Association("Comments").Append(&dbmodel.Comment{})

	// Определение внешних ключей
	DBInstance.Model(&dbmodel.Comment{}).AddForeignKey("post_id", "posts(id)", "CASCADE", "CASCADE")
	DBInstance.Model(&dbmodel.Comment{}).AddForeignKey("parent_id", "comments(id)", "CASCADE", "CASCADE")
	DBInstance.Model(&dbmodel.Comment{}).AddForeignKey("author_id", "users(id)", "CASCADE", "CASCADE")

	DBInstance.Model(&dbmodel.Post{}).AddForeignKey("author_id", "users(id)", "CASCADE", "CASCADE")

	log.Println("Database migration completed.....")
}