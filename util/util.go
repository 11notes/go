// analog to docker-util for shell and javascript in go
package util

import (
	"os"
)

type Util struct{}

// checks if the command line argument exists (case-sensitive)
func (c *Util) CommandLineArgumentExists(f string) bool{
	if(len(os.Args) > 1){
		for _, a := range os.Args[1:] {
			if(f == a){
				return true
			}
		}
	}

	return false
}

// checks if an environment variable exists and if not assigns a default value
func (c *Util) Getenv(key string, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}