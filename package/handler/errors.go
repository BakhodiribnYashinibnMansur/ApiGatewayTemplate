package handler

import "errors"

var (

	// ErrNotFound ...
	ErrNotFound = errors.New("not found")

	// ErrInternal ...
	ErrInternalServer = errors.New("internal server error")

	// ErrAlreadyExists ...
	ErrAlreadyExists = errors.New(" already exists")

	// ErrUsernameExists ...
	ErrUsernameExists = errors.New("username exists")

	//ErrPhoneExists
	ErrPhoneExists = errors.New("phone exists")

	// ErrEmailExists ...
	ErrEmailExists = errors.New("email exists")

	// ErrInvalidField ...
	ErrInvalidField = errors.New("invalid field for username/email")

	// ErrMaximumAmount ...
	ErrMaximumAmount = errors.New("maximum amount")

	// ErrNotEnoughCash ...
	ErrNotEnoughCash = errors.New("not enough cash")

	// ErrInvalidFieldForOperations ...
	ErrInvalidFieldForOperations = errors.New("invalid field for operation type")

	//ErrNotValidPhone
	ErrNotValidPhone = errors.New("invalid field for phone type")

	//ErrNotValidFirstName
	ErrNotValidFirstName = errors.New("invalid field for firstname type")

	//ErrNotValidLastName
	ErrNotValidLastName = errors.New("invalid field for lastname type")

	//ErrorNotValidPassword
	ErrorNotValidPassword = errors.New("invalid field for Password")

	//ErrorSignInCorrect
	ErrorSignInCorrect = errors.New("username or password is incorrect")

	ErrorNotANumberLimit = errors.New("query Limit not a number")

	ErrorNotANumberOffset = errors.New("query Offset not a number")

	ErrorParamIsEmpty = errors.New("query Parameter is empty")
)
