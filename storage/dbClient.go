package storage

import (
	"database/sql"
	"fmt"
	"github.com/goldsmithb/spotted_lantern_api/config"
	_ "github.com/lib/pq"
)

// implements core.dbClient
type dbClient struct {
	config *config.Config
	cxn    *sql.DB
}

func NewDbClient(conf *config.Config) *dbClient {
	return &dbClient{
		config: conf,
	}
}

func (db *dbClient) DisConnect() error {
	return db.cxn.Close()
}

func (db *dbClient) Connect() error {
	opts := db.config.Options.Database

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		opts.Host, opts.Port, opts.UserName, opts.Password, opts.DefaultDb, opts.SSLMode, opts.ConnectTimeout)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	err = database.Ping()
	if err == nil {
		fmt.Println("pinged db successfully :)")
	} else {
		fmt.Println("no ping db :(")
	}
	db.cxn = database
	return nil
}

func (db *dbClient) GetAllKills() int {
	return 100
}

func (db *dbClient) GetKillCount(userId string) int {
	//query := `-- SELECT * FROM users WHERE id = $1`

	//rows, _ := db.cxn.Query(query, userId)
	return 10
}
