package diff

// ActionSorter sorts actions by type and then by relevant fields within each type.
type ActionSorter []Action

// Len implements the sort.Interface.
func (a ActionSorter) Len() int {
	return len(a)
}

// Swap implements the sort.Interface.
func (a ActionSorter) Swap(i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less implements the sort.Interface.
func (a ActionSorter) Less(i int, j int) bool { //nolint:varnamelen
	// Define the order of action types.
	order := map[ActionType]int{
		DeleteFeed:     0,
		DeleteCategory: 1,
		CreateCategory: 2,
		CreateFeed:     3,
	}

	// First, sort by action type.
	if order[a[i].Type] != order[a[j].Type] {
		return order[a[i].Type] < order[a[j].Type]
	}

	// Then, sort within each action type by the relevant field.
	switch a[i].Type {
	case DeleteFeed:
		return a[i].FeedURL < a[j].FeedURL

	case DeleteCategory:
		return a[i].CategoryTitle < a[j].CategoryTitle

	case CreateCategory:
		return a[i].CategoryTitle < a[j].CategoryTitle

	case CreateFeed:
		return a[i].FeedURL < a[j].FeedURL

	default:
		return false
	}
}
