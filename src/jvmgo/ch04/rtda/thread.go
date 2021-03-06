package rtda

type Thread struct {
	pc int//程序计数器
	stack *Stack//java虚拟机栈指针
}

func NewThread() *Thread  {
	return &Thread{
		stack:newStack(1024),
	}
}
/**
getter
 */
func (self *Thread) PC() int  {
	return self.pc
}
/**
setter
 */
func (self *Thread) SetPC(pc int)  {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame)  {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame  {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame  {
	return self.stack.top()
}