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

func TestSectionsOverlapped(t *testing.T) {
	if sectionsOverlapped("2-4,6-8") {
		t.Errorf("sections 2-4,6-8 don't overlap")
	}
	if sectionsOverlapped("2-3,4-5") {
		t.Errorf("sections 2-3,4-5 don't overlap")
	}

	if !sectionsOverlapped("5-7,7-9") {
		t.Errorf("sections 5-7,7-9 overlap")
	}
	if !sectionsOverlapped("2-8,3-7") {
		t.Errorf("sections 2-8,3-7 overlap")
	}
	if !sectionsOverlapped("6-6,4-6") {
		t.Errorf("sections 6-6,4-6 overlap")
	}
	if !sectionsOverlapped("2-6,4-8") {
		t.Errorf("sections 2-6,4-8 overlap")
	}
}

func TestGetNbSectionsOverlapped(t *testing.T) {
	result := getNbSectionsOverlapped(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)
	expected := 4
	if result != expected {
		t.Errorf("getNbSectionsOverlapped(...) returned %v, expected %v", result, expected)
	}
}
