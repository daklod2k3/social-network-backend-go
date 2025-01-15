package entity

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"shared/entity"
	logger2 "shared/logger"
	"shared/utils"
)

type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

type DefaultError struct {
	Code int
	Msg  string
}

var (
	logger = logger2.GetLogger()
)

func ParseError(err error) *DefaultError {
	var supabaseError SupabaseError
	var grpcError *status.Status

	logger.Error(err.Error())

	jsonErr := utils.Deserialize(err.Error(), &supabaseError)
	if jsonErr != nil {
		logger.Error(jsonErr.Error())
	}

	switch {

	case jsonErr == nil:
		fmt.Println("supabase error")
		return &DefaultError{supabaseError.Code, supabaseError.Msg}

	case status.Code(err) == codes.Aborted:
		grpcError = status.Convert(err)
		fmt.Println(grpcError.Message())
		ParseError(errors.New(grpcError.Message()))
	}
	return &DefaultError{Code: 500, Msg: "Unknown Error"}
}

func (e *DefaultError) WriteError(c *gin.Context) {
	entity.ResponseJson{Status: e.Code, Message: e.Msg}.WriteError(c)
}

func (err SupabaseError) Error() string {
	//byte, _ := json.Marshal(err)
	return fmt.Sprint(err)
}
