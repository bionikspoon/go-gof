package robot

import (
	"fmt"
	"io"
)

type RobotController struct {
	out       io.Writer
	Commands  *commandQueue
	UndoStack *commandStack
}

func NewRobotController(out io.Writer) *RobotController {
	queue := &commandQueue{}
	stack := &commandStack{}
	return &RobotController{out, queue, stack}
}

func (controller *RobotController) ExecuteCommands() {
	fmt.Fprint(controller.out, "EXECUTING COMMANDS.\n")

	for {
		if controller.Commands.Count <= 0 {
			break
		}
		command := controller.Commands.Dequeue()

		command.Up()

		controller.UndoStack.Push(command)
	}
}

func (controller *RobotController) UndoCommands(n int) {
	fmt.Fprintf(controller.out, "REVERSING %d COMMAND(S).\n", n)

	for {
		if controller.UndoStack.Count <= 0 || n <= 0 {
			break
		}

		command := controller.UndoStack.Pop()
		command.Down()
		n--
	}
}

type commandQueue struct {
	commands []RobotCommand
	Count    int
}

func (c *commandQueue) Enqueue(command RobotCommand) {
	c.commands = append(c.commands, command)
	c.Count++
}

func (c *commandQueue) Dequeue() (command RobotCommand) {
	command, c.commands = c.commands[0], c.commands[1:]
	c.Count--
	return
}

type commandStack struct {
	commands []RobotCommand
	Count    int
}

func (c *commandStack) Push(command RobotCommand) {
	c.commands = append([]RobotCommand{command}, c.commands...)

	c.Count++
}

func (c *commandStack) Pop() (command RobotCommand) {
	command, c.commands = c.commands[0], c.commands[1:]
	c.Count--
	return
}
