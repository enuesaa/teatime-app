package controller

type ApiResponse struct {
	Data interface{} `json:"data"`
}

type IdSchema struct {
	Id string `json:"id"`
}
func NewIdSchema(id string) IdSchema {
	return IdSchema{
		Id: id,
	}
}
type EmptySchema struct{}
func NewEmptySchema() EmptySchema {
	return EmptySchema{}
}