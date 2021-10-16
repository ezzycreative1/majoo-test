package handler

import (
	"net/http"

	"github.com/ezzycreative1/majoo-test/helpers/response_mapping"

	"github.com/gin-gonic/gin"
)

var statusCode = 200

const MSG = "Failed Validation BackEnd"

type ErrorResponse struct {
	ResponseCode        string `json:"response_code"`
	ResponseDescription string `json:"response_description"`
	ResponseData        string `json:"response_data"`
}
type GenericResponse struct {
	Success   bool        `json:"success"`
	Errorid   int         `json:"errorid"`
	MessageEN string      `json:"error_en"`
	MessageIN string      `json:"error_id"`
	Data      interface{} `json:"data"`
	//Transactionid string `json:"transactionid"`
}

// SetStatusCode -- Set status code
func SetStatusCode(statCode int) int {
	statusCode := statCode
	return statusCode
}

/**
This function take two parameters :
	errorId :	this is transaction status
	formName : 	if you want to use specific success message put form name, if not, you can use "" empty string. (only applied if errorId == 0)
*/
// func NewGenericResponse(errorId int, formName string) *GenericResponse {

// 	messageEn := util.GetErrorString(util.EN, errorId)
// 	messageId := util.GetErrorString(util.ID, errorId)

// 	if formName != "" && errorId == 0 {
// 		messageEn = util.GetSuccessString(util.EN, formName)
// 		messageId = util.GetSuccessString(util.ID, formName)
// 	}

// 	return &GenericResponse{
// 		Success:   errorId == 000,
// 		Errorid:   errorId,
// 		MessageEN: messageEn,
// 		MessageIN: messageId,
// 	}
// }

// RespondJSON -- set response to json format
func RespondJSON(c *gin.Context, data interface{}, request interface{}) {
	c.JSON(statusCode, data)
	return
}

// RespondSuccess ..
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": message,
		"data":    data,
	})
	return
}

type BaseError struct {
	Error error
	Code  string
}

// ResponseFailed ..
func ResponseFailed(c *gin.Context, base BaseError) {
	message, ok := response_mapping.ResponseMappingEN[base.Code]
	if !ok {
		c.JSON(400, gin.H{
			"code":    base.Code,
			"message": "Response code not found!",
			"data":    base.Error.Error(),
		})
	}
	c.JSON(400, gin.H{
		"code":    base.Code,
		"message": message,
		"data":    base.Error.Error(),
	})
	return
}

// ResponseSuccess ..
func ResponseSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": message,
		"data":    data,
	})
	return
}

func FailedAPIValidation(message string, c *gin.Context) {

	var res = ErrorResponse{
		"03",
		MSG,
		message,
	}

	c.JSON(400, res)
}

// FailedResponseBackend ..
func FailedResponseBackend(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    http.StatusBadRequest,
		"message": MSG,
		"data":    err.Error()})
	return
}

func GeneralError(message string, c *gin.Context) {
	var res = ErrorResponse{
		"03",
		MSG,
		message,
	}

	c.JSON(500, res)
}

// RespondCreated -- Set response for create process
func RespondCreated(c *gin.Context, message string, data interface{}, request interface{}) {
	if message == "" {
		message = "Resource Created"
	}
	statusCode := SetStatusCode(201)
	c.JSON(statusCode, data)
	return
}

// RespondUpdated -- Set response for update process
func RespondUpdated(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Updated"
	}
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// RespondDeleted -- Set response for delete process
func RespondDeleted(c *gin.Context, message string) {
	if message == "" {
		message = "Resource Deleted"
	}
	statusCode := SetStatusCode(200)
	c.JSON(statusCode, gin.H{"message": message})
	return
}

// RespondError -- Set response for error
func RespondError(c *gin.Context, message interface{}, statusCode int, request interface{}) {
	data := gin.H{"status": statusCode, "message": message}
	c.JSON(statusCode, data)
	return
}

// RespondFailValidation ..
func RespondFailValidation(c *gin.Context, message interface{}, request interface{}) {
	RespondError(c, message, 422, request)
	return
}

// RespondUnauthorized -- Set response not authorized
func RespondUnauthorized(c *gin.Context, message string) {

	c.JSON(http.StatusUnauthorized, gin.H{
		"response_code":        "401",
		"response_description": "Unauthorized",
		"response_data":        message,
	})
	c.Abort()
	return
}

// RespondNotFound -- Set response not found
func RespondNotFound(c *gin.Context, code string, status string, message string, request interface{}) {
	if message == "" {
		message = "Resource Not Found"
	}
	statusCode := SetStatusCode(404)
	data := gin.H{"code": code, "status": status, "message": message}
	c.JSON(statusCode, data)
	return
}

// RespondMethodNotAllowed -- Set response method not allowed
func RespondMethodNotAllowed(c *gin.Context, message string) {
	if message == "" {
		message = "Method Not Allowed"
	}
	statusCode := SetStatusCode(405)
	c.JSON(statusCode, gin.H{"errors": message})
	return
}

func RespondForbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"response_code":        "403",
		"response_description": "Unauthorized",
		"reponse_data":         "forbidden",
	})
	return
}

// RespondErrorAPI ..
func RespondErrorAPI(c *gin.Context, message string, request interface{}) {
	statusCode := SetStatusCode(400)
	data := gin.H{"status": statusCode, "message": message}
	c.JSON(statusCode, data)
	return
}

func GeneralResponse() {

}

func CustomErrorAPI(c *gin.Context, request interface{}, err error) {
	data := gin.H{"status": 400, "data": request, "message": err.Error()}
	c.JSON(400, data)
	return
}
