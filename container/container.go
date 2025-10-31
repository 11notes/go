// analog to docker-util for shell and javascript in go
package container

import (
	"errors"
	"os"
	"io/ioutil"
	"strings"
)

type Container struct{}

// tries to get a secret either from environment variable or from a secrets file set by environment variable
func (c *Container) GetSecret(env string, envPath string) (string, error){
	if value, ok := os.LookupEnv(env); ok {
		return value, nil
	}else{
		if value, ok := os.LookupEnv(envPath); ok {
			bytes, err := ioutil.ReadFile(value)
			if err != nil {
				return "", err
			}
			return strings.TrimSpace(string(bytes)), nil
		}else{
			return "", errors.New(env + " and " + envPath + " do not exist!")
		}
	}
}

// merges default parameters and user set parameters
func (c *Container) Command(d []string) []string{
	if(len(os.Args) > 0){
		args := os.Args[1:]
		for _, value := range args{
			d = append(d, value)
		}
	}
	return(d)
}