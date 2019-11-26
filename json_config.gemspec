Gem::Specification.new do |s|
    s.add_development_dependency("rake", "~> 13.0", ">= 13.0.0")
    s.authors = ["Miles Whittaker"]
    s.date = Time.new.strftime("%Y-%m-%d")
    s.description = "Read/write from/to a JSON config file."
    s.email = "mj@whitta.dev"
    s.files = Dir["lib/**/*.rb"]
    s.homepage = "https://gitlab.com/mjwhitta/json_config"
    s.license = "GPL-3.0"
    s.metadata = {"source_code_uri" => s.homepage}
    s.name = "json_config"
    s.summary = "Read/write from/to a JSON config file"
    s.version = "1.1.1"
end
