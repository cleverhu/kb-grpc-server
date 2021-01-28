package DocGrpModel

type DocGrpImpl struct {
	GroupID       int64           `gorm:"column:group_id;primary_key" json:"group_id"`
	GroupName     string        `gorm:"column:group_name" json:"label"`
	GroupShortUrl string        `gorm:"column:shorturl" json:"url"`
	Children      []interface{} `gorm:"-" json:"children,omitempty"`
}


func New(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {
	d := &DocGrpImpl{}
	DocGrpModelAttrFuncs(attrs).Apply(d)
	return d
}

func (this *DocGrpImpl) Mutate(attrs ...DocGrpModelAttrFunc) *DocGrpImpl {

	DocGrpModelAttrFuncs(attrs).Apply(this)
	return this
}
