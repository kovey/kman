package code

import "github.com/kovey/kow/result"

const (
	Token_Err         result.Codes = 1011
	Token_Expired     result.Codes = 1012
	System_Err        result.Codes = 1013
	Admin_Not_Found   result.Codes = 1014
	Password_Err      result.Codes = 1015
	Admin_Exists      result.Codes = 1016
	Node_Not_Found    result.Codes = 1017
	Config_Key_Exists result.Codes = 1018
	Project_Not_Found result.Codes = 1019
	Config_Not_Found  result.Codes = 1020
	Project_Exists    result.Codes = 1021
	No_Access         result.Codes = 1022
)
