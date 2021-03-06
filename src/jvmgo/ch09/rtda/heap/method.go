package heap

import "jvmgo/ch09/classfile"

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
		//ch08
		methods[i] = newMethod(class,cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfmethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfmethod)
	method.copyAttribute(cfmethod)

	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative(){
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1} // return
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0} // areturn
	case 'D':
		self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F':
		self.code = []byte{0xfe, 0xae} // freturn
	case 'J':
		self.code = []byte{0xfe, 0xad} // lreturn
	default:
		self.code = []byte{0xfe, 0xac} // ireturn
	}
}

func (self *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}
//ch09
func (self *Method) calcArgSlotCount(paramTypes []string)  {
	for _,paramType := range paramTypes{
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