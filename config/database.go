package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gocomprobem/models"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Gagal membaca file .env")
    }

    dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
        "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
        os.Getenv("DB_NAME") + "?parseTime=true"

    // konek
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal koneksi ke database:", err)
    }

    // migrasi
    database.AutoMigrate(&models.User{}, &models.Report{})

    DB = database
}
