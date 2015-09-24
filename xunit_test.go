package main

import "testing"

func TestNewCollectorDirNotExist(t *testing.T) {
	_, err := NewCollector("dir-not-exist")
	if err == nil {
		t.Error("Expected to fail due to dir not exist")
	}
}

func TestNewCollectorDirExist(t *testing.T) {
	_, err := NewCollector("testdata")
	if err != nil {
		t.Error("Expected to create since dir exists")
	}
}

func TestCollectResults(t *testing.T) {
	c, err := NewCollector("testdata")
	if err != nil {
		t.Error("Expected to create since dir exists")
	}

	ts := c.CollectResults()

	if ts.TotalPassed != 5 {
		t.Error("Expect 5, got ", ts.TotalPassed)
	}
	if ts.TotalFailed != 1 {
		t.Error("Expect 1, got ", ts.TotalFailed)
	}
	if ts.TotalTime != 18.50 {
		t.Error("Expect 18.50, got ", ts.TotalTime)
	}

	if len(ts.Results) != 6 {
		t.Error("Expect results size to be 6, got ", len(ts.Results))
	}

	if ts.Results[0].Name != "Test case 1" {
		t.Error("Expect 'Test case 1', got ", ts.Results[0].Name)
	}

	// one with a failure
	if ts.Results[1].Name != "Test case 2" {
		t.Error("Expect 'Test case 2', got ", ts.Results[1].Name)
	}
	if ts.Results[1].Failure.Value != "AssertionError 0 == 1" {
		t.Error("Expect 'AssertionError 0 == 1', got ", ts.Results[1].Failure.Value)
	}

	if ts.Results[5].Name != "Test case 6" {
		t.Error("Expect 'Test case 6', got ", ts.Results[5].Name)
	}
}
