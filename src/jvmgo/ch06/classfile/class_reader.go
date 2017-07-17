package classfile

import "encoding/binary"

type ClassReader struct {
	data [] byte
}
//u1
func (self *ClassReader) readUint8() uint8  {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

//u2
func (self *ClassReader) readUint16() uint16  {
	val := binary.BigEndian.Uint16(self.data)//class文件就是以大端模式在存储
	self.data = self.data[2:]
	return val
}

//u4
func (self *ClassReader) readUint32() uint32  {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64  {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

//读取unit16表，表的大小由开头的unit16数据指出
func (self *ClassReader) readUint16s() []uint16  {
	n := self.readUint16()
	s := make([]uint16,n)
	for i := range s {
		s[i] = self.readUint16()//调用前面的函数
	}
	return s
}

func (self *ClassReader) readBytes(length uint32) []byte  {//这里有个问题如果没有length这么长的数据了怎么办
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}