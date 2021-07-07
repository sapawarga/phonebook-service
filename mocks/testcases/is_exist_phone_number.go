package testcases

import "errors"

// ReponseIsExistPhone ...
type ReponseIsExistPhone struct {
	Result bool
	Error  error
}

// IsExistPhoneNumber ...
type IsExistPhoneNumber struct {
	Description        string
	UsecaseParams      string
	RepositoryParams   string
	UsecaseResponse    ReponseIsExistPhone
	RepositoryResponse ReponseIsExistPhone
}

// IsExistPhoneNumberData ...
var IsExistPhoneNumberData = []IsExistPhoneNumber{
	{
		Description:      "success_get_check_phone_number",
		UsecaseParams:    "022-1234",
		RepositoryParams: "022-1234",
		UsecaseResponse: ReponseIsExistPhone{
			Result: true,
			Error:  nil,
		},
		RepositoryResponse: ReponseIsExistPhone{
			Result: true,
			Error:  nil,
		},
	}, {
		Description:      "failed_check_phone_number",
		UsecaseParams:    "sample",
		RepositoryParams: "sample",
		UsecaseResponse: ReponseIsExistPhone{
			Result: false,
			Error:  errors.New("failed_check"),
		},
		RepositoryResponse: ReponseIsExistPhone{
			Result: false,
			Error:  errors.New("failed_check"),
		},
	},
}

// IsExistDescription ...
func IsExistDescription() []string {
	var arr = []string{}
	for _, data := range IsExistPhoneNumberData {
		arr = append(arr, data.Description)
	}
	return arr
}
