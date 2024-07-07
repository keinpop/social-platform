package db

import (
	"fmt"
	"mai-platform/internal/clients/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	Port     uint   `yaml:"port"`
	DataBase string `yaml:"database"`
}

type DB struct {
	db  *gorm.DB
	cfg *Config
}

func NewDB(cfg *Config) *DB {
	return &DB{cfg: cfg}
}

func (d *DB) Init() error {
	// dbURL := "postgres://postgres:db-frknz@db:5432/example" // example
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		d.cfg.Login,
		d.cfg.Password,
		d.cfg.Address,
		d.cfg.Port,
		d.cfg.DataBase,
	)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{TranslateError: true})
	if err != nil {
		return err
	}

	d.db = db

	err = db.SetupJoinTable(&models.User{}, "Companies", &models.UserCompanies{})
	if err != nil {
		return err
	}

	return nil
}

func (d *DB) Migrate() error {
	err := d.db.AutoMigrate(
		&models.Company{},
		&models.Programm{},
		&models.Role{},
		&models.Techonology{},
		&models.User{},
		&models.UserCompanies{},
		&models.Student{},
		&models.Teacher{},
		&models.Admin{},
	)

	return err
}

func (d *DB) AddCompany(title string) (*models.Company, error) {
	c := models.Company{Title: title}

	if result := d.db.Create(&c); result.Error != nil {
		return nil, result.Error
	}

	return &c, nil
}

func (d *DB) GetCompanies() ([]models.Company, error) {
	var companies []models.Company

	if result := d.db.Find(&companies); result.Error != nil {
		return nil, result.Error
	}

	return companies, nil
}

func (d *DB) AddProgramm(title string, duration uint64) (*models.Programm, error) {
	p := models.Programm{Title: title, Duration: duration}

	if result := d.db.Create(&p); result.Error != nil {
		return nil, result.Error
	}

	return &p, nil
}

func (d *DB) GetProgrammes() ([]models.Programm, error) {
	var programmes []models.Programm

	if result := d.db.Find(&programmes); result.Error != nil {
		return nil, result.Error
	}

	return programmes, nil
}

func (d *DB) AddRole(title string) (*models.Role, error) {
	r := models.Role{Title: title}

	if result := d.db.Create(&r); result.Error != nil {
		return nil, result.Error
	}

	return &r, nil
}

func (d *DB) GetRoles() ([]models.Role, error) {
	var roles []models.Role

	if result := d.db.Find(&roles); result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (d *DB) AddTechonology(title string) (*models.Techonology, error) {
	t := models.Techonology{Title: title}

	if result := d.db.Create(&t); result.Error != nil {
		return nil, result.Error
	}

	return &t, nil
}

func (d *DB) GetTechonologies() ([]models.Techonology, error) {
	var techonologies []models.Techonology

	if result := d.db.Find(&techonologies); result.Error != nil {
		return nil, result.Error
	}

	return techonologies, nil
}
