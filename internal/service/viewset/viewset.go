package viewset

import (
	"github.com/gin-gonic/gin"
	"go-sso/internal/apierror"
	"net/http"
)

type ViewSet struct {
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// handle api error
func (v ViewSet) ErrorHandler(f func(c *gin.Context) error, c *gin.Context) {
	err := f(c)
	switch err := err.(type) {
	case nil:
	case apierror.ApiError:
		v.ErrorResponse(c, err.(apierror.ApiError))
	default:
		v.FailResponse(c, apierror.WrapError(err))
	}
}

// deal error by error code
func (v ViewSet) ErrorResponse(c *gin.Context, e apierror.ApiError) {
	switch e.Code() {
	case apierror.NotFoundCode:
		v.NotFoundResponse(c)
	default:
		v.FailResponse(c, e)
	}
}

func GetSuccessResponse(data interface{}) Response {
	return Response{
		Code: apierror.SuccessCode,
		Msg:  apierror.SuccessMsg,
		Data: data,
	}
}

func GetFailResponse(err apierror.ApiError, data interface{}) Response {
	return Response{
		Code: err.Code(),
		Msg:  err.GetMsg(),
		Data: data,
	}
}

func (v ViewSet) GetId(c *gin.Context) string {
	if i := c.Param("id"); i != "" {
		return i
	}
	return ""
}

// 封装通用response
// Response 返回的数据

func (v ViewSet) SuccessResponse(c *gin.Context, data interface{}) error {
	c.JSON(http.StatusOK, GetSuccessResponse(data))
	return nil
}

func (v ViewSet) SuccessBlankResponse(c *gin.Context) error {
	c.JSON(http.StatusOK, GetSuccessResponse(nil))
	return nil
}

func (v ViewSet) SuccessListResponse(c *gin.Context, data interface{}, PageNum, PageSize, Total int) error {
	c.JSON(http.StatusOK,
		GetSuccessResponse(map[string]interface{}{
			"page_num":  PageNum,
			"page_size": PageSize,
			"total":     Total,
			"data":      data,
		}),
	)
	return nil
}

func (v ViewSet) FailResponse(c *gin.Context, err apierror.ApiError, data ...interface{}) {
	c.JSON(http.StatusBadRequest, GetFailResponse(err, data))
}

func (v ViewSet) NotFoundResponse(c *gin.Context) {
	c.JSON(http.StatusNotFound, GetFailResponse(apierror.ErrNotFound, nil))
}
