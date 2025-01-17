package globals

type AppCode int

const (
	CodeSuccess      AppCode = 200
	CodeFailed       AppCode = 402
	CodeHttpBad      AppCode = 502
	CodeSercerError  AppCode = 500
	CodeUnauthorized AppCode = 403
)
