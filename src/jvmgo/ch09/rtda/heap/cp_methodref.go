package heap

import "jvmgo/ch09/classfile"
//方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}


func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}
//解析符号引用
//类d想通过方法符号引用访问类c的某个方法
func (self *MethodRef) resolveMethodRef() {
	d := self.cp.class
	c := self.ResolvedClass()
	if c.IsInterface() {//c是接口
		panic("java.lang.IncompatiableClassChangeError")
	}

	method := lookupMethod(c,self.name,self.descriptor)
	if method==nil{
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d){//是否有权限
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class,name,descriptor)
	if method == nil{
		method = lookupMethodInInterfaces(class.interfaces,name,descriptor)
	}
	return method
}

