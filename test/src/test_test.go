package test

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
