package postgres

import (
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log/slog"
	"time"
)

const (
	_defaultMaxPoolSize  = 10
	_defaultConnAttempts = 10
	_defaultConnTimeout  = time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	DB *gorm.DB
}

func New(dbURL string, log *slog.Logger, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  _defaultMaxPoolSize,
		connAttempts: _defaultConnAttempts,
		connTimeout:  _defaultConnTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(pg)
	}

	var db *gorm.DB
	var err error

	fmt.Println("Connecting to database with URL:", dbURL)

	for pg.connAttempts > 0 {
		db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err == nil {
			sqlDB, err := db.DB()
			if err != nil {
				return nil, fmt.Errorf("postgres - New - db.DB: %w", err)
			}
			sqlDB.SetMaxOpenConns(pg.maxPoolSize)
			sqlDB.SetConnMaxLifetime(time.Hour)

			pg.DB = db
			break
		}

		log.Debug(
			"postgres is trying to connect, attempts left: %d", pg.connAttempts,
		)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		log.Error("failed to connect to database: %s", err.Error())
		return nil, err
	}

	return pg, nil
}

func (p *Postgres) Close() {
	sqlDB, err := p.DB.DB()
	if err != nil {
		log.Info("Error getting underlying database connection: %s", err.Error())
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Info("Error closing database connection: %s", err.Error())
	}
}
