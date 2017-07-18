package rtda

import "jvmgo/ch07/rtda/heap"

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
func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}
func (self *Thread) CurrentFrame() *Frame  {
	return self.stack.top()
}
func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

//
////ch05新增
//func (self *Thread) NewFrame(maxLocals,maxStack uint) *Frame  {
//	return NewFrame(self,maxLocals,maxStack)
//}
//ch06
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self,method)
}