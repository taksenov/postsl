/*
Copyright © 2025 taksenov@gmail.com
*/

// Package filesystem -- utilities for working with the file system.
package filesystem

import (
	"os"
	"testing"
)

func TestGetManifestNames(t *testing.T) {
	dir := "tst"

	err := os.Mkdir(dir, 0o777)
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(dir)

	// Create __default__ dir
	err = os.MkdirAll(dir+"/__default__", 0o777)
	if err != nil {
		t.Fatalf("Failed to create '__default__' directory: %v", err)
	}

	// Create test files
	files := []string{dir + "/file1.txt", dir + "/file2.txt"}
	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(f.Name())
	}

	// Test GetManifestNames function
	names, _ := GetManifestNames(dir)
	expected := []string{"file1.txt", "file2.txt"}
	if len(names) != len(expected) {
		t.Errorf("Unexpected length of manifest names. Expected %d, got %d", len(expected), len(names))
	}

	for i, name := range names {
		if name != expected[i] {
			t.Errorf("Unexpected manifest name. Expected %s, got %s", expected[i], name)
		}
	}
}

func TestGetManifestNamesErrorHandling(t *testing.T) {
	dir := "/nonexistent"

	names, err := GetManifestNames(dir)
	if err != nil {
		if len(names) != 0 {
			t.Errorf("Expected empty manifest names, got %v", names)
		}
	}
}
