package db

import (
	"Chat/pkg/models"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
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
	log.Println("Trying to insert user " + user.Login)
	err := d.db.QueryRow(context.Background(), `insert into public.users(login, password) values($1, $2) returning id`, user.Login, user.Password).Scan(&user.ID)
	if err != nil {
		return 0, err
	}
	log.Printf("User %v %v\n added successfully", user.ID, user.Login)
	return user.ID, nil
}

func (d *DB) GetUserByID(id int) (models.User, error) {
	user := models.User{ID: id}
	var login pgtype.Text
	var password pgtype.Text
	err := d.db.QueryRow(context.Background(), `select login, password from public.users where id=$1`, id).Scan(&login, &password)
	user.Login = login.String
	user.Password = password.String
	if err != nil {
		return models.User{}, err
	}
	log.Printf("Returning user %v %v\n", user.ID, user.Password)
	return user, nil
}

func (d *DB) GetUserByLogin(login string) (models.User, error) {
	user := models.User{Login: login}
	var user_id pgtype.Int4
	var password pgtype.Text
	err := d.db.QueryRow(context.Background(), `select id, password from public.users where login=$1`, login).Scan(&user_id, &password)
	user.Login = login
	user.ID = int(user_id.Int32)
	user.Password = password.String
	if err != nil {
		return models.User{}, err
	}
	log.Printf("Returning user %v %v %v\n", user.ID, user.Login, user.Password[:10])
	return user, nil
}

func (d *DB) CheckSameLogin(login string) (bool, error) {
	var id int
	err := d.db.QueryRow(context.Background(), `select id from public.users where login=$1`, login).Scan(&id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *DB) GetRefreshToken(id int) (string, error) {
	var pgtoken pgtype.Text
	err := d.db.QueryRow(context.Background(), `select refresh_token from public.users where id=$1`, id).Scan(&pgtoken)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return "", nil
		}
		return "", err
	}
	return pgtoken.String, nil
}

func (d *DB) InsertRefreshToken(token string, id int) error {
	_, err := d.db.Exec(context.Background(), `update public.users set refresh_token=$1 where id=$2`, token, id)
	return err
}
