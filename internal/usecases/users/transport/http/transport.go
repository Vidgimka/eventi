package user_transport_http

type UserHTTPHandler struct {
	userService UserService
}

type UserService struct {
}

func NewUserHTTPHandler(userService UserService) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: userService,
	}
}
