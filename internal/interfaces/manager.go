package interfaces

type ManagedItem interface {
	// Returns the id
	ID() int
}

type ClientManager interface {
	// Add managed item
	AddClient(item ManagedItem)

	// Get managed item by id
	GetClient(id int) (ManagedItem, error)

	// Remove managed item by id
	DisconnectClient(id int)
}

type RoomManager interface {
	// Add managed item
	AddRoom(item ManagedItem)

	// Get managed item by id
	GetRoom(id int) (ManagedItem, error)

	// Remove managed item by id
	RemoveRoom(id int)
}
