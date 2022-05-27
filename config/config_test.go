package config

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type Anonymous struct {
	Description string
}

type testConfig struct {
	APPName string `default:"config" json:",omitempty"`
	Hosts   []string

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306" json:",omitempty"`
		SSL      bool   `default:"true" json:",omitempty"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}

	Anonymous `anonymous:"true"`

	private string
}

func generateDefaultConfig() testConfig {
	return testConfig{
		APPName: "config",
		Hosts:   []string{"http://example.org", "http://jinzhu.me"},
		DB: struct {
			Name     string
			User     string `default:"root"`
			Password string `required:"true" env:"DBPassword"`
			Port     uint   `default:"3306" json:",omitempty"`
			SSL      bool   `default:"true" json:",omitempty"`
		}{
			Name:     "config",
			User:     "config",
			Password: "config",
			Port:     3306,
			SSL:      true,
		},
		Contacts: []struct {
			Name  string
			Email string `required:"true"`
		}{
			{
				Name:  "Jinzhu",
				Email: "wosmvp@gmail.com",
			},
		},
		Anonymous: Anonymous{
			Description: "This is an anonymous embedded struct whose environment variables should NOT include 'ANONYMOUS'",
		},
		private: "",
	}
}

func TestLoadNormaltestConfig(t *testing.T) {
	cfg := generateDefaultConfig()
	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}

			var result testConfig
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(result, cfg) {
				t.Errorf("result should equal to original configuration")
			}
		}
	} else {
		t.Errorf("failed to marshal cfg")
	}
}

func TestLoadtestConfigFromTomlWithExtension(t *testing.T) {
	var (
		cfg    = generateDefaultConfig()
		buffer bytes.Buffer
	)

	if err := toml.NewEncoder(&buffer).Encode(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config.toml"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(buffer.Bytes()); err != nil {
				t.Error(err)
			}

			var result testConfig
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(result, cfg) {
				t.Errorf("result should equal to original configuration")
			}
		}
	} else {
		t.Errorf("failed to marshal cfg")
	}
}

func TestLoadtestConfigFromTomlWithoutExtension(t *testing.T) {
	var (
		cfg    = generateDefaultConfig()
		buffer bytes.Buffer
	)

	if err := toml.NewEncoder(&buffer).Encode(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(buffer.Bytes()); err != nil {
				t.Error(err)
			}

			var result testConfig
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(result, cfg) {
				t.Errorf("result should equal to original configuration")
			}
		}
	} else {
		t.Errorf("failed to marshal cfg")
	}
}

func TestDefaultValue(t *testing.T) {
	cfg := generateDefaultConfig()
	cfg.APPName = ""
	cfg.DB.Port = 0
	cfg.DB.SSL = false

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}

			var result testConfig
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			if !reflect.DeepEqual(result, generateDefaultConfig()) {
				t.Errorf("result should be set default value correctly")
			}
		}
	} else {
		t.Errorf("failed to marshal cfg")
	}
}

func TestMissingRequiredValue(t *testing.T) {
	cfg := generateDefaultConfig()
	cfg.DB.Password = ""

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}

			var result testConfig
			if _, err := Load(&result, file.Name()); err == nil {
				t.Errorf("Should got error when load configuration missing db password")
			}
		}
	} else {
		t.Errorf("failed to marshal cfg")
	}
}

func TestUnmatchedKeyInTomltestConfigFile(t *testing.T) {
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

	if err := toml.NewEncoder(file).Encode(cfg); err == nil {

		var result cfgStruct

		// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
		if err := New(&Settings{}).Load(&result, filename); err != nil {
			t.Errorf("Should NOT get error when loading configuration with extra keys")
		}

		// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
		err := New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename)
		if err == nil {
			t.Errorf("Should get error when loading configuration with extra keys")
		}

		// The error should be of type UnmatchedTomlKeysError
		tomlErr, ok := err.(*UnmatchedTomlKeysError)
		if !ok {
			t.Errorf("Should get UnmatchedTomlKeysError error when loading configuration with extra keys")
		}

		// The error.Keys() function should return the "Test" key
		keys := GetStringTomlKeys(tomlErr.Keys)
		if len(keys) != 1 || keys[0] != "Test" {
			t.Errorf("The UnmatchedTomlKeysError should contain the Test key")
		}

	} else {
		t.Errorf("failed to marshal cfg")
	}

	// Add .toml to the file name and test again
	err = os.Rename(filename, filename+".toml")
	if err != nil {
		t.Errorf("Could not add suffix to file")
	}
	filename = filename + ".toml"
	defer os.Remove(filename)

	var result cfgStruct

	// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
	if err := New(&Settings{}).Load(&result, filename); err != nil {
		t.Errorf("Should NOT get error when loading configuration with extra keys. Error: %v", err)
	}

	// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
	err = New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename)
	if err == nil {
		t.Errorf("Should get error when loading configuration with extra keys")
	}

	// The error should be of type UnmatchedTomlKeysError
	tomlErr, ok := err.(*UnmatchedTomlKeysError)
	if !ok {
		t.Errorf("Should get UnmatchedTomlKeysError error when loading configuration with extra keys")
	}

	// The error.Keys() function should return the "Test" key
	keys := GetStringTomlKeys(tomlErr.Keys)
	if len(keys) != 1 || keys[0] != "Test" {
		t.Errorf("The UnmatchedTomlKeysError should contain the Test key")
	}

}

func TestUnmatchedKeyInYamltestConfigFile(t *testing.T) {
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

	if data, err := yaml.Marshal(cfg); err == nil {
		if _, err := file.WriteString(string(data)); err != nil {
			t.Error(err)
		}

		var result cfgStruct

		// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
		if err := New(&Settings{}).Load(&result, filename); err != nil {
			t.Errorf("Should NOT get error when loading configuration with extra keys. Error: %v", err)
		}

		// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
		if err := New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename); err == nil {
			t.Errorf("Should get error when loading configuration with extra keys")

			// The error should be of type *yaml.TypeError
		} else if _, ok := err.(*yaml.TypeError); !ok {
			// || !strings.Contains(err.Error(), "not found in struct") {
			t.Errorf("Error should be of type yaml.TypeError. Instead error is %v", err)
		}

	} else {
		t.Errorf("failed to marshal cfg")
	}

	// Add .yaml to the file name and test again
	err = os.Rename(filename, filename+".yaml")
	if err != nil {
		t.Errorf("Could not add suffix to file")
	}
	filename = filename + ".yaml"
	defer os.Remove(filename)

	var result cfgStruct

	// Do not return error when there are unmatched keys but ErrorOnUnmatchedKeys is false
	if err := New(&Settings{}).Load(&result, filename); err != nil {
		t.Errorf("Should NOT get error when loading configuration with extra keys. Error: %v", err)
	}

	// Return an error when there are unmatched keys and ErrorOnUnmatchedKeys is true
	if err := New(&Settings{ErrorOnUnmatchedKeys: true}).Load(&result, filename); err == nil {
		t.Errorf("Should get error when loading configuration with extra keys")

		// The error should be of type *yaml.TypeError
	} else if _, ok := err.(*yaml.TypeError); !ok {
		// || !strings.Contains(err.Error(), "not found in struct") {
		t.Errorf("Error should be of type yaml.TypeError. Instead error is %v", err)
	}
}

func TestLoadtestConfigurationByEnvironment(t *testing.T) {
	cfg := generateDefaultConfig()
	cfg2 := struct {
		APPName string
	}{
		APPName: "cfg2",
	}

	if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
		defer file.Close()
		defer os.Remove(file.Name())
		cfgBytes, _ := yaml.Marshal(cfg)
		cfg2Bytes, _ := yaml.Marshal(cfg2)
		if err := ioutil.WriteFile(file.Name()+".yaml", cfgBytes, 0644); err != nil {
			t.Error(err)
		}
		defer os.Remove(file.Name() + ".yaml")
		if err := ioutil.WriteFile(file.Name()+".production.yaml", cfg2Bytes, 0644); err != nil {
			t.Error(err)
		}
		defer os.Remove(file.Name() + ".production.yaml")

		var result testConfig
		os.Setenv("CONFIG_ENV", "production")
		defer os.Setenv("CONFIG_ENV", "")
		if _, err := Load(&result, file.Name()+".yaml"); err != nil {
			t.Errorf("No error should happen when load configurations, but got %v", err)
		}

		var defaultConfig = generateDefaultConfig()
		defaultConfig.APPName = "cfg2"
		if !reflect.DeepEqual(result, defaultConfig) {
			t.Errorf("result should be load configurations by environment correctly")
		}
	}
}

func TestLoadtestConfigurationByEnvironmentSetBytestConfig(t *testing.T) {
	cfg := generateDefaultConfig()
	cfg2 := struct {
		APPName string
	}{
		APPName: "production_cfg2",
	}

	if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
		defer file.Close()
		defer os.Remove(file.Name())
		cfgBytes, _ := yaml.Marshal(cfg)
		cfg2Bytes, _ := yaml.Marshal(cfg2)
		if err := ioutil.WriteFile(file.Name()+".yaml", cfgBytes, 0644); err != nil {
			t.Error(err)
		}
		defer os.Remove(file.Name() + ".yaml")
		if err := ioutil.WriteFile(file.Name()+".production.yaml", cfg2Bytes, 0644); err != nil {
			t.Error(err)
		}
		defer os.Remove(file.Name() + ".production.yaml")

		var result testConfig
		var cfg = New(&Settings{Environment: "production"})
		if err := cfg.Load(&result, file.Name()+".yaml"); err != nil {
			t.Errorf("No error should happen when load configurations, but got %v", err)
		}

		var defaultConfig = generateDefaultConfig()
		defaultConfig.APPName = "production_cfg2"
		if !reflect.DeepEqual(result, defaultConfig) {
			t.Errorf("result should be load configurations by environment correctly")
		}

		if cfg.GetEnvironment() != "production" {
			t.Errorf("config's environment should be production")
		}
	}
}

func TestOverwritetestConfigurationWithEnvironmentWithDefaultPrefix(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}
			var result testConfig
			os.Setenv("CONFIG_APPNAME", "cfg2")
			os.Setenv("CONFIG_HOSTS", "- http://example.org\n- http://jinzhu.me")
			os.Setenv("CONFIG_DB_NAME", "db_name")
			defer os.Setenv("CONFIG_APPNAME", "")
			defer os.Setenv("CONFIG_HOSTS", "")
			defer os.Setenv("CONFIG_DB_NAME", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.APPName = "cfg2"
			defaultConfig.Hosts = []string{"http://example.org", "http://jinzhu.me"}
			defaultConfig.DB.Name = "db_name"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestOverwritetestConfigurationWithEnvironment(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}

			var result testConfig
			os.Setenv("CONFIG_ENV_PREFIX", "app")
			os.Setenv("APP_APPNAME", "cfg2")
			os.Setenv("APP_DB_NAME", "db_name")
			defer os.Setenv("CONFIG_ENV_PREFIX", "")
			defer os.Setenv("APP_APPNAME", "")
			defer os.Setenv("APP_DB_NAME", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.APPName = "cfg2"
			defaultConfig.DB.Name = "db_name"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestOverwritetestConfigurationWithEnvironmentThatSetBytestConfig(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}

			os.Setenv("APP1_APPName", "cfg2")
			os.Setenv("APP1_DB_Name", "db_name")
			defer os.Setenv("APP1_APPName", "")
			defer os.Setenv("APP1_DB_Name", "")

			var result testConfig
			var cfg = New(&Settings{ENVPrefix: "APP1"})
			if err := cfg.Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.APPName = "cfg2"
			defaultConfig.DB.Name = "db_name"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestResetPrefixToBlank(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}
			var result testConfig
			os.Setenv("CONFIG_ENV_PREFIX", "-")
			os.Setenv("APPNAME", "cfg2")
			os.Setenv("DB_NAME", "db_name")
			defer os.Setenv("CONFIG_ENV_PREFIX", "")
			defer os.Setenv("APPNAME", "")
			defer os.Setenv("DB_NAME", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.APPName = "cfg2"
			defaultConfig.DB.Name = "db_name"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestResetPrefixToBlank2(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}
			var result testConfig
			os.Setenv("CONFIG_ENV_PREFIX", "-")
			os.Setenv("APPName", "cfg2")
			os.Setenv("DB_Name", "db_name")
			defer os.Setenv("CONFIG_ENV_PREFIX", "")
			defer os.Setenv("APPName", "")
			defer os.Setenv("DB_Name", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.APPName = "cfg2"
			defaultConfig.DB.Name = "db_name"
			if !reflect.DeepEqual(result, defaultConfig) {

				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestReadFromEnvironmentWithSpecifiedEnvName(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}
			var result testConfig
			os.Setenv("DBPassword", "db_password")
			defer os.Setenv("DBPassword", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.DB.Password = "db_password"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestAnonymousStruct(t *testing.T) {
	cfg := generateDefaultConfig()

	if bytes, err := json.Marshal(cfg); err == nil {
		if file, err := ioutil.TempFile("/tmp", "config"); err == nil {
			defer file.Close()
			defer os.Remove(file.Name())
			if _, err := file.Write(bytes); err != nil {
				t.Error(err)
			}
			var result testConfig
			os.Setenv("CONFIG_DESCRIPTION", "environment description")
			defer os.Setenv("CONFIG_DESCRIPTION", "")
			if _, err := Load(&result, file.Name()); err != nil {
				t.Error(err)
			}

			var defaultConfig = generateDefaultConfig()
			defaultConfig.Anonymous.Description = "environment description"
			if !reflect.DeepEqual(result, defaultConfig) {
				t.Errorf("result should equal to original configuration")
			}
		}
	}
}

func TestENV(t *testing.T) {
	if ENV() != "test" {
		t.Errorf("Env should be test when running `go test`, instead env is %v", ENV())
	}

	os.Setenv("CONFIG_ENV", "production")
	defer os.Setenv("CONFIG_ENV", "")
	if ENV() != "production" {
		t.Errorf("Env should be production when set it with CONFIG_ENV")
	}
}

type slicetestConfig struct {
	Test1 int
	Test2 []struct {
		Test2Ele1 int
		Test2Ele2 int
	}
}

func TestSliceFromEnv(t *testing.T) {
	var tc = slicetestConfig{
		Test1: 1,
		Test2: []struct {
			Test2Ele1 int
			Test2Ele2 int
		}{
			{
				Test2Ele1: 1,
				Test2Ele2: 2,
			},
			{
				Test2Ele1: 3,
				Test2Ele2: 4,
			},
		},
	}

	var result slicetestConfig
	os.Setenv("CONFIG_TEST1", "1")
	os.Setenv("CONFIG_TEST2_0_TEST2ELE1", "1")
	os.Setenv("CONFIG_TEST2_0_TEST2ELE2", "2")

	os.Setenv("CONFIG_TEST2_1_TEST2ELE1", "3")
	os.Setenv("CONFIG_TEST2_1_TEST2ELE2", "4")
	_, err := Load(&result)
	if err != nil {
		t.Fatalf("load from env err:%v", err)
	}

	if !reflect.DeepEqual(result, tc) {
		t.Fatalf("unexpected result:%+v", result)
	}
}

func TestConfigFromEnv(t *testing.T) {
	type config struct {
		LineBreakString string `required:"true"`
		Count           int64
		Slient          bool
	}

	cfg := &config{}

	os.Setenv("CONFIG_ENV_PREFIX", "CONFIG")
	os.Setenv("CONFIG_LineBreakString", "Line one\nLine two\nLine three\nAnd more lines")
	os.Setenv("CONFIG_Slient", "1")
	os.Setenv("CONFIG_Count", "10")
	if _, err := Load(cfg); err != nil {
		t.Error(err)
	}

	if os.Getenv("CONFIG_LineBreakString") != cfg.LineBreakString {
		t.Error("Failed to load value has line break from env")
	}

	if !cfg.Slient {
		t.Error("Failed to load bool from env")
	}

	if cfg.Count != 10 {
		t.Error("Failed to load number from env")
	}
}

type Menu struct {
	Key      string `json:"key" yaml:"key"`
	Name     string `json:"name" yaml:"name"`
	Icon     string `json:"icon" yaml:"icon"`
	Children []Menu `json:"children" yaml:"children"`
}

type MenuList struct {
	Top []Menu `json:"top"  yaml:"top"`
}

func TestLoadNestedConfig(t *testing.T) {
	adminConfig := MenuList{}
	if err := New(&Settings{Verbose: true}).Load(&adminConfig, "./admin.yml"); err != nil {
		t.Error(err)
	}
}
