package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/tedsuo/rata"
)

const viewsDir = "public"
const reportDir = "./logs"

type TestCases []TestCase

type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Tests     string     `xml:"tests,attr"`
	Time      string     `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}
type TestCase struct {
	Name      string `xml:"name,attr" json:"name"`
	Time      string `xml:"time,attr" json:"time"`
	Classname string `xml:"classname,attr" json:"package"`
	Failure   string `xml:"failure"`
	Result    string `json:"result"`
}

type testSummary struct {
	TotalPassed int       `json:"total_passed"`
	TotalFailed int       `json:"total_failed"`
	TotalTime   float32   `json:"total_time"`
	Results     TestCases `json:"results"`
}

type byModTime []os.FileInfo

func (f byModTime) Len() int           { return len(f) }
func (f byModTime) Less(i, j int) bool { return f[i].ModTime().Unix() < f[j].ModTime().Unix() }
func (f byModTime) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

func collectResults() testSummary {
	t := testSummary{}

	files, _ := ioutil.ReadDir(reportDir)
	sort.Sort(byModTime(files))
	for _, f := range files {
		xmlFile, err := ioutil.ReadFile(reportDir + "/" + f.Name())
		if err != nil {
			fmt.Println("Error opening file:", err)
			break
		}

		ts := TestSuite{}
		xml.Unmarshal(xmlFile, &ts)

		for _, tc := range ts.TestCases {
			if tc.Failure != "" {
				tc.Result = "failed"
				t.TotalFailed++
			} else {
				tc.Result = "pass"
				t.TotalPassed++
			}

			f, _ := strconv.ParseFloat(tc.Time, 32)
			t.TotalTime += float32(f)

			t.Results = append(t.Results, tc)

		}
	}
	return t
}

func newTestsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(collectResults())
	})
}

func main() {

	routes := rata.Routes{
		{Name: "get_index", Method: "GET", Path: "/"},
		{Name: "get_tests", Method: "GET", Path: "/tests"},
	}

	handlers := map[string]http.Handler{
		"get_index": http.FileServer(http.Dir(viewsDir)),
		"get_tests": newTestsHandler(),
	}

	router, err := rata.NewRouter(routes, handlers)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
