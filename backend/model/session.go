package model

type sessionMap map[string]*User

func (s sessionMap) Get(id string) (*User, bool) {
	if user, ok := s[id]; ok {
		return user, true
	}
	return nil, false
}

func (s sessionMap) Store(id string, user *User) {
	s[id] = user
}
