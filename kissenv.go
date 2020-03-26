package kissenv

import (
	"os"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

var envFile = ".env"

// GetEnv returns
func GetEnv(key string) string  {
  LoadEnvFile(envFile)

	value, exists := os.LookupEnv(key)

	if exists {
		log.Print(
      fmt.Sprintf(
        "%s env var exists",
        envFile,
      ),
    )
	}

	return value

}

// LoadEnvFile returns
func LoadEnvFile(filename string) {
  err := godotenv.Load(filename)

  if err != nil {
    log.Print(
      fmt.Sprintf(
        "Error loading %s file",
        envFile,
      ),
    )
  }

}
