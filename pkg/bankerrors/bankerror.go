package bankerrors

var (
	ErrorNotFound = NewBankError("we are unable to locate the requested bank account")
)

type BankError struct {
	Message string
}

func (pe *BankError) Error() string {
	return pe.Message
}

func NewBankError(message string) error {
	return &BankError{
		Message: message,
	}
}
