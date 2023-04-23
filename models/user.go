package models

import (
	"errors"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserNotFound    = errors.New("user not found")
	InvalidPassword = errors.New("password did not match")
)

func RegisterUser(username string, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	return redisClient.Set("user:"+username, hash, 0).Err()
}

func AuthenticateUser(username, password string) error {
	hash, err := redisClient.Get("user:" + username).Bytes()
	if err == redis.Nil {
		return UserNotFound
	} else if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return InvalidPassword
	}
	return nil
}
