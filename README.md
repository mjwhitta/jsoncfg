# json_config

## What is this?

This Go module allows you to read/write and get/set configuration
options from/to a JSON file.

## How to install

Open a terminal and run the following:

```
$ go get -u gitlab.com/mjwhitta/json_config
```

## Usage

```
package main

import (
    "gitlab.com/mjwhitta/json_config"
)

TODO
# Sample implementation
class MyConfig < JSONConfig
    extend JSONConfig::Keys

    add_key("myarray")
    add_bool_key("mybool")
    add_key("mystr")

    def initialize(file = nil)
        file ||= "~/.config/somedir/rc"
        @defaults = {
            "myarray" => Array.new,
            "myboolkey" => true,
            "mystr" => "asdf"
        }
        autosave = false # Should changes be written immediately?
        super(file, autosave)
    end
end

# Sample usage

## Create cofnig
config = MyConfig.new

## Check if option exists, then get current value
if (config.mystr?)
    puts "mystr exists and is equal to #{config.get_mystr}"
end

## Set new value (changes aren't written unless autosave was true)
config.set_mystr("test")

## Manually save changes
config.save

## More changes plus save
config.set_myarray([1, 2, 3, 4])
config.save

## Only want to save the changes from default values?
config.savediff # Diffs are calculated from last manual save

## You can also reset changes
config.reset
```

## Links

- [Source](https://gitlab.com/mjwhitta/json_config)
