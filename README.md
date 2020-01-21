# jsoncfg

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

    config.SetDefault("a", true)
    config.SetDefault("b", "asdf")
    config.SetDefault("c", 1234)
    config.SetDefault("d", []string{"blah", "test"})
    config.SetDefault(
        "e",
        map[string]interface{}{"float": 1.2, "int": 0},
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

    // Check if config has a key and print it
    if config.Has("a") {
        fmt.Println(config.GetBool("a"))
    }

    // Set new value (changes aren't written unless autosave was used)
    config.Set("a", false)
    fmt.Println(config.GetBool("a"))

    // Manually save changes
    config.Save()

    // More changes plus save
    config.Set("b", "asdfasdf")
    config.Set("c", 4321)
    config.Set("d", []string{"asdf", "asdf"})
    config.Save()
    fmt.Println(config.GetString("b"))
    fmt.Println(config.GetInt64("c"))
    fmt.Println(config.GetStringArray("d"))

    // You can also reset changes (unless autosave was used)
    config.Set(
        "e",
        map[string]interface{}{"bool": true, "string": "test"},
    )
    fmt.Println(config.GetMap("e"))
    config.Reset()
    fmt.Println(config.GetMap("e"))

    // Get nested keys
    fmt.Println(config.GetFloat64("e", "float"))

    // Only want to save the changes from default values?
    config.Set("a", false)
    config.SaveDiff() // Diffs are calculated from last manual save

    // Reset to default values
    config.Default()
    config.Save()
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/jsoncfg)
