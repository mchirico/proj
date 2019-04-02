package cmd

import (
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/process"
	"github.com/mchirico/proj/tlib"
	"strings"
	"testing"
)

func TestCreateFile(t *testing.T) {

	defer tlib.ConstructDir()()

	proj := []string{"junkTest"}

	cp := process.Cp{}

	cp.CreateProject(proj)

	data, _ := fileStrings.GetPath("./junkTest/setpath")
	match := strings.Contains(string(data), "git clone git@github.")
	if !match {
		t.Fatalf("No data...")
	}

	tlib.Rmdir(proj[0])

}
