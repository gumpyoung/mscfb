package mscfb

import (
	"io"
	"os"
	"testing"
)

var (
	testDoc = "test/test.doc"
	testXls = "test/test.xls"
	testPpt = "test/test.ppt"
	testMsg = "test/test.msg"
	entries = []*DirectoryEntry{
		&DirectoryEntry{Name: "Root Node",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: 1},
		},
		&DirectoryEntry{Name: "Alpha",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: 2, childID: noStream},
		},
		&DirectoryEntry{Name: "Bravo",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: 3, childID: 5},
		},
		&DirectoryEntry{Name: "Charlie",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: 7},
		},
		&DirectoryEntry{Name: "Delta",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: noStream},
		},
		&DirectoryEntry{Name: "Echo",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: 4, rightSibID: 6, childID: 9},
		},
		&DirectoryEntry{Name: "Foxtrot",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: noStream},
		},
		&DirectoryEntry{Name: "Golf",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: 10},
		},
		&DirectoryEntry{Name: "Hotel",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: noStream},
		},
		&DirectoryEntry{Name: "Indigo",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: 8, rightSibID: noStream, childID: 11},
		},
		&DirectoryEntry{Name: "Jello",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: noStream},
		},
		&DirectoryEntry{Name: "Kilo",
			fn:                   mockFN,
			directoryEntryFields: &directoryEntryFields{leftSibID: noStream, rightSibID: noStream, childID: noStream},
		},
	}
)

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func mockFN(e *DirectoryEntry) {}

func empty(sl []byte) bool {
	for _, v := range sl {
		if v != 0 {
			return false
		}
	}
	return true
}

func testFile(t *testing.T, path string) {
	file, _ := os.Open(path)
	defer file.Close()
	doc, err := New(file)
	if err != nil {
		t.Errorf("Error opening file; Returns error: ", err)
	}
	for entry, _ := doc.Next(); entry != nil; entry, _ = doc.Next() {
		buf := make([]byte, 512)
		_, err := doc.Read(buf)
		if err != nil && err != ErrNoStream && err != io.EOF {
			t.Errorf("Error reading entry name, %v", entry.Name)
		}
		if len(entry.Name) < 1 {
			t.Errorf("Error reading entry name")
		}
	}

}

func TestTraverse(t *testing.T) {
	r := new(Reader)
	r.Entries = entries
	if r.traverse() != nil {
		t.Error("Error traversing")
	}
	expect := []int{0, 1, 2, 4, 5, 8, 9, 11, 6, 3, 7, 10}
	for i, v := range r.indexes {
		if v != expect[i] {
			t.Errorf("Error traversing: expecting %d at index %d; got %d", expect[i], i, v)
		}
	}
	if r.Entries[10].Path[0] != "Charlie" {
		t.Errorf("Error traversing: expecting Charlie got %s", r.Entries[10].Path[0])
	}
	if r.Entries[10].Path[1] != "Golf" {
		t.Errorf("Error traversing: expecting Golf got %s", r.Entries[10].Path[1])
	}
}

func TestWord(t *testing.T) {
	testFile(t, testDoc)
}

func TestXls(t *testing.T) {
	testFile(t, testXls)
}

func TestPpt(t *testing.T) {
	testFile(t, testPpt)
}

func TestMsg(t *testing.T) {
	testFile(t, testMsg)
}
