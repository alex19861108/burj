package errors

import "errors"

var (
	ERROR_DB_NOT_FOUND_NODE = errors.New("failed to find a nonexistent node")
)
