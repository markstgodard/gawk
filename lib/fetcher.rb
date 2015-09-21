require "nokogiri"

class ResultsFetcher

  def self.parse_testcases(testcases)
    results = []
    testcases.each do |t|
      puts "test case: #{t} #{t.class}"
      results << { name: t.attribute("name"), time: t.attribute("time"), package: t.attribute("classname"), result: "pass" }
    end
    results
  end

  def self.fetch
    dir = "./logs"
    results = []

    files_oldest_first = Dir[dir + "/*.xml"].sort_by{ |f| File.mtime(f) }

    files_oldest_first.each do |f|
      xml = Nokogiri::XML(File.read(f))
      testcases = xml.xpath("//testcase")
      results.concat(parse_testcases(testcases))
    end
    results
  end

end
