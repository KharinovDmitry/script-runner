package migrator

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

func MustRun(driverName, connStr, migrationsDir string) {
	db, err := sqlx.Open(driverName, connStr)
	if err != nil {
		panic("Ошибка во время открытия соединения: " + err.Error())
	}
	defer db.Close()

	if err := goose.SetDialect(driverName); err != nil {
		panic("Ошибка во время  выбора диалекта: " + err.Error())
	}

	if err := goose.Up(db.DB, migrationsDir); err != nil {
		panic("Ошибка во время выполнение миграций: " + err.Error())
	}
}
