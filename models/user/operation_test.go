package user

import (
	"fmt"
	"os"
	"testing"
)

func TestTest(t *testing.T) {
	var list = []string{"a", "b", "c", "d", "e"}
	var res = append(list[:0], list[1:]...)
	fmt.Println(res)
}

func TestInit(t *testing.T) {
	os.MkdirAll(Conf.FilePath, 0777)
}
