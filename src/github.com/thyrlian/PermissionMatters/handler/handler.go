package handler

import (
	"../inspector"
	"fmt"
	"os"
)

func Process(conclusion inspector.Conclusion) {
	delimiter := "======================================================================"
	switch result := conclusion.Result; result {
	case inspector.Pass:
		fmt.Println(delimiter)
		fmt.Println("No permission is changed.")
		fmt.Println(delimiter)
		os.Exit(0)
	case inspector.Fail:
		fmt.Println(delimiter)
		fmt.Println("Fail")
		fmt.Println(delimiter)
		os.Exit(1)
	case inspector.PassWithAttention:
		fmt.Println(delimiter)
		fmt.Println("Brilliant!  You got permissions removed:")
		fmt.Println(delimiter)
		os.Exit(0)
	}
}
