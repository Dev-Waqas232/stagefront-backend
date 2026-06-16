package config

import (
	"bufio"
	"os"
	"slices"
	"strings"
)

type Env struct {
	DB_URL string
	PORT   string
}

func LoadEnv() {
	path, err := os.Open(".env")

	if err != nil {
		return
	}

	buffer := bufio.NewScanner(path)

	for buffer.Scan() {
		textLine := buffer.Text()

		// Checking if the line is empty
		if strings.TrimSpace(textLine) == "" {
			continue
		}

		strSlice := strings.Split(textLine, "=")

		key, value := strSlice[0], slices.Concat(strSlice[1:])[0]
		os.Setenv(key, value)
	}
}

func LoadConfig() Env {
	LoadEnv()

	env := Env{DB_URL: os.Getenv("DB_URL"), PORT: (os.Getenv("PORT"))}

	return env

}
