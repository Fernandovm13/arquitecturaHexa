package entities

import (
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Category struct {
	ID     int32          `json:"id"`
	Name   string         `json:"name"`
	Secret sql.NullString `json:"-"`
}

type CategoryDTO struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

func NewCategory(name string, secret string) *Category {
	return &Category{
		Name: name,
		Secret: sql.NullString{
			String: secret,
			Valid:  secret != "",
		},
	}
}

func (c *Category) EncryptSecret() error {
	if !c.Secret.Valid || c.Secret.String == "" {
		return errors.New("el secreto no puede estar vacío")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(c.Secret.String), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	c.Secret.String = string(hash)
	c.Secret.Valid = true
	return nil
}

func (c *Category) CompareSecret(plainSecret string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Secret.String), []byte(plainSecret))
	return err == nil
}
