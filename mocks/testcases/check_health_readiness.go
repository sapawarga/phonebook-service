package testcases

import "errors"

type CheckReadiness struct {
	Description        string
	MockCheckReadiness error
	MockUsecase        error
}

var CheckReadinessData = []CheckReadiness{
	{
		Description:        "check_readiness_database_ok",
		MockCheckReadiness: nil,
		MockUsecase:        nil,
	}, {
		Description:        "check_readiness_database_not_ok",
		MockCheckReadiness: errors.New("something_went_wrong"),
		MockUsecase:        errors.New("something_went_wrong"),
	},
}

func CheckReadinessDescription() []string {
	var arr = []string{}
	for _, data := range CheckReadinessData {
		arr = append(arr, data.Description)
	}
	return arr
}
