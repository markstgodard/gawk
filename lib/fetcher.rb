require "nokogiri"

module Gawk
  class ResultsFetcher

    def self.failed?(test, result)
      if test.children.size > 0
        failed = test.children.at("//failure")
        result[:message] = failed.attribute("message").value
        result[:details] = failed.text
      end
    end

    def self.parse_testcases(summary, testcases)
      testcases.each do |t|
        result = "pass"

        testcase = { name: t.attribute("name").value, time: t.attribute("time").value, package: t.attribute("classname").value }

        if failed?(t, testcase)
          result = "failed"
          summary[:total_failed] += 1
        else
          summary[:total_passed] += 1
        end

        testcase[:result] = result

        summary[:total_time] += testcase[:time].to_f
        summary[:results] << testcase
      end
    end

    def self.fetch(dir)
      summary = { total_time: 0.00, total_passed: 0, total_failed: 0, results: [] }

      files_oldest_first = Dir[dir + "/*.xml"].sort_by{ |f| File.mtime(f) }

      files_oldest_first.each do |f|
        xml = Nokogiri::XML(File.read(f))
        testcases = xml.xpath("//testcase")
        parse_testcases(summary, testcases)
      end

      summary
    end

  end
end
