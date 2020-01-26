package process

import (
	"errors"
	"fmt"
	"github.com/mchirico/proj/fileStrings"
	"github.com/mchirico/proj/pkg"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Cp struct {
	Force   bool
	Empty   bool
	baseDir string
	projDir string
}

func (cp *Cp) Analyze(dir string) error {
	path, _ := filepath.Abs(".")
	list := strings.Split(path, "/")

	cp.baseDir = dir
	cp.projDir = dir

	if list[len(list)-1] == dir {
		cp.baseDir = "."
	}

	if dir == "." {
		return errors.New("Can't work with .")
	}

	return nil
}

func (cp *Cp) CreateProject(args []string) error {

	if len(args) == 1 {
		cp.Analyze(args[0])
		_, err := pkg.StatDir(args[0])
		if err == nil {
			if cp.Force != true {
				return errors.New("Directory Exists... remove first.")
			}
		}

		if cp.baseDir != "." {
			pkg.CreateDir(cp.baseDir)
		}

		pkg.Write(cp.baseDir+"/setpath", fileStrings.Setpath, 0700)
		pkg.Write(cp.baseDir+"/recreateWithCobra.sh", fileStrings.RecreateCobra(cp.projDir), 0700)

		pkg.CreateDir(cp.baseDir + "/src")
		pkg.CreateDir(cp.baseDir + "/src/github.com")

		if cp.Empty == true {
			return nil
		}

		pkg.CreateDir(cp.baseDir + "/src/github.com/mchirico")
		pkg.CreateDir(cp.baseDir + "/src/github.com/mchirico/" + cp.projDir)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/.travis.yml",
			fileStrings.Travis, 0644)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/Notes",
			fileStrings.Notes(cp.projDir), 0644)
		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/README.md",
			fileStrings.Readme(cp.projDir), 0644)
		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/.gitignore",
			fileStrings.GitIgnore, 0600)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/start.sh",
			fileStrings.StartSh, 0700)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/azure-pipelines.docker.yml",
			fileStrings.AzurePipeline(cp.projDir), 0644)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/Dockerfile",
			fileStrings.AzureDocker(cp.projDir), 0644)

		pkg.Write(cp.baseDir+"/src/github.com/mchirico/"+cp.projDir+"/Makefile",
			fileStrings.Makefile(cp.projDir), 0644)

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
