package auth

type Authentication struct {
	UserId int
	Token  string
}

var authSingleton *Authentication

func (auth Authentication) SetAuthentication() *Authentication {
	authSingleton = &auth

	return authSingleton
}

func GetAuthentication() Authentication {
	return *authSingleton
}
