package options

import (
	"log"
	"strings"

	"github.com/shipt/config/lib"
)

var ValidateOption = func(c *lib.Configuration) {
	var fails []string
	if !c.Has("AWS.AWSAccessKeyID") {
		fails = append(fails, `Error: "AWSAccessKeyID" has not been defined`)
	}

	if !c.Has("AWS.AWSSecretAccessKey") {
		fails = append(fails, `Error: "AWSAccessKey" has not been defined`)
	}

	if !c.Has("AWS.AWSRegion") {
		fails = append(fails, `Error: "AWSRegion" has not been defined`)
	}

	if !c.Has("RollbarToken") {
		fails = append(fails, `Error: "RollbarToken" has not been defined`)
	}

	if c.Get("ServiceName") == "your mom" {
		fails = append(fails, `"your mom" exceeds size threshold`)
	}

	if len(fails) > 0 {
		log.Fatalf(strings.Join(fails, "\n"))
	}
}
