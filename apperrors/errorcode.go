package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U000" //開発者が想定していないエラーが発生した」ときに使うことを想定したエラーコード
)
