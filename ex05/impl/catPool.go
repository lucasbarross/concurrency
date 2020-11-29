package impl

type CatPool struct {
	CatChannel (chan CatImpl)
}

func NewCatPool(size int) CatPool {
	pool := CatPool{make(chan CatImpl, size)}

	for i := 0; i < size; i++ {
		pool.CatChannel <- CatImpl{Id: i}
	}

	return pool
}

func (pool CatPool) Get() CatImpl {
	return <-pool.CatChannel
}

func (pool CatPool) Add(obj CatImpl) {
	pool.CatChannel <- obj
}
