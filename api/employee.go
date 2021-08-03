package api

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

/**
查询公司全量人员信息
@see http://docs.open.fenbeitong.com/open-api/2apijie-ru/21zu-zhi-jia-gou/2226-gen-ju-gong-si-id-cha-xun-gong-si-quan-liang-ren-yuan-xin-xi.html
*/
func (d *FenBeiTong) ThridInfo() (employeeData ThirdEmployeeListData, err error) {
	req := map[string]string{
		"companyId": d.AppId,
		"type":      "1",
	}
	employeeData = ThirdEmployeeListData{}

	data, err := d.Post("/open/api/third/company/third/info", req)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &employeeData)
	return
}

/**
查询第三方员工信息
@see http://docs.open.fenbeitong.com/open-api/2apijie-ru/21zu-zhi-jia-gou/218-cha-xun-yuan-gong-xin-606f28-bu-zhi-chi-pi-liang-cha-8be229.html
*/
func (d *FenBeiTong) EmployeesInfo(userCode string) (ret interface{}, err error) {
	req := map[string]string{
		"employee_id": userCode,
		"type":        "1",
	}
	return d.Post("/open/api/third/company/third/info", req)
}

/**
删除第三方员工
@see http://docs.open.fenbeitong.com/open-api/2apijie-ru/21zu-zhi-jia-gou/229-shan-chu-yuan-gong.html
*/
func (d *FenBeiTong) EmployeesDelete(userCode []string) (ret interface{}, err error) {
	req := map[string][]string{
		"third_employee_ids": userCode,
	}
	return d.Post("/open/api/third/employees/delete", req)
}

func (d *FenBeiTong) EmployeesUpdate(employee_list EmployeeReq) (ret interface{}, err error) {
	return d.Post("/open/api/third/employees/v2/update", employee_list)
}

func (d *FenBeiTong) EmployeesSave(employee_list EmployeeReq) (ret interface{}, err error) {

	return d.Post("/open/api/auth/third/user/batch/org_save", employee_list)
}
