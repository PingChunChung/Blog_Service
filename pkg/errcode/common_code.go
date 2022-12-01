package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(1000000, "服務內部錯誤")
	InvalidParams             = NewError(1000000, "導入參數錯誤")
	NotFound                  = NewError(1000000, "找不到")
	UnauthorizedAuthNotExist  = NewError(1000000, "驗證失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1000000, "驗證失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(1000000, "驗證失敗，Token逾時")
	UnauthorizedTokenGenerate = NewError(1000000, "驗證失敗，Token產生失敗")
	TooManyRequests           = NewError(1000000, "請求過多")
)
