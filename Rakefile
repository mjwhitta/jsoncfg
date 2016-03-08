task :default => :gem

desc "Clean up"
task :clean do
    system("rm -f *.gem Gemfile.lock")
    system("chmod -R go-rwx lib")
end

desc "Build gem"
task :gem do
    system("chmod -R u=rwX,go=rX lib")
    system("gem build -V *.gemspec")
end

desc "Build and install gem"
task :install => :gem do
    system("gem install *.gem")
end

desc "Push gem to rubygems.org"
task :push => [:clean, :gem] do
    system("gem push *.gem")
end
