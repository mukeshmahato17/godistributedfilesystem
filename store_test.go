package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestTransformFunc(t *testing.T) {
	key := "mypic"
	pathname := CASPathTransformFunc(key)
	fmt.Println(pathname)
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpg bytes"))
	if err := s.writeStream("mySpecialPicture", data); err != nil {
		t.Error(err)
	}
}
