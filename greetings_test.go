package greetings

import (
	"strings"
	"testing"
)

func TestGreetings(t *testing.T) {

	var usuario = "UnNombre"

	saludo, err := Hello(usuario)

	if saludo == "" {
		t.Fatalf("No cumple con el resultado esperado para un paramatro")

	}

	if !strings.Contains(saludo, "UnNombre") {
		t.Fatalf("No cumple con el resultado esperado para un paramatro")

	}

	if err != nil {
		t.Fatalf("No cumple con el resultado esperado para un paramatro")

	}

}
