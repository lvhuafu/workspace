package native

import "jvmgo/ch09/rtda"
//本地方法注册表
type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}//hash表

func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className+"~"+methodName+"~"+methodDescriptor
	registry[key] = method
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className+"~"+methodName+"~"+methodDescriptor
	if method,ok := registry[key];ok{
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives"{
		return emptyNativeMethod
	}
	return nil
}
func emptyNativeMethod(frame *rtda.Frame)  {
	//donothing
}
