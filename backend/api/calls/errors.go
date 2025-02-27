package calls

import "errors"

var (
	ErrFailedToDecode = errors.New("failed to decode")
	ErrFailedToFetch  = errors.New("failed to fetch")
)
