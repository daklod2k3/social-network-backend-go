package entity

import (
	"fmt"
	"shared/utils"
)

type SupabaseError struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"error_code"`
	Msg       string `json:"msg"`
}

func Error(err string) SupabaseError {
	var parse SupabaseError
	e := utils.Deserialize(err, &parse)
	if e != nil {
		// supabase error convert
		fmt.Println(e)
	}
	return parse
}
