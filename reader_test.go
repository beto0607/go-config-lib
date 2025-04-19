package goconfiglib_test

import (
	"goconfiglib"
	"testing"
)

func TestMissingFile(t *testing.T) {
	_, err := goconfiglib.LoadConfigs("missing_file.ini", goconfiglib.Settings{
		UseXDGConfigHome: false,
	})
	if err == nil {
		t.Error("There was an error")
	}
}

func TestLoadConfigs(t *testing.T) {
	configs, err := goconfiglib.LoadConfigs("test-files/test.ini", goconfiglib.Settings{
		UseXDGConfigHome: false,
	})

	if err != nil {
		t.Error("There was an error")
	}

	if len(configs.Root.Subsections) != 5 {
		t.Errorf("Wrong number of Subsections. Got %d", len(configs.Root.Subsections))
	}

	configs.Print()
}
