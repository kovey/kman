package code

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	Token_Err         codes.Code = 1011
	Token_Expired     codes.Code = 1012
	System_Err        codes.Code = 1013
	Admin_Not_Found   codes.Code = 1014
	Password_Err      codes.Code = 1015
	Admin_Exists      codes.Code = 1016
	Node_Not_Found    codes.Code = 1017
	Config_Key_Exists codes.Code = 1018
	Project_Not_Found codes.Code = 1019
	Config_Not_Found  codes.Code = 1020
	Project_Exists    codes.Code = 1021
	No_Access         codes.Code = 1022
)

func SystemErr[T any](err error) (T, error) {
	var m T
	return Err(System_Err, err, m)
}

func Err[T any](code codes.Code, err error, m T) (T, error) {
	return m, status.Errorf(code, err.Error())
}

func Succ[T any](m T) (T, error) {
	return m, nil
}
