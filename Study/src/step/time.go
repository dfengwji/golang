// time
package step

import (
	"fmt"
	"time"
)

func StudyTime() {
	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
}
