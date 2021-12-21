package jsoncfg

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gitlab.com/mjwhitta/errors"
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

// New will return a pointer to a new JSONCfg instance that requires
// manual calls to Save() to write the config to disk.
func New(file ...string) *JSONCfg {
	var config *jq.JSON
	var diff *jq.JSON

	config, _ = jq.New()
	diff, _ = jq.New()

	return &JSONCfg{
		autosave: false,
		config:   config,
		diff:     diff,
		File:     pathname.ExpandPath(filepath.Join(file...)),
	}
}

// NewAutosave will return a pointer to a new JSONCfg instance that is
// immediately written to disk on change.
func NewAutosave(file ...string) *JSONCfg {
	var c *JSONCfg = New(file...)

	c.autosave = true

	return c
}

// Append will append the specified value to the specified key in the
// config, if it is an array.
func (c *JSONCfg) Append(
	value interface{},
	keys ...interface{},
) error {
	var e error

	if e = c.AppendDefault(value, keys...); e != nil {
		return e
	}

	return c.write(false)
}

// AppendDefault will append the specified value to the specified key
// in the config, if it is an array. It will not write changes to disk
// ever and is intended to be used prior to SaveDefault().
func (c *JSONCfg) AppendDefault(
	value interface{},
	keys ...interface{},
) error {
	var e error

	if e = c.config.Append(value, keys...); e != nil {
		return e
	}

	return c.diff.Append(value, keys...)
}

// Clear will erase the config.
func (c *JSONCfg) Clear() {
	c.config.SetBlob()
	c.diff.SetBlob()
	c.write(false)
}

// Default will return the config to a pre-configured default state.
func (c *JSONCfg) Default() error {
	var e error

	if e = c.config.SetBlob(c.defaultConfig); e != nil {
		return errors.Newf("failed to reset config: %w", e)
	}

	if e = c.diff.SetBlob(c.defaultConfig); e != nil {
		return errors.Newf("failed to reset diff: %w", e)
	}

	return c.write(false)
}

// GetKeys will return a list of valid keys if the specified key
// returns an arry or map.
func (c *JSONCfg) GetKeys(keys ...interface{}) []string {
	return c.config.GetKeys(keys...)
}

// HasKey will return true if the config has the specified key, false
// otherwise.
func (c *JSONCfg) HasKey(keys ...interface{}) bool {
	return c.config.HasKey(keys...)
}

// MustGetKeys will return a list of valid keys if the specified key
// returns an arry or map.
func (c *JSONCfg) MustGetKeys(keys ...interface{}) ([]string, error) {
	return c.config.MustGetKeys(keys...)
}

// Reset will read the config from disk, erasing any unsaved changes.
func (c *JSONCfg) Reset() error {
	var config []byte
	var e error

	if c.File == "" {
		return nil
	}

	if !pathname.DoesExist(c.File) {
		c.Default()
		c.write(true)
	}

	if config, e = ioutil.ReadFile(c.File); e != nil {
        return errors.Newf("failed to read config %s: %w", c.File, e)
	}

	if e = c.config.SetBlob(string(config)); e != nil {
        return errors.Newf("failed to parse config %s: %w", c.File, e)
	}

	if c.defaultConfig == "" {
		c.defaultConfig = string(config)
	}

	return c.diff.SetBlob(c.defaultConfig)
}

// Save will save any unsaved changes to disk.
func (c *JSONCfg) Save() error {
	var e error

	if e = c.diff.SetBlob(c.defaultConfig); e != nil {
		return errors.Newf("failed to save config: %w", e)
	}

	return c.write(true)
}

// SaveDiff will save only the changes from default to disk.
func (c *JSONCfg) SaveDiff() error {
	var diff string
	var e error

	if diff, e = c.diff.GetBlob(); e != nil {
		return errors.Newf("failed to save diff: %w", e)
	}

	if e = c.config.SetBlob(diff); e != nil {
		return errors.Newf("failed to parse diff: %w", e)
	}

	return c.write(true)
}

// SaveDefault will save the default map for use by Default().
func (c *JSONCfg) SaveDefault() error {
	var config string
	var e error

	if config, e = c.config.GetBlob(); e != nil {
		return errors.Newf("failed to get default config: %w", e)
	}

	c.defaultConfig = config
	return nil
}

// Set will set the specified value for the specified key in the
// config.
func (c *JSONCfg) Set(value interface{}, keys ...interface{}) error {
	var e error

	if e = c.SetDefault(value, keys...); e != nil {
		return e
	}

	return c.write(false)
}

// SetDefault will set the specified value for the specified key in
// the config. It will not write changes to disk ever and is intended
// to be used prior to SaveDefault().
func (c *JSONCfg) SetDefault(
	value interface{},
	keys ...interface{},
) error {
	var e error

	if e = c.config.Set(value, keys...); e != nil {
		return errors.Newf("failed to set key %v: %w", keys, e)
	}

	return c.diff.Set(value, keys...)
}

// String will return a string representation of a config.
func (c *JSONCfg) String() string {
	return c.config.String()
}

func (c *JSONCfg) write(force bool) error {
	if (c.File == "") || (!c.autosave && !force) {
		return nil
	}

	var config string
	var e error

	e = os.MkdirAll(pathname.Dirname(c.File), os.ModePerm)
	if e != nil {
		e = errors.Newf(
			"failed to create directory tree for %s: %w",
			c.File,
			e,
		)
		return e
	}

	if config, e = c.config.GetBlob("  "); e != nil {
		return errors.Newf("failed to read config: %w", e)
	}

	if e = ioutil.WriteFile(c.File, []byte(config), 0600); e != nil {
		e = errors.Newf("failed to write config to %s: %w", c.File, e)
		return e
	}

	return nil
}
