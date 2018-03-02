package handler

import (
	"../inspector"
	"fmt"
	"os"
)

func Process(conclusion inspector.Conclusion) {
	linebreak := "======================================================================"
	switch result := conclusion.Result; result {
	case inspector.Pass:
		fmt.Println(linebreak)
		fmt.Println("No permission is changed.")
		fmt.Println(linebreak)
		os.Exit(0)
	case inspector.Fail:
		more := conclusion.More
		less := conclusion.Less
		fmt.Println(linebreak)
		fmt.Println("Warning!")
		fmt.Println(fmt.Sprintf("\n%d new permission(s) added:", len(more)))
		for i := range more {
			fmt.Println(fmt.Sprintf("    %s", more[i].Name))
		}
		if len(less) > 0 {
			fmt.Println(fmt.Sprintf("\n%d old permission(s) removed:", len(less)))
			for i := range less {
				fmt.Println(fmt.Sprintf("    %s", less[i].Name))
			}
		}
		fmt.Println(linebreak)
		os.Exit(1)
	case inspector.PassWithAttention:
		less := conclusion.Less
		fmt.Println(linebreak)
		fmt.Println(fmt.Sprintf("Brilliant!  You got %d permission(s) removed:\n", len(less)))
		for i := range less {
			fmt.Println(fmt.Sprintf("    %s", less[i].Name))
		}
		fmt.Println(linebreak)
		os.Exit(0)
	}
}
