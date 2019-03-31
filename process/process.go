package process

import (
	"errors"
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/pkg"
)

func CreateProject(args []string) error {

	if len(args) == 1 {

		_, err := pkg.StatDir(args[0])
		if err == nil {
			return errors.New("Directory Exists... remove first.")
		}

		pkg.CreateDir(args[0])

		pkg.Write(args[0]+"/setpath", fileStrings.Setpath, 0700)

		pkg.CreateDir(args[0] + "/src")
		pkg.CreateDir(args[0] + "/src/github.com")
		pkg.CreateDir(args[0] + "/src/github.com/mchirico")
		pkg.CreateDir(args[0] + "/src/github.com/mchirico/" + args[0])

		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/.travis.yml", fileStrings.Travis, 0600)

		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/Notes", fileStrings.Notes, 0600)
		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/README.md", fileStrings.Readme(), 0600)

	}

	return nil
}
