package authorizer

type Account struct {
}

func (a *Authorizer) AuthorizeAccount(secret []byte) (Account, error) {
	return Account{}, nil
}

func (a *Authorizer) CheckAuthCode(code []byte) bool {
	return true
}
