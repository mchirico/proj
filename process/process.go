package process

import (
	"errors"
	"fmt"
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/pkg"
	"os"
	"os/exec"
	"path/filepath"
)

type Cp struct {
	Force bool
	Empty bool
}


func (cp *Cp) CreateProject(args []string) error {

	if len(args) == 1 {

		_, err := pkg.StatDir(args[0])
		if err == nil {
			if cp.Force != true {
				return errors.New("Directory Exists... remove first.")
			}
		}

		pkg.CreateDir(args[0])

		pkg.Write(args[0]+"/setpath", fileStrings.Setpath, 0700)

		pkg.CreateDir(args[0] + "/src")
		pkg.CreateDir(args[0] + "/src/github.com")

		if cp.Empty == true {
			return nil
		}

		pkg.CreateDir(args[0] + "/src/github.com/mchirico")
		pkg.CreateDir(args[0] + "/src/github.com/mchirico/" + args[0])

		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/.travis.yml", fileStrings.Travis, 0600)

		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/Notes", fileStrings.Notes, 0600)
		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/README.md", fileStrings.Readme(), 0600)
		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/.gitignore", fileStrings.GitIgnore, 0600)

		pkg.Write(args[0]+"/src/github.com/mchirico/"+args[0]+"/start.sh", fileStrings.StartSh, 0700)

	}

	return nil
}

// Not implemented
func InstallGoPackages(proj string) {

	absPath, _ := filepath.Abs(proj)

	gopath := fmt.Sprintf("GOPATH=%s", absPath)
	path := fmt.Sprintf("PATH=\"%s/bin:$PATH\"", absPath)
	gobin := fmt.Sprintf("GOBIN=\"%s/bin\"", absPath)

	cmd := exec.Command("/bin/bash", ".", "junkTest/setpath")
	cmd.Env = append(os.Environ(),
		gopath,
		path,
		gobin,
	)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output)) // write the output with ResponseWriter

}
