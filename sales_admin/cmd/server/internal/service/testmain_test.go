package service_test

import (
	"flag"
	"os"
	"testing"
)

var templatesGlob = "../../../../web/templates/*"

func TestMain(m *testing.M) {
	flag.StringVar(&templatesGlob, "template-glob", "../../../../web/templates/*", "the file path to the templates that will be used by the service")
	flag.Parse()
	os.Exit(m.Run())
}
