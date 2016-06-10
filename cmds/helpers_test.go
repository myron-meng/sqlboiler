package cmds

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/nullbio/sqlboiler/dbdrivers"
)

func TestCombineTypeImports(t *testing.T) {
	imports1 := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
		},
		thirdparty: importList{
			`"github.com/nullbio/sqlboiler/boil"`,
		},
	}

	importsExpected := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
			`"time"`,
		},
		thirdparty: importList{
			`"github.com/nullbio/sqlboiler/boil"`,
			`"gopkg.in/nullbio/null.v4"`,
		},
	}

	cols := []dbdrivers.Column{
		dbdrivers.Column{
			Type: "null.Time",
		},
		dbdrivers.Column{
			Type: "null.Time",
		},
		dbdrivers.Column{
			Type: "time.Time",
		},
		dbdrivers.Column{
			Type: "null.Float",
		},
	}

	res1 := combineTypeImports(imports1, sqlBoilerTypeImports, cols)

	if !reflect.DeepEqual(res1, importsExpected) {
		t.Errorf("Expected res1 to match importsExpected, got:\n\n%#v\n", res1)
	}

	imports2 := imports{
		standard: importList{
			`"errors"`,
			`"fmt"`,
			`"time"`,
		},
		thirdparty: importList{
			`"github.com/nullbio/sqlboiler/boil"`,
			`"gopkg.in/nullbio/null.v4"`,
		},
	}

	res2 := combineTypeImports(imports2, sqlBoilerTypeImports, cols)

	if !reflect.DeepEqual(res2, importsExpected) {
		t.Errorf("Expected res2 to match importsExpected, got:\n\n%#v\n", res1)
	}
}

func TestCombineImports(t *testing.T) {
	t.Parallel()

	a := imports{
		standard:   importList{"fmt"},
		thirdparty: importList{"github.com/nullbio/sqlboiler", "gopkg.in/nullbio/null.v4"},
	}
	b := imports{
		standard:   importList{"os"},
		thirdparty: importList{"github.com/nullbio/sqlboiler"},
	}

	c := combineImports(a, b)

	if c.standard[0] != "fmt" && c.standard[1] != "os" {
		t.Errorf("Wanted: fmt, os got: %#v", c.standard)
	}
	if c.thirdparty[0] != "github.com/nullbio/sqlboiler" && c.thirdparty[1] != "gopkg.in/nullbio/null.v4" {
		t.Errorf("Wanted: github.com/nullbio/sqlboiler, gopkg.in/nullbio/null.v4 got: %#v", c.thirdparty)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	t.Parallel()

	hasDups := func(possible []string) error {
		for i := 0; i < len(possible)-1; i++ {
			for j := i + 1; j < len(possible); j++ {
				if possible[i] == possible[j] {
					return fmt.Errorf("found duplicate: %s [%d] [%d]", possible[i], i, j)
				}
			}
		}

		return nil
	}

	if len(removeDuplicates([]string{})) != 0 {
		t.Error("It should have returned an empty slice")
	}

	oneItem := []string{"patrick"}
	slice := removeDuplicates(oneItem)
	if ln := len(slice); ln != 1 {
		t.Error("Length was wrong:", ln)
	} else if oneItem[0] != slice[0] {
		t.Errorf("Slices differ: %#v %#v", oneItem, slice)
	}

	slice = removeDuplicates([]string{"hello", "patrick", "hello"})
	if ln := len(slice); ln != 2 {
		t.Error("Length was wrong:", ln)
	}
	if err := hasDups(slice); err != nil {
		t.Error(err)
	}

	slice = removeDuplicates([]string{"five", "patrick", "hello", "hello", "patrick", "hello", "hello"})
	if ln := len(slice); ln != 3 {
		t.Error("Length was wrong:", ln)
	}
	if err := hasDups(slice); err != nil {
		t.Error(err)
	}
}

func TestCombineStringSlices(t *testing.T) {
	t.Parallel()

	var a, b []string
	slice := combineStringSlices(a, b)
	if ln := len(slice); ln != 0 {
		t.Error("Len was wrong:", ln)
	}

	a = []string{"1", "2"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] {
		t.Errorf("Slice mismatch: %#v %#v", a, slice)
	}

	b = a
	a = nil
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 2 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != b[0] || slice[1] != b[1] {
		t.Errorf("Slice mismatch: %#v %#v", b, slice)
	}

	a = b
	b = []string{"3", "4"}
	slice = combineStringSlices(a, b)
	if ln := len(slice); ln != 4 {
		t.Error("Len was wrong:", ln)
	} else if slice[0] != a[0] || slice[1] != a[1] || slice[2] != b[0] || slice[3] != b[1] {
		t.Errorf("Slice mismatch: %#v + %#v != #%v", a, b, slice)
	}
}