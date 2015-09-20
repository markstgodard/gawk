require "rubygems"
require "bundler/setup"
require "sinatra"

configure do
  set :views, "#{File.dirname(__FILE__)}/views"
  set :show_exceptions, :after_handler
end

configure :production, :development do
  enable :logging
end

# main page
get "/" do
  erb :index
end
