package jsoncfg

import (
	"io/ioutil"
	"os"

	"gitlab.com/mjwhitta/jq"
	"gitlab.com/mjwhitta/pathname"
)

// JSONCfg is a struct that handles a JSON formatted config file on
// disk. It contains the filename, the running config, the default
// config, and any changes from default. If autosave is true, changes
// are written to disk immediately.
type JSONCfg struct {
	autosave      bool
	config        *jq.JSON
	defaultConfig string
	diff          *jq.JSON
	File          string
}

// New is a JSONCfg constructor where autosave is false.
func New(file string) *JSONCfg {
	var config *jq.JSON
	var diff *jq.JSON

	config, _ = jq.New("{}")
	diff, _ = jq.New("{}")

	return &JSONCfg{
		autosave:      false,
		config:        config,
		defaultConfig: "",
		diff:          diff,
		File:          pathname.ExpandPath(file),
	}
}

// NewAutosave is a JSONCfg constructor where autosave is true.
func NewAutosave(file string) *JSONCfg {
	var c *JSONCfg

	c = New(file)
	c.autosave = true

	return c
}

// Clear will erase the config struct.
func (c *JSONCfg) Clear() {
	c.config.SetBlob("{}")
	c.diff.SetBlob("{}")
	c.write(false)
}

// Default will return the config struct to a pre-configured default
// state.
func (c *JSONCfg) Default() error {
	var e error

	if e = c.config.SetBlob(c.defaultConfig); e != nil {
		return e
	}

	if e = c.diff.SetBlob(c.defaultConfig); e != nil {
		return e
	}

	return c.write(false)
}

// Has will return true if the config struct has the specified key,
// false otherwise.
func (c *JSONCfg) Has(key string) bool {
	return c.diff.Has(key)
}

func (c *JSONCfg) read() error {
	var config []byte
	var e error

	if !pathname.DoesExist(c.File) {
		c.Default()
		c.write(true)
	}

	if config, e = ioutil.ReadFile(c.File); e != nil {
		return e
	}

	if e = c.config.SetBlob(string(config)); e != nil {
		return e
	}

	return c.diff.SetBlob(c.defaultConfig)
}

// Reset will read the config from disk, erasing any unsaved changes.
func (c *JSONCfg) Reset() error {
	return c.read()
}

// Save will save any unsaved changes to disk.
func (c *JSONCfg) Save() error {
	var e error

	if e = c.diff.SetBlob(c.defaultConfig); e != nil {
		return e
	}

	return c.write(true)
}

// SaveDiff will save only the changes from default to disk.
func (c *JSONCfg) SaveDiff() error {
	var diff string
	var e error

	if diff, e = c.diff.GetBlob(); e != nil {
		return e
	}

	if e = c.config.SetBlob(diff); e != nil {
		return e
	}

	return c.write(true)
}

// SaveDefault will save the default map for use by Default().
func (c *JSONCfg) SaveDefault() error {
	var config string
	var e error

	if config, e = c.config.GetBlob(); e != nil {
		return e
	}

	c.defaultConfig = config
	return nil
}

// Set will set the specified value for the specified key in the
// config struct.
func (c *JSONCfg) Set(key string, value interface{}) error {
	c.config.Set(key, value)
	c.diff.Set(key, value)
	return c.write(false)
}

// SetDefault will set the specified value for the specified key in
// the config struct. It will not write changes to disk ever and is
// intended to be used prior to SaveDefault().
func (c *JSONCfg) SetDefault(key string, value interface{}) {
	c.config.Set(key, value)
	c.diff.Set(key, value)
}

func (c *JSONCfg) String() string {
	var toString string
	toString, _ = c.config.GetBlobIndent("", "  ")
	return toString
}

func (c *JSONCfg) write(force bool) error {
	if !c.autosave && !force {
		return nil
	}

	var config string
	var e error

	e = os.MkdirAll(pathname.Dirname(c.File), os.ModePerm)
	if e != nil {
		return e
	}

	if config, e = c.config.GetBlobIndent("", "  "); e != nil {
		return e
	}

	return ioutil.WriteFile(c.File, []byte(config), 0600)
}
