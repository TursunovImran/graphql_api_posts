package mytest

import (
    "fmt"
    "os"
    "testing"

    "github.com/TursunovImran/graphql_api_posts/graph/dbmodel"
    "github.com/jinzhu/gorm"
    "github.com/stretchr/testify/assert"
    _ "github.com/lib/pq"
)

var dbHost, dbPort, dbName, dbUser, dbPass, dbSslMode string

func TestConnectDB(t *testing.T) {
    dbHost   = os.Getenv("DB_HOST")
    dbPort   = os.Getenv("PG_PORT")
    dbName   = os.Getenv("PG_BNAME")
    dbUser   = os.Getenv("PG_USERNAME")
    dbPass   = os.Getenv("PG_PASSWORD")
    dbSslMode = os.Getenv("PG_SSLMODE")
    // Подготовка тестовой строки подключения
    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode)

    // Подключение к базе данных
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Проверка подключения
    assert.NoError(t, err)
    assert.NotNil(t, db)
}

func TestMigrateDB(t *testing.T) {
    dbHost   = os.Getenv("DB_HOST")
    dbPort   = os.Getenv("PG_PORT")
    dbName   = os.Getenv("PG_BNAME")
    dbUser   = os.Getenv("PG_USERNAME")
    dbPass   = os.Getenv("PG_PASSWORD")
    dbSslMode = os.Getenv("PG_SSLMODE")
    // Подготовка тестовой строки подключения
    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode)

    // Подключение к базе данных
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Выполнение миграции
    err = db.AutoMigrate(&dbmodel.User{}, &dbmodel.Post{}, &dbmodel.Comment{}).Error
    if err != nil {
        t.Fatal(err)
    }

    // Проверка миграции
    var count int
    err = db.Model(&dbmodel.Comment{}).Count(&count).Error
    if err != nil {
        t.Fatal(err)
    }
    assert.Equal(t, 3, count)
}

func TestSetTestData(t *testing.T) {
    dbHost   = os.Getenv("DB_HOST")
    dbPort   = os.Getenv("PG_PORT")
    dbName   = os.Getenv("PG_BNAME")
    dbUser   = os.Getenv("PG_USERNAME")
    dbPass   = os.Getenv("PG_PASSWORD")
    dbSslMode = os.Getenv("PG_SSLMODE")
    // Подготовка тестовой строки подключения
    connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode)

    // Подключение к базе данных
    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        t.Fatal(err)
    }
    defer db.Close()

    // Добавление тестовых данных
    err = db.Exec("INSERT INTO users (username) VALUES ('testAuthorName');").Error
    if err != nil {
        t.Fatal(err)
    }

    // Проверка добавления тестовых данных
    var userCount int
    err = db.Model(&dbmodel.User{}).Count(&userCount).Error
    if err != nil {
        t.Fatal(err)
    }
    assert.Equal(t, 1, userCount)

    // Закрытие соединения с базой данных
    db.Close()
}
