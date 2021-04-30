package matchfile_test

import (
	"fmt"
	"log"

	"github.com/suzuki-shunsuke/matchfile-parser/matchfile"
)

func Example() {
	parser := matchfile.NewParser()
	checkedFiles := []string{"services/foo.txt"}
	rawConditions := []string{"glob services/*"}
	conditions, err := parser.ParseConditions(rawConditions)
	if err != nil {
		log.Fatal(err)
	}
	f, err := parser.Match(checkedFiles, conditions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	// Output:
	// true
}
