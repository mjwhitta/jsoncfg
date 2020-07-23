# jsoncfg

<a href="https://www.buymeacoffee.com/mjwhitta">🍪 Buy me a cookie</a>

[![Go Report Card](https://goreportcard.com/badge/gitlab.com/mjwhitta/jsoncfg)](https://goreportcard.com/report/gitlab.com/mjwhitta/jsoncfg)

## What is this?

This Go module allows you to read/write and get/set configuration
options from/to a JSON file.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/jsoncfg
```

## Usage

```
package main

import (
    "fmt"

    "gitlab.com/mjwhitta/jsoncfg"
)

var config *jsoncfg.JSONCfg

func init() {
    // Create a jsoncfg object
    config = jsoncfg.New("/tmp/rc")

    // Or if you want changes to be written to disk immediately:
    // config = jsoncfg.NewAutosave("/tmp/rc")

    config.SetDefault(true, "a")
    config.SetDefault("asdf", "b")
    config.SetDefault(1234, "c")
    config.SetDefault([]string{"blah", "test"}, "d")
    config.SetDefault(
        map[string]interface{}{"float": 1.2, "int": 0},
        "e",
    )
    config.SaveDefault()
    config.Reset()
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r.(error).Error())
        }
    }()

    var a bool
    var b string
    var c int64
    var d []string
    var e = map[string]interface{}{}
    var err error
    var f float64
    var keys []string

    fmt.Println(config.String())

    // Check if config has a key and print it
    if config.HasKey("a") {
        if a, err = config.MustGetBool("a"); err != nil {
            panic(err)
        }
        fmt.Printf("a = %v\n", a)
    }

    // Set new value (changes aren't written unless autosave was used)
    config.Set(false, "a")

    a = config.GetBool("a")
    fmt.Printf("a is now = %v\n", a)

    config.Reset()
    fmt.Println("Config reset")

    a = config.GetBool("a")
    fmt.Printf("a on disk still = %v\n", a)

    config.Set(false, "a")

    // Manually save changes
    config.Save()
    fmt.Println("Config saved")

    // More changes plus save
    config.Set("asdfasdf", "b")
    config.Set(4321, "c")
    config.Set("asdf", "d", 0)
    config.Set("asdf", "d", 1)
    config.Set([]string{"blah", "blah"}, "d")
    config.Save()

    b = config.GetString("b")
    c = config.GetInt64("c")
    d = config.GetStringArray("d")

    fmt.Printf("b = %s\n", b)
    fmt.Printf("c = %d\n", c)
    fmt.Printf("d = %v\n", d)

    // You can also reset changes (unless autosave was used)
    config.Set(true, "e", "bool")
    config.Set("test", "e", "string")
    config.Set(
        map[string]interface{}{"bool": true, "string": "test"},
        "e",
    )

    e = config.GetMap("e")
    fmt.Printf("e = %+v\n", e)

    config.Reset()
    fmt.Println("Config reset")

    e = config.GetMap("e")
    fmt.Printf("e = %+v\n", e)

    // Get nested keys
    f = config.GetFloat64("e", "float")
    fmt.Printf("e->float = %0.1f\n", f)

    // Only want to save the changes from default values?
    config.Set(false, "a")
    config.SaveDiff() // Diffs are calculated from last manual save

    // Reset to default values
    config.Default()
    config.Save()

    // Get sub-keys
    if keys, err = config.MustGetKeys("a"); err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println(keys)
    }

    keys = config.GetKeys("d")
    fmt.Println(keys)
    keys = config.GetKeys("e")
    fmt.Println(keys)
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/jsoncfg)
