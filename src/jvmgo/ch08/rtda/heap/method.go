package heap

import "jvmgo/ch08/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	//ch07
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
		//ch07
		methods[i].calcArgSlotCount()
	}
	return methods
}

func (self *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
//ch07
func (self *Method) calcArgSlotCount()  {
	parsedDescriptor := parseMethodDescriptor(self.descriptor)//分解方法描述符
	for _,paramType := range parsedDescriptor.parameterTypes{
		self.argSlotCount++
		if paramType== "J"||paramType=="D"{//
			self.argSlotCount++
		}
	}
	if !self.IsStatic(){//实例方法有个this参数
		self.argSlotCount++
	}
}
//访问标志检查
func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags & ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags & ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags & ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags & ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags & ACC_STRICT
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}
func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}