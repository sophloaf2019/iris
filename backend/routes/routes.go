package routes

type Response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func ErrorResponse(data error) Response {
	return Response{
		Success: false,
		Data:    data.Error(),
	}
}

func SuccessResponse(data any) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

func SmartResponse(data any, err error) (int, Response) {
	if err != nil {
		return 500, ErrorResponse(err)
	}
	return 200, SuccessResponse(data)
}
