# JSONConfig

## What is this?

This ruby gem allows you to read/write and get/set configuration
options from/to a JSON file.

## How to install

```
$ gem install json_config
```

## Usage

```ruby
#!/usr/bin/env ruby

require "json_config"

class MyConfig < JSONConfig
    def default_config
        set("mykey", "myval")
    end
end

config = MyConfig.new("./test_config")
puts config.get("mykey")
config.set("newkey", "newval")
```

## Links

- [Source](https://gitlab.com/mjwhitta/json_config)
- [RubyGems](https://rubygems.org/gems/json_config)

## TODO

- Better README
- RDoc
