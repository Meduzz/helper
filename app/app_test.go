package app_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/Meduzz/helper/app"
)

type (
	goodConfig struct {
		Name string `json:"name"`
		Age  int    `json:"age" env:"AGE"`
	}

	badConfig struct {
		Name string `json:"name"`
		Age  int    `json:"age" env:"AGE"`
	}
)

var demand = fmt.Errorf("I demand ... ")

func (g *goodConfig) Start() error {
	return nil
}

func (b *badConfig) Start() error {
	return demand
}

func TestHappyCase(t *testing.T) {
	config := &goodConfig{}
	err := app.Initiate("test_config.json", config)

	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if config.Name != "Test Testsson" {
		t.Errorf("Name was not Test Testsson but %s\n", config.Name)
		t.FailNow()
	}

	if config.Age != 100 {
		t.Errorf("Age was not 100 but %d\n", config.Age)
		t.FailNow()
	}
}

func TestSadCase(t *testing.T) {
	config := &badConfig{}
	err := app.Initiate("test_config.json", config)

	if err == nil {
		t.Error("expected an error")
		t.FailNow()
	}

	if !errors.Is(err, demand) {
		t.Errorf("error was not %v but %v\n", demand, err)
		t.FailNow()
	}
}
