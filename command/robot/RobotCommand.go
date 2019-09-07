package robot

type RobotCommand interface {
	Up()
	Down()
}

type MoveCommand struct {
	Robot           *Robot
	ForwardDistance int
}

func (c MoveCommand) Up() {
	c.Robot.Move(c.ForwardDistance)
}
func (c MoveCommand) Down() {
	c.Robot.Move(-c.ForwardDistance)
}

type RotateCommand struct {
	Robot        *Robot
	LeftRotation int
}

func (c RotateCommand) Up() {
	c.Robot.Rotate(c.LeftRotation)
}

func (c RotateCommand) Down() {
	c.Robot.Rotate(-c.LeftRotation)
}

type ScoopCommand struct {
	Robot        *Robot
	ScoopUpwards bool
}

func (c ScoopCommand) Up() {
	c.Robot.Scoop(c.ScoopUpwards)
}

func (c ScoopCommand) Down() {
	c.Robot.Scoop(!c.ScoopUpwards)
}
