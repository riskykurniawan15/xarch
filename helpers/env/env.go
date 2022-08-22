package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv(file string) error {
	return godotenv.Load(file)
}

func LoadFlags() {
	for i := 0; i < len(os.Args); i++ {
		if os.Args[i][0] == '$' {
			value := strings.Split(strings.Replace(os.Args[i], "$", "", -1), "=")
			if len(value) == 2 {
				os.Setenv(value[0], value[1])
			}
		}
	}
}

func GetString(val string) string {
	return os.Getenv(val)
}

func GetInt(val string) int {
	result, err := strconv.Atoi(os.Getenv(val))
	if err != nil {
		return 0
	}
	return result
}

func GetFloat(val string) float64 {
	result, err := strconv.ParseFloat((os.Getenv(val)), 32)
	if err != nil {
		return 0.0
	}
	return result
}

func GetBoolean(val string) bool {
	result, err := strconv.ParseBool(os.Getenv(val))
	if err != nil {
		return false
	}
	return result
}
