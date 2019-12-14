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
var config = jsoncfg.New("/tmp/asdf/rcfile")

// Or if you want changes to be written to disk immediately
// var config = jsoncfg.NewAutosave("/tmp/asdf/rcfile")

// Initialize default values
func init() {
    config.SetDefault("myArray", []int{})
    config.SetDefault("myBool", true)
    config.SetDefault("myMap", map[string]uint64{})
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
        hl.Printf(
            "myBool exists and is %v\n",
            config.GetBool("myBool"),
        )
    }

    // Set new value (changes aren't written unless autosave was used)
    config.Set("myBool", false)
    hl.Printf("myBool is now %v\n", config.GetBool("myBool"))

    // Manually save changes
    config.Save()

    // More changes plus save
    hl.Printf("myArray is %v\n", config.GetIntArray("myArray"))
    config.Set("myArray", []int{1, 2, 3, 4})
    hl.Printf("myArray is now %v\n", config.GetIntArray("myArray"))
    config.Save()

    // You can also reset changes
    hl.Printf("myMap is %v\n", config.GetUint64Map("myMap"))
    config.Set("myMap", map[string]uint64{"asdf": 1, "blah": 2})
    hl.Printf("myMap is now %v\n", config.GetUint64Map("myMap"))
    config.Reset()
    hl.Printf("myMap is now %v\n", config.GetUint64Map("myMap"))

    // Only want to save the changes from default values?
    config.Set("myStr", "blah")
    config.SaveDiff() // Diffs are calculated from last manual save

    // Reset to default values
    config.Default()
    config.Save()
}
```

## Links

- [Source](https://gitlab.com/mjwhitta/jsoncfg)
