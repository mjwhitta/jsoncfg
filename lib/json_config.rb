require "fileutils"
require "json"
require "pathname"

class JSONConfig
    module Keys
        def add_bool_key(key)
            define_method "#{key}?" do
                return get(key)
            end
            define_method key do
                set(key, true)
            end
            define_method "no_#{key}" do
                set(key, false)
            end
        end

        def add_key(key)
            define_method "#{key}?" do
                return !@config[key].nil?
            end
            define_method "get_#{key}" do
                return get(key)
            end
            define_method "set_#{key}" do |val|
                set(key, val)
            end
        end
    end

    def clear
        @config = Hash.new
        @diff = Hash.new
        write_config
    end

    def default
        @config = @defaults.clone
        @diff = @defaults.clone
        write_config(true)
    end

    def get(key)
        case @config[key]
        when /^\s*false\s*$/i, false
            return false
        when /^\s*true\s*$/i, true
            return true
        when /^\s*\d+\s*$/
            return @config[key].to_i
        when /^\s*\d+\.\d+\s*$/
            return @config[key].to_f
        else
            return @config[key]
        end
    end

    def getdiff(key)
        case @diff[key]
        when /^\s*false\s*$/i, false
            return false
        when /^\s*true\s*$/i, true
            return true
        else
            return @diff[key]
        end
    end

    def initialize(file, autosave = true)
        @defaults ||= Hash.new
        @autosave = autosave
        @config_file = Pathname.new(file).expand_path
        read_config
    end

    def read_config
        if (!@config_file.exist? && !@config_file.symlink?)
            @config = @defaults.clone
            write_config(true)
        end
        @config = JSON.parse(File.read(@config_file))
        @diff = @defaults.clone
    end
    private :read_config

    def reset
        read_config
    end

    def save
        @diff = @defaults.clone
        write_config(true)
    end

    def savediff
        @diff = @config
        write_config(true)
    end

    def set(key, value)
        case value
        when /^\s*false\s*$/i, false
            unsetbool(key)
        when /^\s*true\s*$/i, true
            setbool(key)
        when /^\s*\d+\s*$/
            @config[key] = value.to_i
            @diff[key] = value.to_i
            write_config
        when /^\s*\d+\.\d+\s*$/
            @config[key] = value.to_f
            @diff[key] = value.to_f
            write_config
        else
            @config[key] = value
            @diff[key] = value
            write_config
        end
    end

    def setbool(key)
        @config[key] = true
        @diff[key] = true
        write_config
    end

    def unsetbool(key)
        @config[key] = false
        @diff[key] = false
        write_config
    end

    def write_config(force = false)
        return if (!@autosave && !force)

        FileUtils.mkdir_p(@config_file.dirname)
        File.open(@config_file, "w") do |file|
            file.write(JSON.pretty_generate(@config))
        end
    end
    private :write_config
end
