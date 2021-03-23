package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func (d *FenBeiTong) DepartmentAdd(org string) {
	req := DepartmentAddReq{
		CompanyId:     d.AppId,
		OrgUnitName:   org,
		ThirdOrgId:    org,
		ThirdParentId: d.AppId,
	}
	d.Post("/open/api/third/departments/add", req)
}

func (d *FenBeiTong) DepartmentQuery(company_id string) []Department {
	req := DepartmentQuery{
		CompanyId: company_id,
	}
	ret := []Department{}
	postret, err := d.Post("/open/api/third/departments/query", req)
	if err != nil {
		logrus.Error(err.Error())
		return ret
	}
	b, err := json.Marshal(postret)
	if err != nil {
		logrus.Error(err.Error())
		return ret
	}
	err = json.Unmarshal(b, &ret)
	if err != nil {
		logrus.Error(err.Error())
		return ret
	}
	return ret
}
