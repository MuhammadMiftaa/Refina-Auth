package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"time"
	"unicode"

	"refina-auth/config/env"
	"refina-auth/internal/types/dto"
	"refina-auth/internal/types/model"
	"refina-auth/internal/utils/data"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
)

func EmailValidator(str string) bool {
	email_validator := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return email_validator.MatchString(str)
}

func PasswordValidator(str string) (bool, bool, bool) {
	var hasLetter, hasDigit, hasMinLen bool
	for _, char := range str {
		switch {
		case unicode.IsLetter(char):
			hasLetter = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if len(str) >= 8 {
		hasMinLen = true
	}

	return hasMinLen, hasLetter, hasDigit
}

func PasswordHashing(str string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func ConvertToResponseType(data interface{}) interface{} {
	return dto.UsersResponse{
		ID:    data.(model.Users).ID.String(),
		Name:  data.(model.Users).Name,
		Email: data.(model.Users).Email,
	}
}

func GenerateToken(ID string, username string, email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"id":       ID,
		"username": username,
		"email":    email,
		"exp":      expirationTime.Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(env.Cfg.Server.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(jwtToken string) (dto.UserData, error) {
	token, _ := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("parsing token error occured")
		}
		return []byte(env.Cfg.Server.JWTSecretKey), nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return dto.UserData{}, errors.New("token is invalid")
	}

	return dto.UserData{
		ID:       claims["id"].(string),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
	}, nil
}

func ComparePass(hashPassword, reqPassword string) bool {
	hash, pass := []byte(hashPassword), []byte(reqPassword)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func StorageIsExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func GetGoogleOAuthConfig() (*oauth2.Config, string, error) {
	googleOauthConfig := &oauth2.Config{
		ClientID:     env.Cfg.OAuth.Google.GOClientID,
		ClientSecret: env.Cfg.OAuth.Google.GOClientSecret,
		RedirectURL:  "http://localhost:" + env.Cfg.Client.Port + "/v1/auth/callback/google",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	if env.Cfg.Server.Mode == data.STAGING_MODE || env.Cfg.Server.Mode == data.PRODUCTION_MODE {
		googleOauthConfig.RedirectURL = env.Cfg.Client.Url + "/v1/auth/callback/google"
	}

	return googleOauthConfig, env.Cfg.Client.Url, nil
}

func GetGithubOAuthConfig() (*oauth2.Config, string, error) {
	githubOauthConfig := &oauth2.Config{
		ClientID:     env.Cfg.OAuth.Github.GHClientID,
		ClientSecret: env.Cfg.OAuth.Github.GHClientSecret,
		RedirectURL:  "http://localhost:" + env.Cfg.Client.Port + "/v1/auth/callback/github",
		Scopes: []string{
			"read:user",
			"user:email",
		},
		Endpoint: github.Endpoint,
	}

	if env.Cfg.Server.Mode == data.STAGING_MODE || env.Cfg.Server.Mode == data.PRODUCTION_MODE {
		githubOauthConfig.RedirectURL = env.Cfg.Client.Url + "/v1/auth/callback/github"
	}

	return githubOauthConfig, env.Cfg.Client.Url, nil
}

func GetMicrosoftOAuthConfig() (*oauth2.Config, string, error) {
	microsoftOauthConfig := &oauth2.Config{
		ClientID:     env.Cfg.OAuth.Microsoft.MSClientID,
		ClientSecret: env.Cfg.OAuth.Microsoft.MSClientSecret,
		RedirectURL:  "http://localhost:" + env.Cfg.Client.Port + "/v1/auth/callback/microsoft",
		Scopes: []string{
			"User.Read",
		},
		Endpoint: microsoft.AzureADEndpoint("common"),
	}

	if env.Cfg.Server.Mode == data.STAGING_MODE || env.Cfg.Server.Mode == data.PRODUCTION_MODE {
		microsoftOauthConfig.RedirectURL = env.Cfg.Client.Url + "/v1/auth/callback/microsoft"
	}

	return microsoftOauthConfig, env.Cfg.Client.Url, nil
}

func GenerateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
