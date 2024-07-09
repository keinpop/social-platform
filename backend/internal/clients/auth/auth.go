package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Auth struct {
	config *Config
}

func NewAuth(cfg *Config) *Auth {
	return &Auth{
		config: cfg,
	}
}

func (a *Auth) Register(login, password string) error {
	postBody, _ := json.Marshal(map[string]string{
		"login":    login,
		"password": password,
	})

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(a.getURL()+"/register", "application/json", responseBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func (a *Auth) CheckToken(token string) (bool, error) {
	req, err := http.NewRequest("GET", a.getURL()+"/check", nil)
	if err != nil {
		return false, err
	}
	req.Header.Set("Authorization", token)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		// var Resp struct {
		// 	Login string `json:"login"`
		// }

		// Прочитать респ боди в респ

		// Положить логин в контекст джина

		// Написать отдельную middleware, которая достает логин из контекста джина
		// и проверяет, что человек с данным майлом имеет роль админа
		// тогда доходим до c.Next()

		//

		return true, nil
	}

	return false, nil
}

func (a *Auth) getURL() string {
	return fmt.Sprintf("http://%s:%d", a.config.Host, a.config.Port)
}
