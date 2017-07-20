package control

import (
	"jvmgo/ch08/instructions/base"
	"jvmgo/ch08/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}
type ARETURN struct {
	base.NoOperandsInstruction
}
type DRETURN struct {
	base.NoOperandsInstruction
}
type FRETURN struct {
	base.NoOperandsInstruction
}
type IRETURN struct {
	base.NoOperandsInstruction
}
type LRETURN struct {
	base.NoOperandsInstruction
}
//6条return指令，用于返回不同类型的值
func (self *RETURN) Execute(frame *rtda.Frame)  {
	frame.Thread().PopFrame()
}
func (self *ARETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}
func (self *DRETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retVal)
}
func (self *FRETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retVal)
}
func (self *IRETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retVal)
}
func (self *LRETURN) Execute(frame *rtda.Frame)  {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	retVal := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retVal)
}
