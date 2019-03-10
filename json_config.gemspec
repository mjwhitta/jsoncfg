Gem::Specification.new do |s|
    s.name = "json_config"
    s.version = "0.3.2"
    s.date = Time.new.strftime("%Y-%m-%d")
    s.summary = "Read/write from/to a JSON config file"
    s.description = "Read/write from/to a JSON config file."
    s.authors = [ "Miles Whittaker" ]
    s.email = "mjwhitta@gmail.com"
    s.files = Dir["lib/**/*.rb"]
    s.homepage = "https://gitlab.com/mjwhitta/json_config"
    s.license = "GPL-3.0"
    s.add_development_dependency("rake", "~> 12.3", ">= 12.3.2")
end
