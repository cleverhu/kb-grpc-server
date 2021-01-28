package services

import (
	"context"
	"encoding/json"
	"fmt"
	"nuxt-grpc-server/common"
	"nuxt-grpc-server/models/DocGrpModel"
	"nuxt-grpc-server/models/DocModel"
)

type KbInfoService struct {
}

func (this *KbInfoService) UpdateKbDetailList(ctx context.Context, req *KbInfoRequest) (*KbInfoResponse, error) {
	for _, v := range req.Id {
		detail := GetKbDetail(v)
		data, err := json.Marshal(detail)
		if err != nil {
			return &KbInfoResponse{
				Result: false,
			}, err
		}
		fmt.Println(string(data))
		common.Rds.Set("kb:"+fmt.Sprintf("%d", v), string(data), -1)
	}
	return &KbInfoResponse{
		Result:               true,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}

func GetKbDetail(kbID int64) []*DocGrpModel.DocGrpImpl {
	var dgm []*DocGrpModel.DocGrpImpl

	kbName := &struct {
		Name string `gorm:"column:kb_name"`
	}{}

	common.Orm.Table("kbs").Raw("select kb_name from kbs where kb_id = ? ", kbID).First(&kbName)
	if kbName.Name != "" {
		getKbDetail(kbName.Name, kbID, 0, &dgm)
	}

	return dgm
}

func getKbDetail(kbName string, kbID, groupID int64, result *[]*DocGrpModel.DocGrpImpl) []*DocGrpModel.DocGrpImpl {
	//先找到分组
	common.Orm.Table("doc_grps").Raw(`select group_id,group_name,shorturl from doc_grps 
where kb_id = ? and pid = ? 
order by group_order`, kbID, groupID).Find(&result)

	//遍历分组找文档
	for _, v := range *result {
		//找到文档
		var docs []*DocModel.DocImpl
		common.Orm.Table("docs").Raw(`select doc_id,doc_title,shorturl from docs 
where kb_id = ? and group_id = ? 
order by doc_id`, kbID, v.GroupID).Find(&docs)

		for _, doc := range docs {
			//给文档追加到子元素
			doc.DocHref = "/" + kbName + "/" + v.GroupShortUrl + "/" + doc.DocShortUrl
			v.Children = append(v.Children, doc)
		}

		//寻找子分组 这个数据是临时的，不会返回真实数据
		var grp []*DocGrpModel.DocGrpImpl
		getKbDetail(kbName, kbID, v.GroupID, &grp)
		for _, item := range grp {
			v.Children = append(v.Children, item)
		}
	}
	return *result
}
