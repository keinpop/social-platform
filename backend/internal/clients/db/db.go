package db

import (
	"fmt"
	"log"
	"mai-platform/internal/clients/db/models"

	"golang.org/x/crypto/bcrypt"

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

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{TranslateError: true, AllowGlobalUpdate: true})
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
		&models.UserHash{},
	)

	return err
}

func (d *DB) Register(login, passwordHash string) error {
	u := models.UserHash{
		Login:        login,
		PasswordHash: passwordHash,
	}

	result := d.db.Create(&u)

	return result.Error
}

func (d *DB) CheckHash(login, password string) (bool, error) {
	var uh models.UserHash
	result := d.db.First(&uh, "login = ?", login)
	if result.Error != nil {
		return false, result.Error
	}

	tmpHash := []byte(uh.PasswordHash)
	tmpPass := []byte(login + password)

	if err := bcrypt.CompareHashAndPassword(tmpHash, tmpPass); err != nil {
		log.Println("[error] : Invalid password")
		return false, err
	}

	return true, nil
}

func (d *DB) AddUser(mail string, isStudent bool) (*models.User, error) {
	u := models.User{Mail: mail}
	if isStudent {
		s := models.Student{}
		u.Student = &s
	} else {
		t := models.Teacher{}
		u.Teacher = &t
	}

	if result := d.db.Create(&u); result.Error != nil {
		return nil, result.Error
	}

	return &u, nil
}

func (d *DB) GetUser(id uint) (*models.User, []models.UserCompanies, error) {
	res := models.User{
		Id: id,
	}

	if query := d.db.
		Preload("Teacher").
		Preload("Student").
		Preload("Admin").
		Preload("Technologies").
		Preload("Companies").
		Find(&res); query.Error != nil {
		return nil, nil, query.Error
	}

	uc := []models.UserCompanies{}
	if query := d.db.Find(&uc).Where("user_id = ?", id); query.Error != nil {
		return nil, nil, query.Error
	}

	return &res, uc, nil
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

func (d *DB) DeleteCompany(company models.Company) error {
	if company.Id != 0 {
		if result := d.db.Delete(&models.Company{}, company.Id); result.Error != nil {
			return result.Error
		}
	} else {
		if result := d.db.Where("title = ?", company.Title).Delete(&models.Company{}); result.Error != nil {
			return result.Error
		}
	}

	return nil
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

func (d *DB) DeleteProgramm(programm models.Programm) error {
	if programm.Id != 0 {
		if result := d.db.Delete(&models.Programm{}, programm.Id); result.Error != nil {
			return result.Error
		}
	} else {
		if result := d.db.Where("title = ?", programm.Title).Delete(&models.Programm{}); result.Error != nil {
			return result.Error
		}
	}

	return nil
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

func (d *DB) DeleteRole(role models.Role) error {
	if role.Id != 0 {
		if result := d.db.Delete(&models.Programm{}, role.Id); result.Error != nil {
			return result.Error
		}
	} else {
		if result := d.db.Where("title = ?", role.Title).Delete(&models.Programm{}); result.Error != nil {
			return result.Error
		}
	}

	return nil
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

func (d *DB) DeleteTechnology(tech models.Techonology) error {
	if tech.Id != 0 {
		if result := d.db.Delete(&models.Techonology{}, tech.Id); result.Error != nil {
			return result.Error
		}
	} else {
		if result := d.db.Where("title = ?", tech.Title).Delete(&models.Techonology{}); result.Error != nil {
			return result.Error
		}
	}

	return nil
}
