package pkg

// Ref: https://www.youtube.com/watch?v=oeFrGpAjm8s
import (
	"testing"
)

func TestCreateDir(t *testing.T) {

	dir := "public"

	err := CreateDir(dir)
	if err != nil {
		t.Fatalf("Can not create...")
	}

	stat, err := StatDir(dir)
	if stat.Name() != dir {
		t.Fatalf("Did not create directory: %s\n", dir)
	}

	err = RmDir("public")
	if err != nil {
		t.Fatalf("Can not create...")
	}

	stat, err = StatDir(dir)
	if err == nil {
		t.Fatalf("Should have deleted directory: %s\n", dir)
	}

}

