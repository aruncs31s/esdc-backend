package utils

import "errors"

type reaction struct {
	ReactionMessage string
	err             error
}

func NewReaction(err error) reaction {
	return reaction{err: err}

}
func (r *reaction) Reaction() string {
	r.GetReactionBasedOnError()
	return r.ReactionMessage
}

func (r *reaction) GetReactionBasedOnError() {
	switch r.err {
	case ErrGeneratingJWT:
		r.ReactionMessage = "Could not create token"
	case ErrNotFound:
		r.ReactionMessage = "Resource not found"
	case ErrForbidden:
		r.ReactionMessage = "You don't have permission to access this resource"
	case ErrConflict:
		r.ReactionMessage = "Resource already exists"
	case ErrInternalServer:
		r.ReactionMessage = "Internal server error"
	default:
		r.ReactionMessage = "An unexpected error occurred"
	}
}

var (
	ErrBadRequest           = errors.New("bad request")
	ErrUserNotExists        = errors.New("user does not exists")
	ErrPasswordDoesNotMatch = errors.New("password does not match")
	ErrGeneratingJWT        = errors.New("error generating jwt")
	ErrForbidden            = errors.New("forbidden")
	ErrNotFound             = errors.New("not found")
	ErrConflict             = errors.New("conflict")
	ErrInternalServer       = errors.New("internal server error")
	ErrEmailorPasswordEmpty = errors.New("email or password is empty")
)

var (
	ErrDetailBadRequestJSONPayload = errors.New("bad json payload")
)
