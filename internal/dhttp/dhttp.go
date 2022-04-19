package dhttp

const (
	OK             = "00000"  //一切OK
	CheckFailure   = "A1001" //参数校验失败
	ParaInvaild    = "A1002"
	CreateDirError = "B1001"
	DeleteDirError = "B1002" //删除目录错误
	//  CRUD(Create, Read, Update and Delete)
	DatabaseCError = "D1001" //数据库增加数据错误
	DatabaseDError = "D1002" //数据库删除数据错误
	DatabaseUError = "D1003" //数据库更新数据错误
	DatabaseRError = "D1004" //数据库查询数据错误

)
