package tests

import (
	"reflect"
	"testing"

	"github.com/gaurav2721/database/pkg/bptree"
)

func TestBPlusTreeIndex_InsertAndRangeLookup(t *testing.T) {
	index := bptree.NewBPlusTreeIndex()
	index.Insert(100, "f1")
	index.Insert(200, "f2")
	index.Insert(300, "f3")
	index.Insert(150, "f4")

	result := index.RangeLookup(120, 250)
	expected := []string{"f4", "f2"} // 150, 200

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBPlusTreeIndex_EmptyRange(t *testing.T) {
	index := bptree.NewBPlusTreeIndex()
	index.Insert(500, "f1")
	index.Insert(800, "f2")

	result := index.RangeLookup(100, 300)
	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
}

func TestBPlusTreeIndex_InsertDuplicateKeys(t *testing.T) {
	index := bptree.NewBPlusTreeIndex()
	index.Insert(100, "f1")
	index.Insert(100, "f2")
	index.Insert(200, "f3")

	result := index.RangeLookup(100, 100)
	expected := []string{"f1", "f2"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestBPlusTreeIndex_NewIndex(t *testing.T) {
	index := bptree.NewBPlusTreeIndex()

	// Test that the index is properly initialized
	if index == nil {
		t.Error("Expected non-nil index")
	}
}
