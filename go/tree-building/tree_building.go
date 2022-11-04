package tree

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var root *Node

	for _, record := range records {
		// create the root
		if root == nil {
			root = &Node{ID: record.ID, Children: []*Node{}}
		} else {
			// insert the node
			// update parent
			if record.Parent == 0 && record.ID == 0 {
				temp := root
				root = &Node{ID: record.ID, Children: []*Node{temp}}
			} else {
				// insert where you can
			}
		}
	}

	return root, nil
}
