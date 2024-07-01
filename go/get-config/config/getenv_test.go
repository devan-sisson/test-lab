package config

import (
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

func prepareEnv() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	filepath := path.Join(basepath, ".env")

	// In case we are testing, try to load .env using relative path and ignore the error.
	err := godotenv.Load(filepath)
	if err != nil && !strings.HasSuffix(err.Error(), "no such file or directory") {
		panic(err)
	}
}

func TestTryGet(t *testing.T) {
	prepareEnv()

	var tests = []struct {
		envVar string
		want   any
	}{
		{"TEST_One", "one"},
		{"TEST_TWO", "two"},
		{"test_three", "three"},
	}

	for _, tt := range tests {
		t.Run(tt.envVar, func(t *testing.T) {
			got, _ := TryGet(tt.envVar)

			if got != tt.want {
				t.Errorf("got %s, want %s", got, tt.want)
			}
		})
	}
}

func TestTryGetBool(t *testing.T) {
	prepareEnv()

	var tests = []struct {
		envVar string
		want   bool
	}{
		{"TEST_BOOL1", true},
		{"TEST_BOOL2", false},
		{"TEST_Bool3", true},
		{"TEST_bool4", false},
	}

	for _, tt := range tests {
		t.Run(tt.envVar, func(t *testing.T) {
			got, _ := TryGetBool(tt.envVar)

			if got != tt.want {
				t.Errorf("got %t, want %s", got, strconv.FormatBool(tt.want))
			}
		})
	}
}

func TestTryGetInt(t *testing.T) {
	prepareEnv()

	var tests = []struct {
		envVar string
		want   int
	}{
		{"TEST_One1", 1},
		{"TEST_TWO2", 2},
		{"TEST_three3", 3},
	}

	for _, tt := range tests {
		t.Run(tt.envVar, func(t *testing.T) {
			got, _ := TryGetInt(tt.envVar)

			if got != tt.want {
				t.Errorf("got %d, want %s", got, strconv.Itoa(tt.want))
			}
		})
	}
}
