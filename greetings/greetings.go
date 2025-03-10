package greetings

import (
	"errors"

	"fmt"
	"math/rand"
)

func Hello(name string)(string, error){

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormats(), name)

	return message, nil
}

func randomFormats()string {

	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! We meet!",
	}

	return formats[rand.Intn(len(formats))]

}