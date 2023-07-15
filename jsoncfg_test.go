package jsoncfg_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/mjwhitta/jsoncfg"
	assert "github.com/stretchr/testify/require"
)

var def = map[string]any{
	"a": true,
	"b": "asdf",
	"c": 1234,
	"d": [2]string{"blah", "test"},
	"e": map[string]any{
		"aFloat": 1.2,
		"anInt":  17,
		"more": map[string]any{
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

var testcfg string = "/tmp/jsoncfg_test"

func TestAppend(t *testing.T) {
	var actual string
	var cfg *jsoncfg.JSONCfg
	var e error
	var expected string

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	e = cfg.Append("asdf", "d")
	assert.Nil(t, e)

	expected = "[blah test asdf]"
	actual = fmt.Sprintf("%v", cfg.GetArray("d"))
	assert.Equal(t, expected, actual)

	e = cfg.Append(2, "d")
	assert.Nil(t, e)

	expected = "[blah test asdf 2]"
	actual = fmt.Sprintf("%v", cfg.GetArray("d"))
	assert.Equal(t, expected, actual)
}

func TestClear(t *testing.T) {
	var cfg *jsoncfg.JSONCfg
	var expected string = "{}"

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()
	cfg.Clear()

	assert.Equal(t, expected, cfg.String())
}

func TestDefault(t *testing.T) {
	var cfg *jsoncfg.JSONCfg
	var e error

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	_, e = cfg.MustGetDiffArray("d")
	assert.Nil(t, e)

	assert.Equal(t, json, cfg.String())

	cfg.Set(2, "e", "anInt")
	assert.NotEqual(t, json, cfg.String())

	cfg.Default()
	assert.Equal(t, json, cfg.String())
}

func TestHasKey(t *testing.T) {
	var cfg *jsoncfg.JSONCfg = jsoncfg.New(testcfg)

	cfg.Reset()

	assert.True(t, cfg.HasKey("a"))
	assert.False(t, cfg.HasKey("asdf"))
}

func TestKeys(t *testing.T) {
	var actual []string
	var cfg *jsoncfg.JSONCfg
	var e error
	var expected string

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	expected = fmt.Sprintf("%v", []string{"0", "1"})
	actual, e = cfg.MustGetKeys("d")
	assert.Nil(t, e)
	assert.Equal(t, expected, fmt.Sprintf("%v", actual))

	expected = fmt.Sprintf("%v", []string{"aFloat", "anInt", "more"})
	actual, e = cfg.MustGetKeys("e")
	assert.Nil(t, e)
	assert.Equal(t, expected, fmt.Sprintf("%v", actual))

	actual = cfg.GetKeys("a")
	assert.Equal(t, 0, len(actual))

	_, e = cfg.MustGetKeys("a")
	assert.NotNil(t, e)
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
	var newMap map[string]any

	cfg = jsoncfg.New(testcfg)
	cfg.Reset()

	assert.Equal(t, json, cfg.String())

	e = cfg.Set("asdf", "d", 0)
	assert.Nil(t, e)

	actual, e = cfg.MustGetString("d", 0)
	assert.Nil(t, e)
	assert.Equal(t, "asdf", actual)

	e = cfg.Set("asdf", "d", "asdf")
	assert.NotNil(t, e)

	e = cfg.Set("asdf", "e", 0)
	assert.NotNil(t, e)

	e = cfg.Set("asdf", "e", "asdf", "blah")
	assert.NotNil(t, e)

	newMap = map[string]any{"asdf": "blah", "anInt": 7}

	e = cfg.Set(newMap)
	assert.Nil(t, e)

	actual = fmt.Sprintf("%+v", cfg.GetMap())
	expected = fmt.Sprintf("%+v", newMap)
	assert.Equal(t, expected, actual)

	cfg.Reset()
	assert.Equal(t, json, cfg.String())
}
