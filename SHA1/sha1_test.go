package SHA1

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_Getsha1(t *testing.T) {
	var tests = []struct {
		in  string
		out [20]byte
	}{
		//{"D:/Work/Project/Go/src/mytest/1.jpg",[58 6 1 49 26 155 176 80 20 94 106 49 169 150 5 204 71 55 46 183]}
		{"D:/Work/Project/Go/src/mytest/Hash/Combination/result.txt",[40 90 221 215 167 242 17 188 28 161 182 19 187 40 53 214 197 94 201 30]}
	    {"",[] }
	}
	for test:=range tests{
		actual:=GetSha1(tests.in)
		if actual!=tests.out {
			t.Errorf("GetSha1(%s)=%v;out %s", tests.in,actual,tests.out)
		}
	}
}
