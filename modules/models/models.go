package models

type (
	PaginateReq struct {
		Start string 	`query:"start" validate:"max=64"`
		Limit int 		`query:"limit" validate:"required,min=2,max=10"`
	}
	PaginateRes struct {
		Data any 				`json:"data"`
		Limit int 				`json:"limit"`
		Total int64 			`json:"total"`
		First FirestPaginate 	`json:"first"`
		Next NextPaginate 		`json:"next"`
	}

	FirestPaginate struct {
		Href string		`json:"href"`
		Start string	`json:"start"`

	}

	NextPaginate struct {
		Href string 	`json:"href"`
	}

	KafkaOffset struct {
		Offset int64 `json:"offset" bson:"offent"`
		
	}
)