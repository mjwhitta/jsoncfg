package jsoncfg_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"gitlab.com/mjwhitta/jsoncfg"
)

var def = map[string]interface{}{
	"a": true,
	"b": "asdf",
	"c": 1234,
	"d": [2]string{"blah", "test"},
	"e": map[string]interface{}{
		"aFloat": 1.2,
		"anInt":  17,
		"more": map[string]interface{}{
			"aFloat32": 1.2,
			"anInt64":  19,
		},
	},
}

var json string = strings.Join(
	[]string{
		"{",
		"  \"a\": true,",
		"  \"b\": \"asdf\",",
		"  \"c\": 1234,",
		"  \"d\": [",
		"    \"blah\",",
		"    \"test\"",
		"  ],",
		"  \"e\": {",
		"    \"aFloat\": 1.2,",
		"    \"anInt\": 17,",
		"    \"more\": {",
		"      \"aFloat32\": 1.2,",
		"      \"anInt64\": 19",
		"    }",
		"  }",
		"}",
	},
	"\n",
)

var testcfg = "/tmp/jsoncfg_test"

func TestAppend(t *testing.T) {
	var actual string
	var cfg *jsoncfg.JSONCfg
	var e error
	var expected string

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	if e = cfg.Append("asdf", "d"); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	}

	expected = "[blah test asdf]"
	actual = fmt.Sprintf("%v", cfg.GetArray("d"))
	if actual != expected {
		t.Errorf("\ngot: %s\nwant: %s", actual, expected)
	}

	if e = cfg.Append(2, "d"); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	}

	expected = "[blah test asdf 2]"
	actual = fmt.Sprintf("%v", cfg.GetArray("d"))
	if actual != expected {
		t.Errorf("\ngot: %s\nwant: %s", actual, expected)
	}
}

func TestClear(t *testing.T) {
	var cfg *jsoncfg.JSONCfg
	var expected string = "{}"

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()
	cfg.Clear()

	if cfg.String() != expected {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), expected)
	}
}

func TestDefault(t *testing.T) {
	var cfg *jsoncfg.JSONCfg
	var e error

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	if _, e = cfg.MustGetDiffArray("d"); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	}

	if cfg.String() != json {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), json)
	}

	cfg.Set(2, "e", "anInt")
	if cfg.String() == json {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), json)
	}

	cfg.Default()
	if cfg.String() != json {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), json)
	}
}

func TestHasKey(t *testing.T) {
	var cfg *jsoncfg.JSONCfg = jsoncfg.New(testcfg)

	cfg.Reset()

	if !cfg.HasKey("a") {
		t.Errorf("\ngot: false\nwant: true")
	}

	if cfg.HasKey("asdf") {
		t.Errorf("\ngot: true\nwant: false")
	}
}

func TestKeys(t *testing.T) {
	var actual []string
	var cfg *jsoncfg.JSONCfg
	var e error
	var expected string

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	expected = fmt.Sprintf("%v", []string{"0", "1"})
	if actual, e = cfg.MustGetKeys("d"); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	} else if fmt.Sprintf("%v", actual) != expected {
		t.Errorf("\ngot: %v\nwant: %v", actual, expected)
	}

	expected = fmt.Sprintf("%v", []string{"aFloat", "anInt", "more"})
	if actual, e = cfg.MustGetKeys("e"); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	} else if fmt.Sprintf("%v", actual) != expected {
		t.Errorf("\ngot: %v\nwant: %v", actual, expected)
	}

	if actual = cfg.GetKeys("a"); len(actual) > 0 {
		t.Errorf("\ngot: %v\nwant: []", actual)
	}

	expected = "jq: key [a] has no valid sub-keys"
	if _, e = cfg.MustGetKeys("a"); e == nil {
		t.Errorf("\ngot: nil\nwant: %s", expected)
	} else if e.Error() != expected {
		t.Errorf("\ngot: %s\nwant: %s", e.Error(), expected)
	}
}

func TestMain(m *testing.M) {
	var cfg *jsoncfg.JSONCfg
	var ret int

	cfg = jsoncfg.New(testcfg)
	cfg.SetDefault(def)
	cfg.SaveDefault()
	cfg.Reset()

	ret = m.Run()

	os.Remove(testcfg)
	os.Exit(ret)
}

func TestSet(t *testing.T) {
	var actual string
	var cfg *jsoncfg.JSONCfg
	var e error
	var expected string
	var newMap map[string]interface{}

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	if cfg.String() != json {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), json)
	}

	if e = cfg.Set("asdf", "d", 0); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	}

	expected = "asdf"
	if actual, e = cfg.MustGetString("d", 0); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	} else if actual != expected {
		t.Errorf("\ngot: %s\nwant: %s", actual, expected)
	}

	expected = strings.Join(
		[]string{
			"jsoncfg: failed to set key [d asdf]",
			"jq: key [d asdf] is not of type int",
		},
		": ",
	)
	if e = cfg.Set("asdf", "d", "asdf"); e == nil {
		t.Errorf("\ngot: nil\nwant: %s", expected)
	} else if e.Error() != expected {
		t.Errorf("\ngot: %s\nwant: %s", e.Error(), expected)
	}

	expected = strings.Join(
		[]string{
			"jsoncfg: failed to set key [e 0]",
			"jq: key [e 0] is not of type string",
		},
		": ",
	)
	if e = cfg.Set("asdf", "e", 0); e == nil {
		t.Errorf("\ngot: nil\nwant: %s", expected)
	} else if e.Error() != expected {
		t.Errorf("\ngot: %s\nwant: %s", e.Error(), expected)
	}

	expected = strings.Join(
		[]string{
			"jsoncfg: failed to set key [e asdf blah]",
			"jq: key [e asdf] not found",
		},
		": ",
	)
	if e = cfg.Set("asdf", "e", "asdf", "blah"); e == nil {
		t.Errorf("\ngot: nil\nwant: %s", expected)
	} else if e.Error() != expected {
		t.Errorf("\ngot: %s\nwant: %s", e.Error(), expected)
	}

	newMap = map[string]interface{}{"asdf": "blah", "anInt": 7}

	if e = cfg.Set(newMap); e != nil {
		t.Errorf("\ngot: %s\nwant: nil", e.Error())
	}

	actual = fmt.Sprintf("%+v", cfg.GetMap())
	expected = fmt.Sprintf("%+v", newMap)
	if actual != expected {
		t.Errorf("\ngot: %s\nwant: %s", actual, expected)
	}

	cfg.Reset()
	if cfg.String() != json {
		t.Errorf("\ngot: %s\nwant: %s", cfg.String(), json)
	}
}
