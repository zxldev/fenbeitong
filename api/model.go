package api


type BaseParams struct {
	Timestamp    int64 `json:"timestamp"`
	AccessToken string `json:"access_token"`
	EmployeeId   string `json:"employee_id"`
}


type BaseResponse struct {
	RequestId string `json:"request_id"`
	Code int64 `json:"code"`
	Msg string `json:"msg"`
}

type AuthorizeResp struct {
	RequestId string `json:"request_id"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data Token `json:"data"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

type DepartmentAddReq struct {
	CompanyId string `json:"company_id"`
	OrgUnitName string `json:"org_unit_name"`
	ThirdParentId string `json:"third_parent_id"`
	ThirdOrgId string `json:"third_org_id"`
	OperatorId string `json:"operator_id,omitempty"`
}

type Employee struct {
	Name string `json:"name"`
	Role int64 `json:"role"`
	Phone  string `json:"phone"`
	Email string `json:"email"`
	OrgUnitName string `json:"org_unit_name"`
	ThirdEmployeeId string `json:"third_employee_id"`
}


type EmployeeReq struct {
	EmployeeList []Employee `json:"employee_list"`
}