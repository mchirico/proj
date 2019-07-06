package fileStrings

import (
	"io/ioutil"
	"strings"
)

var readme = []byte(`


[![Build Status](https://travis-ci.org/mchirico/{proj}.svg?branch=master)](https://travis-ci.org/mchirico/{proj})
[![codecov](https://codecov.io/gh/mchirico/{proj}/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/{proj})

[![Build Status](https://mchirico.visualstudio.com/{proj}/_apis/build/status/mchirico.{proj}?branchName=master)](https://mchirico.visualstudio.com/{proj}/_build/latest?definitionId=9&branchName=master)


# {proj}



### Checklist:

1. dockerPassword
2. [CodeCov Token](https://codecov.io/gh/mchirico)
3. No Caps in project
4. MONGO_CONNECTION_STRING
5. MONGO_DATABASE 
6. Make Azure Boards Public
7. More Cobra commands? (cobra add say)



## Build with vendor
{rep}
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
{rep}


## Don't forget golint

{rep}

golint -set_exit_status $(go list ./... | grep -v /vendor/)

{rep}


`)

var recreateCobra = []byte(`#!/bin/bash

export GOPATH=$(pwd)
export PATH="$(pwd)/bin:$PATH"
export GOBIN="$(pwd)/bin"

go get github.com/axw/gocov/gocov
go install github.com/axw/gocov/gocov
go get -u github.com/mchirico/date/parse
go get gopkg.in/yaml.v2

go get github.com/spf13/cobra
go get github.com/mitchellh/go-homedir
go get github.com/spf13/viper

rm -rf src/github.com/mchirico/{proj}
cobra init github.com/mchirico/{proj}

proj create {proj}



cd src/github.com/mchirico/{proj}

export GO111MODULE=on
go mod init
go mod vendor
go test -v -mod=vendor ./...
go build -v -mod=vendor ./...

git init
git add cmd
git add *.go
git add .travis.yml
git add .gitignore
git add LICENSE
git add README.md
git add azure-pipelines.docker.yml
git add Dockerfile
git commit -m "first commit"
git remote add origin git@github.com:mchirico/{proj}.git
git push -u origin master --force


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

var notes = []byte(`


git rebase --root -i
git push origin	master --force


git init
git add README.md
git commit -m "first commit"
git remote add origin git@github.com:mchirico/{proj}.git
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
        - go get -v -t -d ./...
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
        - # bash <(curl -s https://codecov.io/bash)
        - # "./cc-test-reporter after-build --exit-code $TRAVIS_TEST_RESULT"


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

start.sh
Notes

`)

var azurePipeline = []byte(`# 
# 
# 
# 

pool:
  vmImage: 'ubuntu-16.04' # examples of other options: 'macOS-10.13', 'vs2017-win2016'

variables:
  GOBIN:  '$(GOPATH)/bin' # Go binaries path
  GOROOT: '/usr/local/go1.11' # Go installation path
  GOPATH: '$(system.defaultWorkingDirectory)/gopath' # Go workspace path
  GOMAXPROCS: 9
  modulePath: '$(GOPATH)/src/github.com/$(build.repository.name)' # Path to the module's code
  dockerId: aipiggybot  
  imageName: {proj}  # Replace with the name of the image you want to publish


steps:
- script: |
    mkdir -p '$(GOBIN)'
    mkdir -p '$(GOPATH)/pkg'
    mkdir -p '$(modulePath)'
    shopt -s extglob
    mv !(gopath) '$(modulePath)'
    echo '##vso[task.prependpath]$(GOBIN)'
    echo '##vso[task.prependpath]$(GOROOT)/bin'
    echo 'Variables:'
    echo ${MONGO_DATABASE}
  displayName: 'Set up the Go workspace'

- script: go get -v -t -d ./...
  workingDirectory: '$(modulePath)'
  displayName: 'go get dependencies'

- script: go build -v .
  workingDirectory: '$(modulePath)'
  displayName: 'Build'

- script: |
    set -e
    go test -race -coverprofile=coverage.txt -covermode=atomic ./...
    if [[ -s coverage.txt ]]; then bash <(curl -s https://codecov.io/bash); fi
  env:
    MONGO_DATABASE: $(MONGO_DATABASE)
    MONGO_CONNECTION_STRING: $(MONGO_CONNECTION_STRING)
    CODECOV_TOKEN: $(CODECOV_TOKEN)
  workingDirectory: '$(modulePath)'
  displayName: 'Run tests'


# Docker
- script: |
    set -e
    docker build --no-cache -t $(dockerId)/$(imageName) .
    echo "${DOCKERPASSWORD}"| docker login -u=$(dockerId) --password-stdin
    docker push $(dockerId)/$(imageName)
    docker logout
  env:
    DOCKERPASSWORD: $(dockerPassword)
  workingDirectory: '$(modulePath)'
  displayName: 'Building docker image and pushing'

`)

var azureDocker = []byte(`FROM golang:1.12.6-alpine3.10 AS build

RUN apk add --no-cache git


WORKDIR /go/src/project

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/project/
RUN go get -v -t -d ./...
RUN go build -o /bin/project

# This results in a single layer image
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /bin/project /bin/project
ENTRYPOINT ["/bin/project"]
CMD ["--help"]

`)

func RecreateCobra(proj string) []byte {
	return []byte(strings.Replace(string(recreateCobra), "{proj}", proj, -1))
}

func Readme(proj string) []byte {

	t := strings.Replace(string(readme), "{rep}", "```", -1)
	t = strings.Replace(t, "{proj}", proj, -1)

	return []byte(t)
}

func Notes(proj string) []byte {
	t := strings.Replace(string(notes), "{proj}", proj, -1)
	return []byte(t)

}

func GetPath(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}

func AzurePipeline(proj string) []byte {
	return []byte(strings.Replace(string(azurePipeline), "{proj}", proj, -1))
}

func AzureDocker(proj string) []byte {
	return []byte(strings.Replace(string(azureDocker), "{proj}", proj, -1))
}
