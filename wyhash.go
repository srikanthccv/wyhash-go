package wyhash

// Wyhash64 is an implementation of the random number generator wyhash
type Wyhash64 struct {
	seed uint64
}

func (w *Wyhash64) next() uint64 {
	w.seed += 0x60bee2bee120fc15
	t1, _ := FromString("a3b195354a39b70d")
	t1 = From64(w.seed).Mul(t1)
	m1 := t1.hi ^ t1.lo

	t2, _ := FromString("1b03738712fad5c9")
	t2 = From64(m1).Mul(t2)
	m2 := t2.hi ^ t1.lo
	return m2
}
