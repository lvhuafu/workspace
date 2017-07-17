package heap
//简化的类加载器
import (
	"jvmgo/ch06/classpath"
	"fmt"
	"jvmgo/ch06/classfile"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp: cp,
		classMap: make(map[string]*Class),
	}
}
//加载类数据到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	if class,ok := self.classMap[name];ok{
		return class//已加载
	}
	return self.loadNonArrayClass(name)
}
//加载类
func (self *ClassLoader) loadNonArrayClass(name string) *Class  {
	data,entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n",name,entry)
	return class
}
//第一步
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data,entry,err :=self.cp.ReadClass(name)
	if err !=nil{
		panic("java.lang.ClassNotFoundException(class_loader): "+name)
	}
	return data,entry
}
//第二步
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}
//class文件数据转换成Class结构体
func parseClass(data []byte) *Class {
	cf,err := classfile.Parse(data)
	if err!=nil{
		panic("java.lang.classFormatError(class_loader)")
	}
	return newClass(cf)
}
//
func resolveSuperClass(class *Class)  {
	if class.name != "java/lang/Object"{
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class)  {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount >0{
		class.interfaces = make([]*Class,interfaceCount)
		for i,interfaceName := range class.interfaceNames{
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}
//第三步 链接
func link(class *Class)  {
	verify(class)
	prepare(class)
}
func verify(class *Class)  {
	//todo java虚拟机规范4.10节
}

func prepare(class *Class)  {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func calcInstanceFieldSlotIds(class *Class)  {
	slotId := uint(0)
	if class.superClass !=nil{
		slotId = class.superClass.instanceSlotCount
	}
	for _,field := range class.fields{
		if !field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class)  {
	slotId := uint(0)
	for _,field := range class.fields{
		if field.IsStatic(){
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble(){
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class)  {
	class.staticVars = newSlots(class.staticSlotCount)
	for _,field := range class.fields{
		if field.IsStatic() && field.IsFinal(){
			initStaticFinalVar(class,field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todoAtChater8")
		}
	}
}
