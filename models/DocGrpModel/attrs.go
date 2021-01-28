package DocGrpModel

type DocGrpModelAttrFunc func(d *DocGrpImpl)

type DocGrpModelAttrFuncs []DocGrpModelAttrFunc

func (this DocGrpModelAttrFuncs) Apply(d *DocGrpImpl) {
	for _, f := range this {
		f(d)
	}
}
