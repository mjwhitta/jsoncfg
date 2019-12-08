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
    hl "gitlab.com/mjwhitta/hilighter"
    "gitlab.com/mjwhitta/jsoncfg"
)

// Create a jsoncfg object
var config = jsoncfg.New("/tmp/rcfile")

// Or if you want changes to be written to disk immediately
// var config = jsoncfg.NewAutosave("/tmp/rcfile")

// Initialize default values
func init() {
    config.SetDefault("myArray", []interface{}{})
    config.SetDefault("myBool", true)
    config.SetDefault("myStr", "asdf")
    config.SaveDefault()
    config.Reset()
}

func main() {
    defer func() {
        if r := recover(); r != nil {
            hl.PrintlnRed(r.(error).Error())
        }
    }()

    // Check if option exists, then get current value
    if config.Has("myBool") {
        hl.Printf("myBool exists and is %v\n", config.Get("myBool"))
    }

    // Set new value (changes aren't written unless autosave was used)
    config.Set("myBool", false)

    // Manually save changes
    config.Save()

    // More changes plus save
    config.Set("myArray", []int{1, 2, 3, 4})
    config.Save()

    // You can also reset changes
    config.Set("myBool", true)
    config.Reset()

    // Only want to save the changes from default values?
    config.Set("myStr", "blah")
    config.SaveDiff() // Diffs are calculated from last manual save
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/jsoncfg)
