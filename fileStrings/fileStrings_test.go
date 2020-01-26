package fileStrings

import (
	"fmt"
	"strings"
	"testing"
)

func TestSetPath(t *testing.T) {
	expected := "WORKDIR $GOPATH/src/github.com/mchirico/test"

	r := AzureDocker("test")
	fmt.Printf("stuff:  %s\n\n", string(r))
	if !strings.Contains(string(r), expected) {
		t.FailNow()
	}

}
