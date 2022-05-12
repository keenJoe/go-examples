package core

type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	//没有返回reply，会自己返回吗？
	return nil
}

func (t *Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	//没有返回reply，会自己返回吗？
	return nil
}
