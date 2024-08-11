package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/harrymuliawan03/go-rest-api/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open connection : ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("failed to ping connection : ", err.Error())
	}

	return db
}
