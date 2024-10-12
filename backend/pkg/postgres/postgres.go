package postgres

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm/schema"
	"time"

	"gorm.io/gorm"
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

func New(ctx context.Context, dbURL string, opts ...Option) (*Postgres, error) {
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

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *Postgres) Close(ctx context.Context) {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return
	}

	if err := sqlDB.Close(); err != nil {
		return
	}
}
