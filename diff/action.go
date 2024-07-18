package diff

// ActionType represents the type of action to be performed.
type ActionType string

const (
	// CreateCategory represents an action to create a category.
	CreateCategory ActionType = "CreateCategory"

	// DeleteCategory represents an action to delete a category.
	DeleteCategory ActionType = "DeleteCategory"

	// CreateFeed represents an action to create a feed.
	CreateFeed ActionType = "CreateFeed"

	// DeleteFeed represents an action to delete a feed.
	DeleteFeed ActionType = "DeleteFeed"
)

// Action represents an action to be performed to sync the local and remote state.
type Action struct {
	Type          ActionType
	CategoryTitle string
	FeedURL       string
}
