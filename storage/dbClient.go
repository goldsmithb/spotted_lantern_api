package storage

import (
	"database/sql"
	"fmt"
	"github.com/goldsmithb/spotted_lantern_api/config"
	"github.com/goldsmithb/spotted_lantern_api/core"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// implements core.dbClient
type dbClient struct {
	config *config.Config
	logger *zap.Logger
	cxn    *sql.DB
}

func NewDbClient(conf *config.Config, l *zap.Logger) *dbClient {
	return &dbClient{
		config: conf,
		logger: l,
	}
}

func (db *dbClient) Disconnect() error {
	return db.cxn.Close()
}

func (db *dbClient) Connect() error {
	opts := db.config.Options.Database

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s connect_timeout=%d",
		opts.Host, opts.Port, opts.UserName, opts.Password, opts.DefaultDb, opts.SSLMode, opts.ConnectTimeout)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		db.logger.Fatal("Failed to connect to database", zap.Error(err))
		return err
	}
	db.cxn = database
	return nil
}

func (db *dbClient) GetUserByEmail(email string) (*core.User, error) {
	var user core.User
	row := db.cxn.QueryRow(`SELECT * FROM lanternfly.users WHERE email = $1`, email)
	err := row.Scan(&user.UserId, &user.Username, &user.Email, &user.Hash, &user.Score)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (db *dbClient) GetHashForEmail(email string) (string, error) {
	var hash string
	const q = `SELECT hash FROM lanternfly.users WHERE email = $1`
	row := db.cxn.QueryRow(q, email)
	err := row.Scan(&hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func (db *dbClient) CreateUser(user core.User) error {
	const query = `INSERT INTO lanternfly.users (user_id, username, email, hash, score ) 
			VALUES ($1, $2, $3, $4, $5)`
	res, err := db.cxn.Exec(query, user.UserId, user.Username, user.Email, user.Hash, user.Score)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func (db *dbClient) GetAllUsers() ([]core.User, error) {
	const q = `SELECT * FROM lanternfly.users`
	users := make([]core.User, 0)
	rows, err := db.cxn.Query(q)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user core.User
		err = rows.Scan(&user.UserId, &user.Username, &user.Email, &user.Hash, &user.Score)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (db *dbClient) GetAllKills() ([]int, error) {
	scores := make([]int, 0)
	rows, err := db.cxn.Query(`SELECT score FROM lanternfly.users`)
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
	row := db.cxn.QueryRow(`SELECT score FROM lanternfly.users WHERE user_id=$1`, userId)
	err := row.Scan(&score)
	if err != nil {
		return -1, err
	}
	return score, nil
}
