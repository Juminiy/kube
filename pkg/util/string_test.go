package util

import (
	"testing"
	"time"
)

// +passed
func TestBytes2StringNoCopy(t *testing.T) {
	bytesOf := []byte{104, 98, 111}
	t.Log(string(bytesOf))
	t.Log(Bytes2StringNoCopy(bytesOf))
}

// +passed
func TestString2BytesNoCopy(t *testing.T) {
	strOf := "Alan"
	t.Log([]byte(strOf))
	t.Log(String2BytesNoCopy(strOf))
}

func TestStringCast(t *testing.T) {
	t.Log(StringCast([]int{1, 2, 3}))
	t.Log(StringCast([]string{"1", "2", "3"}))
	t.Log(StringCast([]time.Time{time.Now(), time.Now().AddDate(1, 0, 0), time.Now().AddDate(0, 1, 0)}))
	t.Log(StringCast([]*int{nil, nil, nil}))
	t.Log(StringCast([]chan int{nil, nil, nil}))
	t.Log(StringCast([]func(){nil, nil, nil}))
	t.Log(StringCast([]interface{}{nil, nil, nil}))
	t.Log(StringCast([]map[string]int{nil, nil, nil}))
	t.Log(StringCast([][]int{nil, nil, nil}))
}
