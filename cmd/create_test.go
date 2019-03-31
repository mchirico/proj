package cmd

import (
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/pkg"
	"github.com/mchirico/proj/process"
	"testing"
)

func TestCreateFile(t *testing.T) {
	proj := []string{"junkTest"}
	process.CreateProject(proj)

	data, _ := fileStrings.GetPath("./junkTest/setpath")
	t.Logf("%s\n", data)

	pkg.RmDir(proj[0])

}
