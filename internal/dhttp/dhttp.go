package dhttp

const (
	OK             = "00000"  //一切OK
	CheckFailure   = "A10001" //参数校验失败
	CreateDirError = "B10001"
	DeleteDirError = "B10002" //删除目录错误
	//  CRUD(Create, Read, Update and Delete)
	DatabaseCError = "D10001" //数据库增加数据错误
	DatabaseDError = "D10002" //数据库删除数据错误
	DatabaseUError = "D10003" //数据库更新数据错误
	DatabaseRError = "D10004" //数据库查询数据错误

)
