package apierrors

type ErrCode string

const (
	Unknown               ErrCode = "U000"
	InsertDataFailed      ErrCode = "S001"
	GetDataFailed         ErrCode = "S002"
	NAData                ErrCode = "S003"
	UpdateDataFailed      ErrCode = "S005"
	ReqBodyDecodeFailed   ErrCode = "R001"
	BadParam              ErrCode = "R002"
	RequiredAuthorization ErrCode = "A001"
	Unauthorizated        ErrCode = "A003"
	NotMatchUser          ErrCode = "A004"
)
