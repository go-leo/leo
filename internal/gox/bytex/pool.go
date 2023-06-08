package bytex

//
// import (
// 	"bytes"
// 	"sync"
// )
//
// type Buffer struct {
// 	*bytes.Buffer
// 	pool *Pool
// }
//
// func (buf *Buffer) Free() {
// 	buf.pool.Put(buf)
// }
//
// type Pool struct {
// 	sync.Pool
// }
//
// func (pool *Pool) Get() *Buffer {
// 	v := pool.Pool.Get()
// 	if v != nil {
// 		buf := v.(*Buffer)
// 		buf.Reset()
// 		return buf
// 	}
// 	return &Buffer{pool: pool, Buffer: &bytes.Buffer{}}
// }
//
// func (pool *Pool) Put(buf *Buffer) {
// 	pool.Pool.Put(buf)
// }
//
// var defaultPool = new(Pool)
//
// func Get() *Buffer {
// 	return defaultPool.Get()
// }
//
// func Put(buf *Buffer) {
// 	defaultPool.Put(buf)
// }
