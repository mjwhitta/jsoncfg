require "json"
require "pathname"

class JSONConfig
    def default_config
        # User will implement
    end

    def get(key)
        return @config[key]
    end

    def initialize(file)
        @config_file = Pathname.new(file).expand_path
        read_config
    end

    def read_config
        if (!@config_file.exist? && !@config_file.symlink?)
            @config = Hash.new
            default_config
        end

        @config = JSON.parse(File.read(@config_file))
    end
    private :read_config

    def set(key, value)
        @config[key] = value
        write_config
    end

    def write_config
        File.open(@config_file, "w") do |file|
            file.write(JSON.pretty_generate(@config))
        end
    end
    private :write_config
end
