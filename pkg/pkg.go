package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func CreateDir(dir string) error {
	newpath := filepath.Join(".", dir)
	err := os.MkdirAll(newpath, os.ModePerm)
	return err
}

func RmDir(dir string) error {
	path := filepath.Join(".", dir)
	err := os.RemoveAll(path)
	return err
}

func StatDir(dir string) (os.FileInfo, error) {
	path := filepath.Join(".", dir)
	return os.Stat(path)
}

func Write(file string, data []byte, perm os.FileMode) error {
	err := ioutil.WriteFile(file, data, perm)

	return err
}

func Cmd(file string) {

	tcmd := `go get github.com/axw/gocov/gocov \
go get -u github.com/mchirico/date/parse \
go get gopkg.in/yaml.v2 \
go get github.com/spf13/cobra \
go get github.com/mitchellh/go-homedir \
go get github.com/spf13/viper 

`
	tcmd = "go"

	scmd := fmt.Sprintf("export GOPATH=./%s;%s", file,tcmd)

	fmt.Println(scmd)

	out, err := exec.Command("/bin/bash", "-c", scmd).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s, %s\n", out,tcmd)

}
