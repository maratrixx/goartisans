package convert

import "testing"

func TestBytesToString(t *testing.T) {
	s := "hello世界"
	b := []byte(s)

	if BytesToString(b) != s {
		t.FailNow()
	}
}

func TestStringToBytes(t *testing.T) {
	s := "hello世界"
	b := []byte(s)

	r := StringToBytes(s)

	if len(r) != len(b) {
		t.FailNow()
	}

	for k, v := range r {
		if v != b[k] {
			t.FailNow()
		}
	}
}

func BenchmarkBytesToString(b *testing.B) {
	bb := []byte("hello世界")
	for i := 0; i < b.N; i++ {
		_ = BytesToString(bb)
	}
}

func BenchmarkBytesToString2(b *testing.B) {
	bb := []byte("hello世界")
	for i := 0; i < b.N; i++ {
		_ = string(bb)
	}
}

func BenchmarkStringToBytes(b *testing.B) {
	s := "hello世界"
	for i := 0; i < b.N; i++ {
		_ = StringToBytes(s)
	}
}

func BenchmarkStringToBytes2(b *testing.B) {
	s := "hello世界"
	for i := 0; i < b.N; i++ {
		_ = []byte(s)
	}
}
