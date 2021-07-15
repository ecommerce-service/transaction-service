package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"time"
)

type IConnection interface {
	Connect() (IConnection, error)
	Pool()
	Migration(migrationDirectory string)
	GetDbInstance() *sql.DB
	Begin() (*sql.Tx, error)
}

type Connection struct {
	Config *Config
	db     *sql.DB
}

func NewConnection(config *Config) IConnection {
	return &Connection{Config: config}
}

func (c *Connection) Connect() (IConnection, error) {
	var err error
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&TimeZone=%s", c.Config.User(), c.Config.Password(), c.Config.Host(), c.Config.Port(), c.Config.DbName(),
		c.Config.SslMode(), c.Config.Location(),
	)

	c.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return c, err
}

func (c *Connection) Pool() {
	c.db.SetMaxOpenConns(c.Config.DBMaxConnection())
	c.db.SetMaxIdleConns(c.Config.DBMAxIdleConnection())
	c.db.SetConnMaxLifetime(time.Duration(c.Config.DBMaxLifeTimeConnection()) * time.Second)
}

func (c *Connection) Migration(migrationDirectory string) {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDirectory,
	}
	n, err := migrate.Exec(c.db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Error migration := ", err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
}

func (c *Connection) GetDbInstance() *sql.DB {
	return c.db
}

func (c *Connection) Begin() (*sql.Tx, error) {
	res, err := c.db.Begin()

	return res, err
}
