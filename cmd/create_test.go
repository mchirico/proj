package cmd

import (
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/process"
	"github.com/mchirico/tlib/util"
	"strings"
	"testing"
)

func FileIn(file, pwd string) bool {
	files := util.ListFiles(pwd)
	for _, value := range files {

		if strings.Contains(value, file) {
			return true
		}
	}
	return false
}

func TestCreateFile(t *testing.T) {

	t.Parallel()

	defer util.NewTlib().ConstructDir()()

	proj := []string{"junkTest"}

	cp := process.Cp{}

	cp.CreateProject(proj)

	data, _ := fileStrings.GetPath("./junkTest/setpath")
	match := strings.Contains(string(data), "git clone git@github.")
	if !match {
		t.Fatalf("No data...")
	}

	tests := []struct {
		name string
		file string
	}{
		{name: "test Dockerfile", file: "Dockerfile"},
		{name: "test Azure pipeline", file: "azure-pipelines.docker.yml"},
		{name: "test travis", file: ".travis.yml"},
		{name: "test README.md", file: "README.md"},
		{name: "test recreateWithCobra.sh", file: "recreateWithCobra.sh"},
		{name: "test setpath", file: "setpath"},
		{name: "test .gitignore", file: ".gitignore"},
		{name: "test start.sh", file: "start.sh"},
		{name: "test github.com", file: "github.com"},
		{name: "test src", file: "src"},
		{name: "test mchirico", file: "mchirico"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if !FileIn(tc.file, util.PWD()) {
				t.Fatalf("can't find: %s\n", tc.file)
			}

		})
	}

}
