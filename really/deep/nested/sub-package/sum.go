package subpkg

/*
int sum(int, int);
*/
import "C"

func Sum(a, b int) int {
	return int(C.sum(C.int(a), C.int(b)))
}
