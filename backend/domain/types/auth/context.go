package auth

type Context struct {
	UserID    int
	Clearance Clearance
	Token     string
}

func NewContext(user *User, token string) *Context {
	return &Context{
		UserID:    user.ID,
		Clearance: user.Clearance,
		Token:     token,
	}
}

func NewAdminContext() *Context {
	return &Context{
		Clearance: ClearanceTopSecret,
	}
}
