package api


func (d *FenBeiTong)DepartmentAdd(org string){
	req := DepartmentAddReq{
		CompanyId:d.AppId,
		OrgUnitName:org,
		ThirdOrgId:org,
		ThirdParentId:d.AppId,
	}
	d.Post("/open/api/third/departments/add",req)
}