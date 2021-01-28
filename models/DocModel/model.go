package DocModel

import "fmt"

type DocImpl struct {
	DocID       int           `gorm:"column:doc_id;primary_key"  json:"doc_id"`
	DocTitle    string        `gorm:"column:doc_title" json:"label"`
	DocShortUrl string        `gorm:"column:shorturl" json:"url"`
	DocHref     string        `gorm:"-"` //拼接地址字符串
	Children    []interface{} `gorm:"-" json:"children,omitempty"`
}


func (this *DocImpl) String() string {
	return fmt.Sprintf("DocID:%d,DocTitle:%s,DocShortUrl:%s,DocHref:%s", this.DocID, this.DocTitle, this.DocShortUrl, this.DocHref)
}

