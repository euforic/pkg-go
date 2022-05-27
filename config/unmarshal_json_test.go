package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestUnmatchedKeyInJsonConfigFile(t *testing.T) {
	type cfgStruct struct {
		Name string
	}
	type cfgFile struct {
		Name string
		Test string
	}
	cfg := cfgFile{Name: "test", Test: "ATest"}

	file, err := ioutil.TempFile("/tmp", "config")
	if err != nil {
		t.Fatal("Could not create temp file")
	}
	defer os.Remove(file.Name())
	defer file.Close()

	filename := file.Name()

	if err := json.NewEncoder(file).Encode(cfg); err == nil {

		var result cfgStruct

		// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
		if err := New(&Settings{}).Load(&result, filename); err != nil {
			t.Errorf("Should NOT get error when loading configuration with extra keys. Error: %v", err)
		}

		// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
		if err := New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename); err == nil || !strings.Contains(err.Error(), "json: unknown field") {

			t.Errorf("Should get unknown field error when loading configuration with extra keys. Instead got error: %v", err)
		}

	} else {
		t.Errorf("failed to marshal cfg")
	}

	// Add .json to the file name and test
	err = os.Rename(file.Name(), file.Name()+".json")
	if err != nil {
		t.Errorf("Could not add suffix to file")
	}
	filename = file.Name() + ".json"
	defer os.Remove(filename)

	var result cfgStruct

	// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
	if err := New(&Settings{}).Load(&result, filename); err != nil {
		t.Errorf("Should NOT get error when loading configuration with extra keys. Error: %v", err)
	}

	// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
	if err := New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename); err == nil || !strings.Contains(err.Error(), "json: unknown field") {

		t.Errorf("Should get unknown field error when loading configuration with extra keys. Instead got error: %v", err)
	}

}
