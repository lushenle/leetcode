package recursion

type Item struct {
	ID    int
	Type  string
	Child *Item
}

type ItemClassifier interface {
	IsDoll() bool
}

func (it *Item) IsDoll() bool {
	return it.Type == "doll"
}

func findDiamond(item Item) Item {
	if item.IsDoll() {
		return findDiamond(*item.Child)
	}

	return item
}
