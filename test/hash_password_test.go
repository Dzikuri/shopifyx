package test

import (
	"fmt"
	"testing"

	"github.com/Dzikuri/shopifyx/internal/helper"
)

func TestHashPassword(t *testing.T) {
	password := "admindev"

	hashPassword, err := helper.HashPassword(password)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("original password: ", password)
	fmt.Println("Hash salt password: ", string(hashPassword))
	
	err = helper.ComparePassword(hashPassword, password)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hash salt not match")
	} else {
			
			fmt.Println("Hash salt password is match")
	}
		
}

func TestComparePassword(t *testing.T) {
	hashPassword := "$2a$08$P7.Bu2UT9LAgbp7.vS4RHOGZzI3awH9zAdfb7TiipJZtqNsNE2k7S"
	password := "admindev"
	fmt.Println("original password: ", password)
	fmt.Println("Hash salt password: ", string(hashPassword))

	err := helper.ComparePassword(hashPassword, password)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Hash salt not match")
	} else {
			fmt.Println("Hash salt password is match")
	}
}