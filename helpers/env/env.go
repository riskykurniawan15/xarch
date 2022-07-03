package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadEnv(file string) error {
	return godotenv.Load(file)
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
