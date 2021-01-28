package DocModel

type DocModelAttrFunc func(d *DocImpl)

type DocModelAttrFuncs []DocModelAttrFunc

func (this DocModelAttrFuncs) Apply(d *DocImpl) {
	for _, f := range this {
		f(d)
	}
}
