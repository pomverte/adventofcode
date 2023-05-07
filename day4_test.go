package main

import (
	"testing"
)

func TestSectionsFullyContained(t *testing.T) {
	if sectionsFullyContained("2-4,6-8") {
		t.Errorf("sectionsFullyContained(\"2-4,6-8\") sections are not fully contained")
	}
	if !sectionsFullyContained("2-8,3-7") {
		t.Errorf("sectionsFullyContained(\"2-8,3-7\") sections are fully contained")
	}
	if !sectionsFullyContained("6-6,4-6") {
		t.Errorf("sectionsFullyContained(\"6-6,4-6\") sections are fully contained")
	}
	if !sectionsFullyContained("28-93,5-94") {
		t.Errorf("sectionsFullyContained(\"28-93,5-94\") sections are fully contained")
	}
}

func TestGetNbSectionsFullyContained(t *testing.T) {
	result := getNbSectionsFullyContained(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)
	expected := 2
	if result != expected {
		t.Errorf("getNbSectionsFullyContained(...) should have 2 fully contained")
	}
}
