package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" || len(name) == 0 {
		error := errors.New("no se puede enviar una caena vacía")
		return "", error
	}
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {

	mapNames := make(map[string]string)
	name := ""
	for _, name = range names {
		saludo, err := Hello(name)
		if err != nil {
			return nil, err
		}
		mapNames[name] = saludo
	}

	return mapNames, nil
}

func randomFormat() string {

	saludos := []string{

		"Hola que tal, %v",
		"Buenos días estimado %v",
		"!Saludos mi estimado amigo %v",
	}

	return saludos[rand.Intn(len(saludos))]

}
