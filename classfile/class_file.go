package classfile

import "fmt"

type ClassFile struct {
	// magic				uint32
	minorVersion 			uint16
	majorVersion			uint16
	constantPool			ConstantPool
	accessFlags				uint16
	thisClass				uint16
	superClass				uint16
	interfaces				[]uint16
	fields					[]*MemberInfo
	methods					[]*MemberInfo
	attributes				[]AttributeInfo
}

/**
	Parse() 函数把[]byte解析成ClassFile结构体
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	} ()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)

	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)

	self.readAndCheckVersion(reader)

	self.constantPool = readConstantPool(reader)

	self.accessFlags = reader.readUint16()

	self.thisClass = reader.readUint16()

	self.superClass = reader.readUint16()

	self.interfaces = reader.readUint16s()

	self.fields = readMembers(reader, self.constantPool)

	self.methods = readMembers(reader, self.constantPool)

	self.attributes = readAttributes(reader, self.constantPool)
}

/**
	魔术
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
	版本号
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()

	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	// todo
}

/**
	类访问标志
 */
func (self *ClassFile) AccessFlag() uint16 {
	// todo
}

func (self *ClassFile) Fields() []*MemberInfo {
	// todo
}

func (self *ClassFile) Methods() []*MemberInfo {
	// todo
}

/**
	类索引
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/**
	超类索引
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有 java.lang.Object 没有超类
}

/**
	接口索引表
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
