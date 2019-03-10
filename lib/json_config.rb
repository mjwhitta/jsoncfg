require "fileutils"
require "json"
require "pathname"

class JSONConfig
    def clear
        @config = Hash.new
        write_config
    end

    def default_config
        # User will implement
        write_config
    end

    def get(key)
        case @config[key]
        when /^\s*false\s*$/i, false
            return false
        when /^\s*true\s*$/i, true
            return true
        else
            return @config[key]
        end
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
        case value
        when /^\s*false\s*$/i, false
            unsetbool(key)
        when /^\s*true\s*$/i, true
            setbool(key)
        else
            @config[key] = value
            write_config
        end
    end

    def setbool(key)
        @config[key] = true
        write_config
    end

    def unsetbool(key)
        @config[key] = false
        write_config
    end

    def write_config
        FileUtils.mkdir_p(@config_file.dirname)
        File.open(@config_file, "w") do |file|
            file.write(JSON.pretty_generate(@config))
        end
    end
    private :write_config
end
