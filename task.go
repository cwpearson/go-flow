package flow

import (
	"io"

	"github.com/labstack/gommon/log"
)

type Task interface {
	Wait()
	Hook(func(interface{}))
	ID() int
}

var goTaskCount int

type GoTask struct {
	Source  string
	Inputs  []string
	Outputs []string

	In  io.ReadCloser
	Out io.Writer

	id int
}

func (t *GoTask) InPipe() io.WriteCloser {
	if t.In != nil {
		log.Error("In already set")
	}
	pr, pw := io.Pipe()
	t.In = pr
	return pw
}

func (t *GoTask) OutPipe() io.ReadCloser {
	if t.Out != nil {
		log.Error("Out already set")
	}
	pr, pw := io.Pipe()
	t.Out = pw
	// t.closeAfterWait = append(e.closeAfterWait, pw)
	return pr
}

// NewGoTaskFromFile loads a task implementation from a file
func NewGoTaskFromFile(path, functionName string) *GoTask {
	return &GoTask{}
}

// NewGoTask produces a new go task
func NewGoTask(source string, inputs, outputs []string) *GoTask {
	g := &GoTask{Source: source, Inputs: inputs, Outputs: outputs, id: goTaskCount}
	goTaskCount++
	return g
}

func (t *GoTask) ID() int {
	return t.id
}

func (t *GoTask) InputInt(i int) *GoTask {
	return t
}

func (t *GoTask) To(other Task) Task {
	Connect(t.ID(), 0, other.ID(), 0)
	return other
}

func (t *GoTask) Merge(preds ...Task) *GoTask {
	for _, p := range preds {
		Connect(p.ID(), 0, t.ID(), 0)
	}
	return t
}

func (t *GoTask) Wait() {

}

func (t *GoTask) Hook(f func(interface{})) {

}
