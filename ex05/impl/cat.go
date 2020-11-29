package impl

type Cat interface {
	Echo(string) string
}

type CatImpl struct {
	Id int
}

func (cat CatImpl) Echo(message string) string {
	return message
}
