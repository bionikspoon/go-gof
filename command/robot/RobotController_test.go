package robot_test

import (
	"bytes"
	"reflect"
	"strings"
	"testing"

	r "github.com/bionikspoon/go-gof/command/robot"
)

func TestRobotController(t *testing.T) {
	buffer := &bytes.Buffer{}
	robot := r.NewRobot(buffer)
	controller := r.NewRobotController(buffer)

	move := r.MoveCommand{Robot: robot, ForwardDistance: 1000}
	controller.Commands.Enqueue(move)

	rotate := r.RotateCommand{Robot: robot, LeftRotation: 45}
	controller.Commands.Enqueue(rotate)

	scoop := r.ScoopCommand{Robot: robot, ScoopUpwards: true}
	controller.Commands.Enqueue(scoop)

	controller.ExecuteCommands()
	controller.UndoCommands(3)

	got := strings.Split(buffer.String(), "\n")
	want := []string{
		"EXECUTING COMMANDS.",
		"Robot moves forwards 1000mm.",
		"Robot rotates left 45 degrees.",
		"Robot gathers soil in scoop.",
		"REVERSING 3 COMMAND(S).",
		"Robot releases soil from scoop.",
		"Robot rotates right 45 degrees.",
		"Robot moves backwards 1000mm.",
		"",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot  %v\nwant %v", got, want)
	}

}
