require "nokogiri"

class ResultsFetcher

  def self.failed?(test, result)
    if test.children.size > 0
      failed = test.children.at("//failure")
      result[:message] = failed.attribute("message").value
      result[:details] = failed.text
    end
  end

  def self.parse_testcases(testcases)
    results = []
    testcases.each do |t|
      result = "pass"

      testcase = { name: t.attribute("name").value, time: t.attribute("time").value, package: t.attribute("classname").value }

      result = "failed" if failed?(t, testcase)
      testcase[:result] = result
      results << testcase
    end
    results
  end


  def self.fetch(dir)
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
