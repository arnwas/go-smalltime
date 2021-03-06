package readme_example_test

import "testing"

import "fmt"
import "github.com/kstenerud/go-smalltime"

func demonstrateSmalltime() {
	noonJan12000 := smalltime.New(2000, 1, 1, 12, 0, 0, 0)
	oneOclockJan12000 := smalltime.New(2000, 1, 1, 13, 0, 0, 0)
	feb151999 := smalltime.New(1999, 2, 15, 12, 8, 45, 9122)

	if oneOclockJan12000 > noonJan12000 {
		fmt.Printf("Comparison: One o'clock is greater than noon.\n")
	}

	if feb151999 < noonJan12000 {
		fmt.Printf("Comparison: Feb 15, 1999 is less than Jan 1, 2000.\n")
	}

	gotime := feb151999.AsTime()
	fmt.Printf("Go Time: %v\n", gotime)

	smtime := smalltime.FromTime(gotime)
	fmt.Printf("Smalltime Raw: 0x%016x\n", smtime)

	fmt.Printf("Smalltime Fields: %04d-%02d-%02d %02d:%02d:%02d.%06d\n",
		smtime.Year(), smtime.Month(), smtime.Day(), smtime.Hour(),
		smtime.Minute(), smtime.Second(), smtime.Microsecond())
}

func TestReadmeExamples(t *testing.T) {
	demonstrateSmalltime()
}
