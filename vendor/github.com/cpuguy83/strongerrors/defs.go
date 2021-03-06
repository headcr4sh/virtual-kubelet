package strongerrors

// ErrNotFound signals that the requested object doesn't exist
type ErrNotFound interface {
	NotFound()
}

// ErrInvalidArgument signals that the user input is invalid
type ErrInvalidArgument interface {
	InvalidArgument()
}

// ErrConflict signals that some internal state conflicts with the requested action and can't be performed.
// A change in state should be able to clear this error.
type ErrConflict interface {
	Conflict()
}

// ErrUnauthorized is used to signify that the user is not authorized to perform a specific action
type ErrUnauthorized interface {
	Unauthorized()
}

// ErrUnauthenticated is used to indicate that the caller cannot be identified.
type ErrUnauthenticated interface {
	Unauthenticated()
}

// ErrUnavailable signals that the requested action/subsystem is not available.
type ErrUnavailable interface {
	Unavailable()
}

// ErrForbidden signals that the requested action cannot be performed under any circumstances.
// When a ErrForbidden is returned, the caller should never retry the action.
type ErrForbidden interface {
	Forbidden()
}

// ErrSystem signals that some internal error occurred.
// An example of this would be a failed mount request.
type ErrSystem interface {
	System()
}

// ErrNotModified signals that an action can't be performed because it's already in the desired state
type ErrNotModified interface {
	NotModified()
}

// ErrAlreadyExists is a special case of ErrNotModified which signals that the desired object already exists
type ErrAlreadyExists interface {
	AlreadyExists()
}

// ErrNotImplemented signals that the requested action/feature is not implemented on the system as configured.
type ErrNotImplemented interface {
	NotImplemented()
}

// ErrUnknown signals that the kind of error that occurred is not known.
type ErrUnknown interface {
	Unknown()
}

// ErrCancelled signals that the action was cancelled.
type ErrCancelled interface {
	Cancelled()
}

// ErrDeadline signals that the deadline was reached before the action completed.
type ErrDeadline interface {
	DeadlineExceeded()
}

// ErrExhausted indicates that the action cannot be performed because some resource is exhausted.
type ErrExhausted interface {
	Exhausted()
}

// ErrDataLoss indicates that data was lost or there is data corruption.
type ErrDataLoss interface {
	DataLoss()
}
