package service

//校验并初始化分页参数
func CheckPageReq(pageNo, pageSize, pageFrom int32) (int32, int32, int32) {
	if pageNo == 0 {
		pageNo = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	pageFrom = (pageNo - 1) * pageSize

	return pageNo, pageSize, pageFrom
}