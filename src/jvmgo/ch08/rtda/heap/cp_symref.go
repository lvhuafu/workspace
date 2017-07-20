package heap
//4种类型的符号引用的共同部分

type SymRef struct {
	cp *ConstantPool
	className string
	class *Class
}
//类符号引用解析
func (self *SymRef) ResolvedClass()*Class  {
	if self.class==nil{
		self.resolveClassRef()
	}
	return self.class
}
func (self *SymRef) resolveClassRef()  {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d){
		panic("java.lang.IllegaAccessError(cp_symref)")
	}
	self.class = c
}
