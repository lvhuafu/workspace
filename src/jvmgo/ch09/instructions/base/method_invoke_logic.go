package base

import (
	"jvmgo/ch09/rtda"
	"jvmgo/ch09/rtda/heap"
)
//4条方法调用指令的相同逻辑：给方法创建新的帧并把它推入java虚拟机栈顶


func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method)  {
	thread := invokerFrame.Thread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotSlot := int(method.ArgSlotCount())
	if argSlotSlot > 0{
		for i := argSlotSlot-1;i>=0;i--{
			slot := invokerFrame.OperandStack().PopSlot()
			newFrame.LocalVars().SetSlot(uint(i),slot)
		}
	}
}
