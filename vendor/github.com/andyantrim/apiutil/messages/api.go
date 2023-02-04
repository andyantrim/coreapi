package messages

// GenericSuccess sucess message struct
type GenericSuccess struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	OK      bool   `json:"ok"`
}

// NewGenericSuccess builds new success message
func NewGenericSuccess(name, message string) GenericSuccess {
	return GenericSuccess{
		Name:    name,
		Message: message,
		OK:      true,
	}
}
