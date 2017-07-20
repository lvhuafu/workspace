package references

import "jvmgo/ch08/instructions/base"
import (
	"jvmgo/ch08/rtda"
	"jvmgo/ch08/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()//得到当前类
	cp := currentClass.ConstantPool()//得到当前常量池
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)//得到方法符号引用
	resolvedClass := methodRef.ResolvedClass()//解析符号引用类
	resolvedMethod := methodRef.ResolvedMethod()//解析符号引用方法

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() !=resolvedClass{//如果是构造方法，这声明类必须是解析类
		panic("java.lang.NoSuchMethodError(invokedpical.go)")
	}
	if resolvedMethod.IsStatic(){//调实例方法，静态方法错误
		panic("java.lang.IncompatibleClassChangeError(invokespical.go)")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)//弹出this引用
	if ref==nil{
		panic("java.lang.NullPointException")
	}

	if resolvedMethod.IsProtected()&&//确保protected方法只能被声明该方法的类或子类调用
		resolvedMethod.Class().IsSuperClassOf(currentClass)&&
		resolvedMethod.Class().GetPackageName() !=currentClass.GetPackageName()&&
		ref.Class() !=currentClass&&
		!ref.Class().IsSubClassOf(currentClass){
		panic("java.lang.IllegalAccessError")
	}

	methodToInvoked := resolvedMethod
	if currentClass.IsSuper()&&
	resolvedClass.IsSuperClassOf(currentClass)&&
	resolvedMethod.Name()!="<init>"{
		methodToInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),methodRef.Name(),methodRef.Descriptor())
	}
	if methodToInvoked == nil || methodToInvoked.IsAbstract(){
		panic("java.lang.AbstractMethodError")
	}

	base.InvokeMethod(frame,methodToInvoked)
}
