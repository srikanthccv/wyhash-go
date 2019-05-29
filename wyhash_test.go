package wyhash

import "testing"

func TestWyhash(t *testing.T) {
	n := Wyhash64{seed: 0}
	if n.next() != 7962096441729272069 {
		t.Error("dope.")
	}
	if n.next() != 15924192883458544138 {
		t.Error("dope.")
	}
	if n.next() != 5439545251478264591 {
		t.Error("dope.")
	}
}
