package fileStrings

import (
	"io/ioutil"
	"strings"
)

var readme = []byte(`

## Build with vendor
{rep}
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./cmd/script
{rep}

`)

var Setpath = []byte(`#!/bin/bash

# Change Project Here:
export PROJ=$(basename $(pwd))

#
mkdir -p src/github.com/mchirico
mkdir -p bin

export GOPATH=$(pwd)
export PATH="$(pwd)/bin:$PATH"
export GOBIN="$(pwd)/bin"
if ! [ -x "$(command -v godep)" ]; then
    echo 'Note: godep is not installed.' >&2


fi


if [ -d "$PWD/src/github.com/mchirico/$PROJ" ]; then
    cd "$PWD/src/github.com/mchirico/$PROJ"
else
    echo -e '

       cd src/github.com/mchirico
       git clone git@github.com:mchirico/${PROJ}.git

 Also install:

       go get github.com/axw/gocov/gocov
       go install github.com/axw/gocov/gocov
       go get -u github.com/mchirico/date/parse
       go get gopkg.in/yaml.v2

       go get github.com/spf13/cobra
       go get github.com/mitchellh/go-homedir
       go get github.com/spf13/viper

'
echo "          cobra init github.com/mchirico/${PROJ}"
echo -e "

    . setpath
    cobra add serve
    cobra add config
    cobra add create -p 'configCmd'

"
echo -e "
git remote add origin https://github.com/mchirico/${PROJ}.git
git push -u origin master
"

fi


`)

var Notes = []byte(`


git rebase --root -i
git push origin	master --force


git remote add origin https://github.com/mchirico/${PROJ}.git
git push -u origin master


`)

var Travis = []byte(`dist: trusty
sudo: false
matrix:
  include:
    - language: go
      go:
        - 1.12.x
      env:
        - GOMAXPROCS=9
      os:
        - linux
      install:
        - go get github.com/axw/gocov/gocov
        - go install github.com/axw/gocov/gocov
        - go get -u github.com/mchirico/date/parse
        - go get gopkg.in/yaml.v2
        - go get golang.org/x/crypto/ssh
      before_install:
        #- openssl aes-256-cbc -k "$super_secret_password" -in fixtures/data.enc -out fixtures/data -d
        #- for i in $(ls fixtures/fixtures_*.enc); do openssl aes-256-cbc -k "$super_secret_password" -in ${i} -out ${i%%.*} -d; done
        - curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64
          > ./cc-test-reporter
        - chmod +x ./cc-test-reporter
        - "./cc-test-reporter before-build"
      script:
        - go test -race -v -coverprofile=c.out ./... && echo -e "\n\n\n ✅ SUCCESS \n\n"
        - gocov test ./... > cc.out
      after_success:
        - cp c.out coverage.txt
        - bash <(curl -s https://codecov.io/bash)
        - "./cc-test-reporter after-build --exit-code $TRAVIS_TEST_RESULT"


`)

var StartSh = []byte(`#!/bin/bash

# These are common packages

go get github.com/axw/gocov/gocov
go install github.com/axw/gocov/gocov
go get -u github.com/mchirico/date/parse
go get gopkg.in/yaml.v2



`)

var GitIgnore = []byte(`# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, build with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# User-specific stuff
.idea/**/workspace.xml
.idea/**/tasks.xml
.idea/**/usage.statistics.xml
.idea/**/dictionaries
.idea/**/shelf

# Generated files
.idea/**/contentModel.xml

# Sensitive or high-churn files
.idea/**/dataSources/
.idea/**/dataSources.ids
.idea/**/dataSources.local.xml
.idea/**/sqlDataSources.xml
.idea/**/dynamic.xml
.idea/**/uiDesigner.xml
.idea/**/dbnavigator.xml

# Gradle
.idea/**/gradle.xml
.idea/**/libraries


`)

func Readme() []byte {
	return []byte(strings.Replace(string(readme), "{rep}", "```", -1))
}

func GetPath(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}
