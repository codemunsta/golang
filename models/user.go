package models

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var (
	UserNotFound    = errors.New("user not found")
	InvalidPassword = errors.New("password did not match")
)

type User struct {
	key      string
	username string
	password []byte
}

func NewUser(username string, password []byte) (*User, error) {
	id, err := redisClient.Incr("user:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	pipeline := redisClient.Pipeline()
	pipeline.HSet(key, "id", id)
	pipeline.HSet(key, "username", username)
	pipeline.HSet(key, "hash", password)
	pipeline.HSet("user:by-username", username, id)
	_, err2 := pipeline.Exec()
	if err2 != nil {
		return nil, err
	}
	return &User{key, username, password}, nil
}

func (user *User) GetUserName() (string, error) {
	return redisClient.HGet(user.key, "username").Result()
}

func (user *User) GetPassword() ([]byte, error) {
	return redisClient.HGet(user.key, "hash").Bytes()
}

func (user *User) Authenticate(password string) error {
	hash, err := user.GetPassword()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return InvalidPassword
	} else {
		return err
	}
}

func GetUserByUserName(username string) (*User, error) {
	id, err := redisClient.HGet("user:by-username", username).Int64()
	if err == redis.Nil {
		return nil, UserNotFound
	} else if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("user:%d", id)
	return &User{key: key}, nil
}

func RegisterUser(username string, password string) error {
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return err
	}
	_, err2 := NewUser(username, hash)
	return err2
	// return redisClient.Set("user:"+username, hash, 0).Err()
}

func AuthenticateUser(username, password string) error {
	user, err := GetUserByUserName(username)
	if err != nil {
		return err
	}
	return user.Authenticate(password)
	//hash, err := redisClient.Get("user:" + username).Bytes()
	//if err == redis.Nil {
	//	return UserNotFound
	//} else if err != nil {
	//	return err
	//}
	//err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	//if err != nil {
	//	return InvalidPassword
	//}
	//return nil
}
