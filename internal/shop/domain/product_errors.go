package domain

type ErrProductNotFound struct {
	Err string
}

func (e *ErrProductNotFound) Error() string {

	return e.Err
}
