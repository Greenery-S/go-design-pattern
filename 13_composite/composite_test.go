package composite

import "testing"

func TestComposite(t *testing.T) {
	file1 := NewComponent(LeafNode, "msg", "hello world")
	file2 := NewComponent(LeafNode, "bool", true)
	file3 := NewComponent(LeafNode, "magic#", 114514)
	file4 := NewComponent(LeafNode, "pi", 3.14)
	dir1 := NewComponent(CompositeNode, "dir1", nil)
	dir2 := NewComponent(CompositeNode, "dir2", nil)
	rootDir := NewComponent(CompositeNode, "root dir", nil)

	dir2.AddChild(file2)
	dir2.AddChild(file3)

	dir1.AddChild(dir2)
	dir1.AddChild(file4)

	rootDir.AddChild(dir1)
	rootDir.AddChild(file1)

	rootDir.Print(">")
}
