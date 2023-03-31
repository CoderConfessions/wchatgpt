package handler

const (
	Ok                 int = 0
	NetworkError       int = 1000
	UnmarshalJsonError int = 1001
	ParamaterError     int = 1100
	NoAuthError        int = 1101
	OpenAIError        int = 2000
)

type UniversalResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,emptyomit"`
}

func (ur *UniversalResp) QuickSetup(code int, message string) UniversalResp {
	ur.Code = code
	ur.Message = message
	return *ur
}
