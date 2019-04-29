package classfile

type ClassFile struct {
	// magic				uint32
	minorVersion 			uint16
	magorVersion			uint16
	constantPool			ConstantPool
	accessFlags				uint16
	thisClass				uint16
	superClass				uint16
	interfaces				[]uint16
	fields					[]*MemberInfo
	methods					[]*MemberInfo
	attributes				[]AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	// todo
}

func (self *ClassFile) read(reader *ClassReader) {
	// todo
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {

}
