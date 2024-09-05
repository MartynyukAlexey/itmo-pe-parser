package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type AuthConfig struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`

	IdToken string `json:"id_token"`

	TokenType string `json:"token_type"`
	Scope     string `json:"scope"`
}

func (cfg *AuthConfig) Refresh() error {
	client := &http.Client{}
	data := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {cfg.RefreshToken},
		"scope":         {"openid profile"},
		"client_id":     {"student-personal-cabinet"},
	}
	req, err := http.NewRequest("POST", "https://id.itmo.ru/auth/realms/itmo/protocol/openid-connect/token", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var newAuthConfig AuthConfig
	err = json.NewDecoder(resp.Body).Decode(&newAuthConfig)
	if err != nil {
		return err
	}

	cfg.AccessToken = newAuthConfig.AccessToken
	cfg.ExpiresIn = newAuthConfig.ExpiresIn
	cfg.RefreshToken = newAuthConfig.RefreshToken
	cfg.RefreshExpiresIn = newAuthConfig.RefreshExpiresIn
	cfg.IdToken = newAuthConfig.IdToken

	return nil
}
