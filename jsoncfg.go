package jsoncfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gitlab.com/mjwhitta/pathname"
)

const Version = "1.1.3"

type jsoncfg struct {
	autosave      bool
	config        map[string]interface{}
	defaultConfig []byte
	diff          map[string]interface{}
	File          string
}

func (c *jsoncfg) Clear() {
	c.config = map[string]interface{}{}
	c.diff = map[string]interface{}{}
	c.write(false)
}

func (c *jsoncfg) Default() error {
	var e = json.Unmarshal(c.defaultConfig, &c.config)
	if e != nil {
		return e
	}

	e = json.Unmarshal(c.defaultConfig, &c.diff)
	if e != nil {
		return e
	}

	c.write(false)
	return nil
}

func (c *jsoncfg) Get(key string) interface{} {
	return c.config[key]
}

func (c *jsoncfg) GetDiff(key string) interface{} {
	return c.diff[key]
}

func (c *jsoncfg) Has(key string) bool {
	var _, hasKey = c.diff[key]
	return hasKey
}

// Constructor
func New(file string) jsoncfg {
	return jsoncfg{
		autosave:      false,
		config:        map[string]interface{}{},
		defaultConfig: []byte{},
		diff:          map[string]interface{}{},
		File:          pathname.ExpandPath(file),
	}
}

// Constructor
func NewAutosave(file string) jsoncfg {
	return jsoncfg{
		autosave:      true,
		config:        map[string]interface{}{},
		defaultConfig: []byte{},
		diff:          map[string]interface{}{},
		File:          pathname.ExpandPath(file),
	}
}

// FIXME remove
func (c *jsoncfg) Print() error {
	var config, e = json.MarshalIndent(c.config, "", "  ")
	if e != nil {
		return e
	}

	fmt.Println(string(config))
	return nil
}

// FIXME remove
func (c *jsoncfg) PrintDefault() {
	fmt.Println(string(c.defaultConfig))
}

func (c *jsoncfg) read() error {
	// TODO check if symlink
	if !pathname.Exists(c.File) {
		c.Default()
		c.write(true)
	}

	var config, e = ioutil.ReadFile(c.File)
	if e != nil {
		return e
	}

	e = json.Unmarshal([]byte(config), &c.config)
	if e != nil {
		return e
	}

	e = json.Unmarshal(c.defaultConfig, &c.diff)
	if e != nil {
		return e
	}

	return nil
}

func (c *jsoncfg) Reset() error {
	return c.read()
}

func (c *jsoncfg) Save() error {
	var e = json.Unmarshal(c.defaultConfig, &c.diff)
	if e != nil {
		return e
	}

	return c.write(true)
}

func (c *jsoncfg) SaveDiff() error {
	var diff, e = json.Marshal(c.diff)
	if e != nil {
		return e
	}

	e = json.Unmarshal(diff, &c.config)
	if e != nil {
		return e
	}

	return c.write(true)
}

func (c *jsoncfg) SaveDefault() error {
	var config, e = json.Marshal(c.config)
	if e != nil {
		return e
	}

	c.defaultConfig = config
	return nil
}

func (c *jsoncfg) Set(key string, value interface{}) error {
	c.config[key] = value
	c.diff[key] = value
	return c.write(false)
}

func (c *jsoncfg) SetDefault(key string, value interface{}) {
	c.config[key] = value
	c.diff[key] = value
}

func (c *jsoncfg) write(force bool) error {
	if !c.autosave && !force {
		return nil
	}

	// TODO make parent directory first
	var config, e = json.MarshalIndent(c.config, "", "  ")
	if e != nil {
		return e
	}

	return ioutil.WriteFile(c.File, config, 0600)
}
