package classpath

type Classpath struct {
	bootClasspath		Entry
	extClasspath		Entry
	userClasspath		Entry
}

/**
	Classpath对象实例的创建接口
 */
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return nil
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// todo
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	// todo
}

func (self *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	// todo
	return nil, nil, nil
}

func (self *Classpath) String() string {
	// todo
	return ""
}


