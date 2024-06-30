package sample

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	// Loading in .env file.
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	filepath := path.Join(basepath, "../../.env")

	// In case we are testing, try to load .env using relative path and ignore the error.
	err := godotenv.Load(filepath)
	if err != nil && !strings.HasSuffix(err.Error(), "no such file or directory") {
		panic(err)
	}
}

func ConfigMissingErr(key string) error {
	return fmt.Errorf("ConfigurationError: Missing %s", key)
}

func TryGet(key string) (string, bool) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return "", false
	}

	return value, true
}

func TryGetInt(key string) (int, bool) {
	value, ok := TryGet(key)
	if !ok {
		return 0, false
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue, true
}

func TryGetBool(key string) (bool, bool) {
	value, ok := TryGet(key)
	if !ok {
		return false, false
	}

	bValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return bValue, true
}

// Get - This function will attempt to get an environment variable, or panic if the variable is not populated.
func Get(key string) string {
	value, ok := TryGet(key)
	if !ok {
		panic(ConfigMissingErr(key))
	}
	return value
}

func GetInt(key string) int {
	value, ok := TryGetInt(key)
	if !ok {
		panic(ConfigMissingErr(key))
	}

	return value
}
