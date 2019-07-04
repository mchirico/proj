package cmd

import (
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/process"
	"github.com/mchirico/tlib/util"
	"strings"
	"testing"
)

func TestCreateFile(t *testing.T) {

	defer util.NewTlib().ConstructDir()()

	proj := []string{"junkTest"}

	cp := process.Cp{}

	cp.CreateProject(proj)

	data, _ := fileStrings.GetPath("./junkTest/setpath")
	match := strings.Contains(string(data), "git clone git@github.")
	if !match {
		t.Fatalf("No data...")
	}

	flist := util.ListFiles(util.PWD())
	if len(flist) < 15 {
		t.Fatalf("Did not create files...")
	}

}
