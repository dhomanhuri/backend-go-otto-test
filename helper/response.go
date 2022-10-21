package helper

import "net/http"

func BadRequest(msg string) (int, map[string]interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"code":    http.StatusBadRequest,
		"status":  "error",
		"message": msg,
	}
}

func SuccessInsert() (int, map[string]interface{}) {
	return http.StatusCreated, map[string]interface{}{
		"code":    http.StatusCreated,
		"status":  "success",
		"message": "insert success",
	}
}

func FailedBadRequestWithMSG(err string) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": err,
		"code":    400,
	}
}

func SuccessGetData(data interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    data,
		"code":    200,
	}
}
