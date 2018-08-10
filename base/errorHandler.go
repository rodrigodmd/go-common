package base

type ErrorHandler struct {
	log Logger
	Err error
}

func (e *ErrorHandler) HandleError(err ...error) {
	if len(err) > 0 {
		e.Err = err[0]
	}
	if e.Err != nil {
		e.LogError()
	}
}

func (e ErrorHandler) LogError() {
	e.log.ErrorObj(e.Err)
}
