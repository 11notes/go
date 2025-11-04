package eleven

import (
	"fmt"
	"time"
	"regexp"
	"os"

	"github.com/11notes/go/util"
	"github.com/11notes/go/container"
	"github.com/11notes/go/http"
)

type New struct{
	Util util.Util
	Container container.Container
	HTTP http.HTTP
}

// output log in json format with time stamp and simple string message
func (c *New) Log(t string, m string, args ...interface{}){
	l := "INF"
	p := true
	switch {
		case regexp.MustCompile(`(?i)DEBUG|DBG|DEB`).MatchString(t): l = "DBG"
		case regexp.MustCompile(`(?i)INFO|INF|IN`).MatchString(t): l = "INF"
		case regexp.MustCompile(`(?i)WARNING|WARN|WRN`).MatchString(t): l = "WRN"
		case regexp.MustCompile(`(?i)ERROR|ERR`).MatchString(t): l = "ERR"
		case regexp.MustCompile(`(?i)START`).MatchString(t): m = fmt.Sprintf("starting %s v%s", os.Getenv("APP_NAME"), os.Getenv("APP_VERSION"))
		case regexp.MustCompile(`(?i)PATCH|FIX`).MatchString(t): l = "FIX"
	}
	if(l == "DBG"){
		if _, ok := os.LookupEnv("DEBUG"); !ok {
			p = false
		}
	}
	if(p){
		fmt.Println(fmt.Sprintf(`{"time":"%s","type":"%s","msg":"%s"}`, time.Now().Format("2006-01-02T15:04:05.000Z"), l, fmt.Sprintf(m, args...)))
	}
}

// output log in json format with time stamp and simple string message and exist process with exit code 1
func (c *New) LogFatal(m string, args ...interface{}){
	c.Log("ERR", m, args...)
	os.Exit(1)
}