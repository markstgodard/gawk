require "rubygems"
require "bundler/setup"
require "sinatra"
require "json"

$LOAD_PATH.unshift("#{File.dirname(__FILE__)}/lib")
Dir.glob("#{File.dirname(__FILE__)}/lib/*.rb") { |lib| require File.basename(lib, '.*') }

configure do
  set :views, "#{File.dirname(__FILE__)}/views"
  set :show_exceptions, :after_handler
end

configure :production, :development do
  enable :logging
end

get "/tests" do
  content_type :json
  ResultsFetcher.fetch.to_json
end

# main page
get "/" do
  erb :index
end
