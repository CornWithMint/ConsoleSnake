package main

import (
	"encoding/json"
	"os"

	"github.com/google/uuid"
)

type User struct {
	PersonalID string `json:"PersonalId"`
	Score      int    `json:"Score"`
}

func NewUser() *User {
	return &User{
		PersonalID: GenerateUUID(),
		Score:      0,
	}
}

func GenerateUUID() string {
	UserUUID := uuid.New().String()
	return UserUUID
}

func Open() *os.File {
	file, err := os.Open("Score.json")
	if err != nil {
		// обработка
	}
	defer file.Close()
	file, err = os.Create("Score.json")
	if err != nil {
		// ...
	}
	defer file.Close()
	return file
}
func Encode(user *User, file *os.File) {
	enc := json.NewEncoder(file)
	enc.Encode(user)
}

func Decode(user *User, file *os.File) *User {
	dec := json.NewDecoder(file)
	dec.Decode(&user)
	return user
}

func GetBestScore(user *User) int {
	file := Open()
	Decode(user, file)
	return user.Score
}
