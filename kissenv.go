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
		log.Debugf(
      fmt.Sprintf(
        "%s ENV VAR exists",
        key,
      ),
    )
	}

	return value

}

// LoadEnvFile returns
func LoadEnvFile(filename string) {
  err := godotenv.Load(filename)

  if err != nil {
    log.Debugf(
      fmt.Sprintf(
				"Cannot load %s file, default ENV VAR values will be used",
        envFile,
      ),
    )
  }

}
