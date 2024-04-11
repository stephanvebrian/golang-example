package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type User struct {
	Username string
	Password string
	Salt     string
}

// utility function, dont have any to do with the example,
// what you need to know is it will generating random string based on the length we pass
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var result string
	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result += string(charset[index.Int64()])
	}
	return result
}

func generatingUserSalt(user *User) {
	user.Salt = generateRandomString(15)
}

func main() {
	inputUser := User{
		Username: "admin",
		Password: "this-is-password",
	}

	fmt.Printf("User salt: %s\n", inputUser.Salt)
	generatingUserSalt(&inputUser)
	fmt.Printf("user salt generated: %s\n", inputUser.Salt)
}
