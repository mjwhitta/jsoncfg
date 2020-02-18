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

    var a bool
    var b string
    var c int64
    var d []string
    var e = map[string]interface{}{}
    var err error
    var f float64

    fmt.Println(config.String())

    // Check if config has a key and print it
    if config.Has("a") {
        if a, err = config.GetBool("a"); err != nil {
            panic(err)
        }
        fmt.Printf("a = %v\n", a)
    }

    // Set new value (changes aren't written unless autosave was used)
    config.Set("a", false)

    a, _ = config.GetBool("a")
    fmt.Printf("a is now = %v\n", a)

    config.Reset()
    fmt.Println("Config reset")

    a, _ = config.GetBool("a")
    fmt.Printf("a on disk still = %v\n", a)

    config.Set("a", false)

    // Manually save changes
    config.Save()
    fmt.Println("Config saved")

    // More changes plus save
    config.Set("b", "asdfasdf")
    config.Set("c", 4321)
    config.Set("d", []string{"asdf", "asdf"})
    config.Save()

    b, _ = config.GetString("b")
    c, _ = config.GetInt64("c")
    d, _ = config.GetStringArray("d")

    fmt.Printf("b = %s\n", b)
    fmt.Printf("c = %d\n", c)
    fmt.Printf("d = %v\n", d)

    // You can also reset changes (unless autosave was used)
    config.Set(
        "e",
        map[string]interface{}{"bool": true, "string": "test"},
    )

    e, _ = config.GetMap("e")
    fmt.Printf("e = %+v\n", e)

    config.Reset()
    fmt.Println("Config reset")

    e, _ = config.GetMap("e")
    fmt.Printf("e = %+v\n", e)

    // Get nested keys
    f, _ = config.GetFloat64("e", "float")
    fmt.Printf("e->float = %0.1f\n", f)

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
