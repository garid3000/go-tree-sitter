package tsx

//#include "parser.h"
//TSLanguage *tree_sitter_tsx();
import "C"
import (
	"unsafe"

	sitter "github.com/garid3000/go-tree-sitter"
)

func GetLanguage() *sitter.Language {
	ptr := unsafe.Pointer(C.tree_sitter_tsx())
	return sitter.NewLanguage(ptr)
}
