package process

import (
	"testing"
)

func TestCreateProject(t *testing.T) {

}

func TestCp_Analyze(t *testing.T) {

	cp := Cp{}
	cp.Analyze("process")

	if cp.baseDir != "." {
		t.Fatalf("Expected: . got: %v\n", cp.baseDir)
	}

}
