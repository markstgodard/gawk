require "rubygems"
require "bundler/setup"
require "sinatra"
require "sinatra/base"
require "sinatra/config_file"
require "json"

$LOAD_PATH.unshift("#{File.dirname(__FILE__)}/lib")
Dir.glob("#{File.dirname(__FILE__)}/lib/*.rb") { |lib| require File.basename(lib, '.*') }

module Gawk
  class App < Sinatra::Base
    register Sinatra::ConfigFile
    config_file 'gawk.yml'

    configure do
      set :views, "#{File.dirname(__FILE__)}/views"
      set :show_exceptions, :after_handler
    end

    configure :production, :development do
      enable :logging
    end

    get "/tests" do
      content_type :json
      Gawk::ResultsFetcher.fetch(settings.reports_dir).to_json
    end

    # main page
    get "/" do
      erb :index
    end
  end
end
