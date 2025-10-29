// analog to docker-util for shell and javascript in go
package util

import (
	"os"
	"os/exec"
	"syscall"
	"io/ioutil"
	"strings"
	"bufio"
	"io"
	"errors"
)

type Util struct{}

// reads a file if it exists and returns the content of the file
func (c *Util) ReadFile(path string) (string, error){
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// writes contents to a file
func (c *Util) WriteFile(path string, txt string) error{
	err := ioutil.WriteFile(path, []byte(txt), os.ModePerm)	
	if err != nil {
		return err
	}
	return nil
}

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

// checks if a file containing an environment variable exists and if not assigns a default value
func (c *Util) GetenvFile(path string, fallback string) string{
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		value, err := c.ReadFile(path)
		if err != nil {
			return fallback
		}
		return value
	}
	return fallback
}

// run an external program and return output
func (c *Util) Run(bin string, params []string) (string, error){
	cmd := exec.Command(bin, params...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid:true}

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	out := []string{}
	go func() {
		stdoutScanner := bufio.NewScanner(io.MultiReader(stdout,stderr))
		for stdoutScanner.Scan() {
			out = append(out, stdoutScanner.Text())
		}
	}()

	err := cmd.Start()
	if err != nil {
		return "", errors.New(err.Error() + strings.Join(out, " "))
	}
	err = cmd.Wait()
	if err != nil {
		return "", errors.New(err.Error() + strings.Join(out, " "))
	}

	return strings.Join(out, " "), nil
}