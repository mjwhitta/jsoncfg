package jsoncfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/mjwhitta/pathname"
)

const Version = "1.1.3"

type jsoncfg struct {
	defaultConfig []byte
	File          string
	Prefs         map[string]interface{}
}

// Constructor
func New(file string) jsoncfg {
	return jsoncfg{
		defaultConfig: []byte{},
		File:          pathname.ExpandPath(file),
		Prefs:         map[string]interface{}{},
	}
}

func (config *jsoncfg) Print() error {
	var content, e = json.MarshalIndent(config.Prefs, "", "  ")
	if e != nil {
		return e
	}

	fmt.Println(string(content))
	return nil
}

func (config *jsoncfg) PrintDefault() {
	fmt.Println(string(config.defaultConfig))
}

func (config *jsoncfg) Read() error {
	var content, e = ioutil.ReadFile(config.File)
	if e != nil {
		return e
	}

	e = json.Unmarshal([]byte(content), &config.Prefs)
	if e != nil {
		return e
	}

	return nil
}

func (config *jsoncfg) Reset() error {
	return json.Unmarshal([]byte(config.defaultConfig), &config.Prefs)
}

func (config *jsoncfg) SaveDefault() error {
	var content, e = json.Marshal(config.Prefs)
	if e != nil {
		return e
	}
	config.defaultConfig = content
	return nil
}

func (config *jsoncfg) Write() error {
	var content, e = json.MarshalIndent(config.Prefs, "", "  ")
	if e != nil {
		return e
	}

	return ioutil.WriteFile(config.File, content, 0600)
}
