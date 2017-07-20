package rtda

import "jvmgo/ch08/rtda/heap"

type Frame struct {
	lower *Frame
	localVars LocalVars//局部变量表
	operandStack *OperandStack//操作数栈
	//ch05 新增
	thread *Thread
	nextPC int

	//ch06
	method       *heap.Method
}
//ch05 改变
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

//func NewFrame(maxLocals, maxStack uint) *Frame {
//	return &Frame{
//		localVars: newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
func (self *Frame) Thread() *Thread {
	return self.thread
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(nextPC int) {
	self.nextPC = nextPC
}
func (self *Frame) Method() *heap.Method {
	return self.method
}
//ch07
func (self *Frame) RevertNextPC()  {
	self.nextPC = self.thread.pc
}