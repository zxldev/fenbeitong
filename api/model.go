package api

type BaseParams struct {
	Timestamp   int64  `json:"timestamp"`
	AccessToken string `json:"access_token"`
	EmployeeId  string `json:"employee_id"`
}

type BaseResponse struct {
	RequestId string      `json:"request_id"`
	Code      int64       `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

type AuthorizeResp struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      string `json:"data"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type DepartmentQuery struct {
	CompanyId string `json:"company_id"`
}

type Department struct {

	//"employeeCount": 0,
	//"orgUnitFullName":"北京分贝通科技有限公司/财务部/财务一部",
	//"orgUnitId": "59e9bcc4279863398a0672eb",
	//"orgThirdUnitId": "",
	//"orgThirdUnitParentId": "",
	//"orgUnitName": "财务一部",
	//"orgUnitParentFullName": "北京分贝通科技有限公司/财务部",
	//"orgUnitParentId": "5abdf19e27986373488a010f",
	//"orgUnitParentName": "财务部"
	EmployeeCount         int    `json:"employeeCount"`
	OrgUnitFullName       string `json:"orgUnitFullName"`
	OrgUnitId             string `json:"orgUnitId"`
	OrgThirdUnitId        string `json:"orgThirdUnitId"`
	OrgThirdUnitParentId  string `json:"orgThirdUnitParentId"`
	OrgUnitName           string `json:"orgUnitName"`
	OrgUnitParentFullName string `json:"orgUnitParentFullName"`
	OrgUnitParentId       string `json:"orgUnitParentId"`
	OrgUnitParentName     string `json:"orgUnitParentName"`
}

type DepartmentAddReq struct {
	CompanyId     string `json:"company_id"`
	OrgUnitName   string `json:"org_unit_name"`
	ThirdParentId string `json:"third_parent_id"`
	ThirdOrgId    string `json:"third_org_id"`
	OperatorId    string `json:"operator_id,omitempty"`
}

type Employee struct {
	Name            string `json:"name"`
	Role            int64  `json:"role"`
	Phone           string `json:"phone"`
	Email           string `json:"email"`
	OrgUnitName     string `json:"org_unit_name"`
	ThirdEmployeeId string `json:"third_employee_id"`
}

type EmployeeReq struct {
	EmployeeList []Employee `json:"employee_list"`
}

type ThirdEmployeeListData struct {
	ThirdEmployeeList []ThirdEmployee `json:"thirdEmployeeList"`
}

type ThirdEmployee struct {
	Manager         bool   `json:"manager"`
	Gender          int    `json:"gender"`
	Phone           string `json:"phone"`
	Name            string `json:"name"`
	Id              string `json:"id"`
	DeptId          string `json:"deptId"`
	ThirdEmployeeId string `json:"thirdEmployeeId"`
	Dept            string `json:"dept"`
	Status          int    `json:"status"`
}
