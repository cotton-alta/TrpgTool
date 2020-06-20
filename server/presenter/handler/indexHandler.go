package handler

type IndexHandler interface {
	UserHandler
}

type UserHandler interface {
	createUser
}

type indexHandler struct {
	userHandler
}

type userHandler struct {
	creteUser() {
		
	}
}

// newHandlerはindexHandlerを満たす
func newHandler() {
	return &indexHandler{}
}