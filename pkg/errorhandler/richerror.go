package errorhandler

import "gameapp/pkg/errorhandler/errorcodestatus"

type RichError struct {
	wrappedError error
	codeStatus   errorcodestatus.CodeStatus
	operation    string
	message      string
	metaData     map[string]interface{}
}

func (re RichError) Error() string {
	return re.message
}

func New() RichError {
	return RichError{}
}

func (re RichError) WithWrappedError(wrapper error) RichError {
	re.wrappedError = wrapper
	return re
}

func (re RichError) WithCodeStatus(code errorcodestatus.CodeStatus) RichError {
	re.codeStatus = code
	return re
}

func (re RichError) WithOperation(operation string) RichError {
	re.operation = operation
	return re
}

func (re RichError) WithMessage(message string) RichError {
	re.message = message
	return re
}

func (re RichError) WithMetaData(metaData map[string]interface{}) RichError {
	re.metaData = metaData
	return re
}
