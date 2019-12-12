package jsoncfg

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gitlab.com/mjwhitta/pathname"
)

// JsonCfg is a struct that handles a JSON formatted config file on
// disk. It contains the filename, the running config, the default
// config, and any changes from default. If autosave is true, changes
// are written to disk immediately.
type JsonCfg struct {
	autosave      bool
	config        map[string]interface{}
	defaultConfig []byte
	diff          map[string]interface{}
	File          string
}

// New is a JsonCfg constructor where autosave is false.
func New(file string) *JsonCfg {
	return &JsonCfg{
		autosave:      false,
		config:        map[string]interface{}{},
		defaultConfig: []byte{},
		diff:          map[string]interface{}{},
		File:          pathname.ExpandPath(file),
	}
}

// NewAutosave is a JsonCfg constructor where autosave is true.
func NewAutosave(file string) *JsonCfg {
	return &JsonCfg{
		autosave:      true,
		config:        map[string]interface{}{},
		defaultConfig: []byte{},
		diff:          map[string]interface{}{},
		File:          pathname.ExpandPath(file),
	}
}

// Clear will erase the config struct.
func (c *JsonCfg) Clear() {
	c.config = map[string]interface{}{}
	c.diff = map[string]interface{}{}
	c.write(false)
}

// Default will return the config struct to a pre-configured default
// state.
func (c *JsonCfg) Default() error {
	var e error

	e = json.Unmarshal(c.defaultConfig, &c.config)
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

// Get will return a key from the config struct.
func (c *JsonCfg) Get(key string) interface{} {
	return c.config[key]
}

// GetDiff will return a key from the diff map in the config struct.
func (c *JsonCfg) GetDiff(key string) interface{} {
	return c.diff[key]
}

// Has will return true if the config struct has the specified key,
// false otherwise.
func (c *JsonCfg) Has(key string) bool {
	var hasKey bool
	_, hasKey = c.diff[key]
	return hasKey
}

func (c *JsonCfg) read() error {
	if !pathname.DoesExist(c.File) {
		c.Default()
		c.write(true)
	}

	var config []byte
	var e error

	config, e = ioutil.ReadFile(c.File)
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

// Reset will read the config from disk, erasing any unsaved changes.
func (c *JsonCfg) Reset() error {
	return c.read()
}

// Save will save any unsaved changes to disk.
func (c *JsonCfg) Save() error {
	var e error

	e = json.Unmarshal(c.defaultConfig, &c.diff)
	if e != nil {
		return e
	}

	return c.write(true)
}

// SaveDiff will save only the changes from default to disk.
func (c *JsonCfg) SaveDiff() error {
	var diff []byte
	var e error

	diff, e = json.Marshal(c.diff)
	if e != nil {
		return e
	}

	e = json.Unmarshal(diff, &c.config)
	if e != nil {
		return e
	}

	return c.write(true)
}

// SaveDefault will save the default map for use by Default().
func (c *JsonCfg) SaveDefault() error {
	var config []byte
	var e error

	config, e = json.Marshal(c.config)
	if e != nil {
		return e
	}

	c.defaultConfig = config
	return nil
}

// Set will set the specified value for the specified key in the
// config struct.
func (c *JsonCfg) Set(key string, value interface{}) error {
	c.config[key] = value
	c.diff[key] = value
	return c.write(false)
}

// SetDefault will set the specified value for the specified key in
// the config struct. It will not write changes to disk ever and is
// intended to be used prior to SaveDefault().
func (c *JsonCfg) SetDefault(key string, value interface{}) {
	c.config[key] = value
	c.diff[key] = value
}

func (c *JsonCfg) write(force bool) error {
	if !c.autosave && !force {
		return nil
	}

	var config []byte
	var e error

	e = os.MkdirAll(pathname.Dirname(c.File), os.ModePerm)
	if e != nil {
		return e
	}

	config, e = json.MarshalIndent(c.config, "", "  ")
	if e != nil {
		return e
	}

	return ioutil.WriteFile(c.File, config, 0600)
}
