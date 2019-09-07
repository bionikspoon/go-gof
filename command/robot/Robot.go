package robot

import (
	"fmt"
	"io"
)

type Robot struct {
	out io.Writer
}

func NewRobot(out io.Writer) *Robot {
	return &Robot{out}
}

func (r *Robot) Move(distance int) {
	if distance > 0 {
		fmt.Fprintf(r.out, "Robot moves forwards %dmm.\n", distance)
	}

	if distance < 0 {
		fmt.Fprintf(r.out, "Robot moves backwards %dmm.\n", -distance)
	}
}

func (r *Robot) Rotate(leftRotation int) {
	if leftRotation > 0 {
		fmt.Fprintf(r.out, "Robot rotates left %d degrees.\n", leftRotation)
	}
	if leftRotation < 0 {
		fmt.Fprintf(r.out, "Robot rotates right %d degrees.\n", -leftRotation)
	}
}
func (r *Robot) Scoop(upwards bool) {
	if upwards {
		fmt.Fprint(r.out, "Robot gathers soil in scoop.\n")
	} else {
		fmt.Fprint(r.out, "Robot releases soil from scoop.\n")

	}
}
