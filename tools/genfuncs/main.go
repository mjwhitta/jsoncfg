package main

import (
	"os"
	"strings"

	"github.com/mjwhitta/errors"
)

func generateFuncs(f *os.File, t string) {
	var rt string = t
	var T string
	var tmp string
	var txt string = `
// Get{A}{B} will return the value for the specified key(s) as a {C}.
func (c *JSONCfg) Get{A}{B}(key ...any) {C} {
    return c.{D}.Get{B}(key...)
}

// Get{A}{B}Array will return an array for the specified key(s) as a
// []{C}.
func (c *JSONCfg) Get{A}{B}Array(key ...any) []{C} {
    return c.{D}.Get{B}Array(key...)
}

// Get{A}{B}Map will return a map for the specified key(s) as a
// map[string]{C}.
func (c *JSONCfg) Get{A}{B}Map(key ...any) map[string]{C} {
    return c.{D}.Get{B}Map(key...)
}

// MustGet{A}{B} will return the value for the specified key(s) as a
// {C}.
func (c *JSONCfg) MustGet{A}{B}(key ...any) ({C}, error) {
    return c.{D}.MustGet{B}(key...)
}

// MustGet{A}{B}Array will return an array for the specified key(s) as
// a []{C}.
func (c *JSONCfg) MustGet{A}{B}Array(key ...any) ([]{C}, error) {
    return c.{D}.MustGet{B}Array(key...)
}

// MustGet{A}{B}Map will return a map for the specified key(s) as a
// map[string]{C}.
func (c *JSONCfg) MustGet{A}{B}Map(
    key ...any,
) (map[string]{C}, error) {
    return c.{D}.MustGet{B}Map(key...)
}
`

	if t != "" {
		T = strings.ToUpper(t[0:1]) + t[1:]
	} else {
		rt = "any"
	}

	tmp = strings.ReplaceAll(txt, "{A}", "")
	tmp = strings.ReplaceAll(tmp, "{B}", T)
	tmp = strings.ReplaceAll(tmp, "{C}", rt)
	tmp = strings.ReplaceAll(tmp, "{D}", "config")
	f.WriteString(tmp)

	tmp = strings.ReplaceAll(txt, "{A}", "Diff")
	tmp = strings.ReplaceAll(tmp, "{B}", T)
	tmp = strings.ReplaceAll(tmp, "{C}", rt)
	tmp = strings.ReplaceAll(tmp, "{D}", "diff")
	f.WriteString(tmp)
}

func header(f *os.File) {
	f.WriteString("// Code generated by genfuncs; DO NOT EDIT.\n")
	f.WriteString("package jsoncfg\n")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			panic(r.(error).Error())
		}
	}()

	var e error
	var f *os.File
	var fn string = "generated.go"
	var types []string = []string{
		"",
		"bool",
		"float32", "float64",
		"int", "int16", "int32", "int64",
		"string",
		"uint", "uint16", "uint32", "uint64",
	}

	if f, e = os.Create(fn); e != nil {
		panic(errors.Newf("failed to create %s: %w", fn, e))
	}
	defer f.Close()

	header(f)

	for _, thetype := range types {
		generateFuncs(f, thetype)
	}
}
