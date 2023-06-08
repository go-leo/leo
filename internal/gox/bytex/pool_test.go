package bytex

//
// func TestPool(t *testing.T) {
// 	defer debug.SetGCPercent(debug.SetGCPercent(-1))
// 	var p Pool
//
// 	buf := p.Get()
// 	t.Logf("%p", buf)
// 	p.Put(buf)
// 	buf = nil
//
// 	newBuf := p.Get()
// 	t.Logf("%p", newBuf)
// 	newBuf.Free()
// 	newBuf = nil
//
// 	runtime.GC()
//
// 	lastBuf := p.Get()
// 	t.Logf("%p", lastBuf)
// 	p.Put(lastBuf)
// }
