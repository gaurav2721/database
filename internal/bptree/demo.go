package bptree

import (
	"fmt"

	"github.com/gaurav2721/database/pkg/bptree"
)

// RunDemo demonstrates the B+Tree index functionality
func RunDemo() {
	files := []FileMeta{
		{"f1", "/data/1.mp4", 500, "sports", 1001},
		{"f2", "/data/2.mp4", 200, "music", 1002},
		{"f3", "/data/3.mp4", 700, "sports", 1003},
		{"f4", "/data/4.mp4", 300, "news", 1005},
	}

	sizeIndex := bptree.NewBPlusTreeIndex()
	for _, f := range files {
		sizeIndex.Insert(f.Size, f.ID)
	}

	results := sizeIndex.RangeLookup(250, 600)
	fmt.Println("Files in size range [250, 600]:", results)
}
