package interfaces

// This is a wrapper for the room
// so we can handle the players inside more easily
type Room interface {
	// Add member to the room by id
	Join(id int)

	// Remove member from the room by id
	Leave(id int)

	// Broadcast to everyone whos in the same room (including the source)
	Broadcast(arg ...string)

	// Announce to others whos in the same room (not including the source)
	Announce(id int, arg ...string)

	// Just send to source
	Send(id int, arg ...string)

	// GETTERS //
	// Returns if the room is empty
	IsEmpty() bool

	// Returns of the room is running
	IsRunning() bool

	// SETTERS //
	// Set the room running or not
	SetRunning(bool)
}
