package behavior

const (
	SUM = "+"
	SUB = "-"
)

type Interpreter interface {
	Eval() int
}

type value int

func (v *value) Eval() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

type operationSub struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Eval() int {
	return a.Left.Eval() + a.Right.Eval()
}

func (s *operationSub) Eval() int {
	return s.Left.Eval() - s.Right.Eval()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSub{
			Left:  left,
			Right: right,
		}
	}

	return nil
}

type operationStack []Interpreter

func (p *operationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *operationStack) Pop() Interpreter {
	length := len(*p)
	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}
	return nil
}
