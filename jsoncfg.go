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
	inMemory      bool
}

// New will return a pointer to a new JSONCfg instance that requires
// manual calls to Save() to write the config to disk.
func New(file string) *JSONCfg {
	var config *jq.JSON
	var diff *jq.JSON

	config, _ = jq.New("{}")
	diff, _ = jq.New("{}")

	return &JSONCfg{
		autosave:      false,
		config:        config,
		defaultConfig: "{}",
		diff:          diff,
		File:          pathname.ExpandPath(file),
		inMemory:      false,
	}
}

// NewAutosave will return a pointer to a new JSONCfg instance that is
// immediately written to disk on change.
func NewAutosave(file string) *JSONCfg {
	var c *JSONCfg

	c = New(file)
	c.autosave = true

	return c
}

// NewInMemory will return a pointer to a new JSONCfg instance that
// exists in memory and is never written to disk. This also means the
// config can not be read back from disk.
func NewInMemory() *JSONCfg {
	var config *jq.JSON
	var diff *jq.JSON

	config, _ = jq.New("{}")
	diff, _ = jq.New("{}")

	return &JSONCfg{
		autosave:      false,
		config:        config,
		defaultConfig: "{}",
		diff:          diff,
		inMemory:      true,
	}
}

// Clear will erase the config.
func (c *JSONCfg) Clear() {
	c.config.SetBlob("{}")
	c.diff.SetBlob("{}")
	c.write(false)
}

// Default will return the config to a pre-configured default state.
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

// GetKeys will return a list of valid keys if the specified key
// returns an arry or map.
func (c *JSONCfg) GetKeys(key ...interface{}) []string {
	return c.config.GetKeys(key...)
}

// HasKey will return true if the config has the specified key, false
// otherwise.
func (c *JSONCfg) HasKey(key ...interface{}) bool {
	return c.config.HasKey(key...)
}

// MustGetKeys will return a list of valid keys if the specified key
// returns an arry or map.
func (c *JSONCfg) MustGetKeys(key ...interface{}) ([]string, error) {
	return c.config.MustGetKeys(key...)
}

// Reset will read the config from disk, erasing any unsaved changes.
func (c *JSONCfg) Reset() error {
	var config []byte
	var e error

	if c.inMemory {
		return nil
	}

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
// config.
func (c *JSONCfg) Set(value interface{}, keys ...interface{}) error {
	c.config.Set(value, keys...)
	c.diff.Set(value, keys...)
	return c.write(false)
}

// SetDefault will set the specified value for the specified key in
// the config. It will not write changes to disk ever and is intended
// to be used prior to SaveDefault().
func (c *JSONCfg) SetDefault(value interface{}, keys ...interface{}) {
	c.config.Set(value, keys...)
	c.diff.Set(value, keys...)
}

// String will return a string representation of a config.
func (c *JSONCfg) String() string {
	return c.config.String()
}

func (c *JSONCfg) write(force bool) error {
	if c.inMemory || (!c.autosave && !force) {
		return nil
	}

	var config string
	var e error

	e = os.MkdirAll(pathname.Dirname(c.File), os.ModePerm)
	if e != nil {
		return e
	}

	if config, e = c.config.GetBlob("  "); e != nil {
		return e
	}

	return ioutil.WriteFile(c.File, []byte(config), 0600)
}
