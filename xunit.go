package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

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

type TestSummary struct {
	TotalPassed int       `json:"total_passed"`
	TotalFailed int       `json:"total_failed"`
	TotalTime   float32   `json:"total_time"`
	Results     TestCases `json:"results"`
}

type byModTime []os.FileInfo

func (f byModTime) Len() int           { return len(f) }
func (f byModTime) Less(i, j int) bool { return f[i].ModTime().Unix() < f[j].ModTime().Unix() }
func (f byModTime) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

type Collector struct {
	ReportsDir string
}

func NewCollector(reportsDir string) *Collector {
	return &Collector{
		ReportsDir: reportsDir,
	}
}

func (c *Collector) CollectResults() TestSummary {
	t := TestSummary{}

	files, _ := ioutil.ReadDir(c.ReportsDir)
	sort.Sort(byModTime(files))
	for _, f := range files {
		xmlFile, err := ioutil.ReadFile(c.ReportsDir + "/" + f.Name())
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
