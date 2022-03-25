package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/patricksferraz/guest-check/domain/entity"
	_ "gorm.io/driver/sqlite"
)

type PostgreSQL struct {
	Db *gorm.DB
}

func NewPostgreSQL(dsnType, dsn string) (*PostgreSQL, error) {
	pg := &PostgreSQL{}

	err := pg.connect(dsnType, dsn)
	if err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *PostgreSQL) connect(dsnType, dsn string) error {
	db, err := gorm.Open(dsnType, dsn)
	if err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}

	p.Db = db

	return nil
}

func (p *PostgreSQL) Debug(enable bool) {
	p.Db.LogMode(enable)
}

func (p *PostgreSQL) Migrate() {
	p.Db.AutoMigrate(
		&entity.Guest{},
		&entity.Place{},
		&entity.Attendant{},
		&entity.GuestCheck{},
		&entity.GuestCheckItem{},
	)
}
