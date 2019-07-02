package subpkg

/*
#cgo

int sum(int, int);
*/
import "C"

func Sum(a, b int) int {
	return C.sum(C.int(a), C.int(b))
}
