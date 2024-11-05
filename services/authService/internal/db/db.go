package db

import (
	"Chat/pkg/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	User     string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"password" env-default:"postgres"`
	DBName   string `yaml:"dbname" env:"DBNAME" env-default:"chat"`
}

type DB struct {
	config *Config
	db     *pgx.Conn
}

func New(cfg *Config) (*DB, error) {
	d := &DB{config: cfg}
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	db, err := pgx.Connect(context.Background(), connection)
	log.Println("Connecting to: " + connection)
	if err != nil {
		return nil, err
	}
	d.db = db
	return d, nil
}

func (d *DB) Close() error {
	return d.db.Close(context.Background())
}

func (d *DB) AddUser(user models.User) (int, error) {
	var id int
	err := d.db.QueryRow(context.Background(), `insert into users(login, password) values($1, $2)`, user.Login, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d *DB) GetUserByID(id int) (models.User, error) {
	user := models.User{ID: id}
	err := d.db.QueryRow(context.Background(), `select login, password from users where id=$1`, id).Scan(&user.Login, &user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (d *DB) GetUserByLogin(login string) (models.User, error) {
	user := models.User{Login: login}
	err := d.db.QueryRow(context.Background(), `select id, password from users where login=$1`, login).Scan(&user.ID, user.Password)
	if err != nil {
		return models.User{}, err
	}
	return user, err
}

func (d *DB) CheckSameLogin(login string) (bool, error) {
	var id int
	err := d.db.QueryRow(context.Background(), `select id from users where login=$1`, login).Scan(&id)
	if err != nil {
		return false, err
	}
	return id > 0, nil
}
