package entities

type Token struct {
	uuid   int64
	userId int64
}

func NewToken(
	uuid int64,
	userId int64,
) *Token {
	return &Token{
		uuid:   uuid,
		userId: userId,
	}
}

func (u *Token) Id() int64 {
	return u.uuid
}

func (u *Token) UserId() int64 {
	return u.userId
}
