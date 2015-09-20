class ResultsFetcher
  class TestCase < Struct.new(:name, :time, :category, :result)
  end

  def self.fetch
    results = []
    results << { name: "Buildpacks binary makes the app reachable via its bound route", time: "18.650551275", package: "Applications", result: "pass" }
    results << { name: "Buildpacks binary makes the app reachable via its bound route 2", time: "12.65275", package: "Applications", result: "pass" }
    results << { name: "Buildpacks binary makes the app reachable via its bound route 3", time: "8.650551275", package: "Applications", result: "pass" }
    results << { name: "Buildpacks binary makes the app reachable via its bound route 4", time: "275", package: "Applications", result: "pass" }
    results << { name: "Buildpacks binary makes the app reachable via its bound route 5", time: "1275", package: "Applications", result: "pass" }
    results
  end

end
