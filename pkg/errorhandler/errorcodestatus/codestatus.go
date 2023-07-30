package errorcodestatus

type CodeStatus int

const (
	EntityNotFound CodeStatus = iota + 1
	InvalidProcess
	Forbidden
	InternalError
)
