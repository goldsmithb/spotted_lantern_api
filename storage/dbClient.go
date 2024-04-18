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

func (db *dbClient) GetAllKills() ([]int, error) {
	scores := make([]int, 0)
	rows, err := db.cxn.Query(`SELECT score FROM users`)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var score int
		err = rows.Scan(&score)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (db *dbClient) GetKillCount(userId string) (int, error) {
	//query := `-- SELECT * FROM users WHERE id = $1`
	var score int
	row := db.cxn.QueryRow(`SELECT score FROM users WHERE id=$1`, userId)
	err := row.Scan(&score)
	if err != nil {
		return -1, err
	}
	return score, nil
}
