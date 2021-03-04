package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/golang/mock/gomock"
	"github.com/sapawarga/phonebook-service/mocks"
	"github.com/sapawarga/phonebook-service/mocks/testcases"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Phone Book", func() {
	var (
		mockPhoneBookRepo *mocks.MockPhoneBookI
		phonebook         usecase.Provider
	)

	BeforeEach(func() {
		logger := kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
		mockSvc := gomock.NewController(GinkgoT())
		mockSvc.Finish()
		mockPhoneBookRepo = mocks.NewMockPhoneBookI(mockSvc)
		phonebook = usecase.NewPhoneBook(mockPhoneBookRepo, logger)
	})

	// DECLARE UNIT TEST FUNCTION

	// GetListLogic ...
	var GetListLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetPhoneBookData[idx]
		mockPhoneBookRepo.EXPECT().GetListPhoneBook(ctx, &data.GetListParams).Return(data.MockGetList.Result, data.MockGetList.Error).Times(1)
		mockPhoneBookRepo.EXPECT().GetMetaDataPhoneBook(ctx, &data.GetMetaDataParams).Return(data.MockGetMetadata.Result, data.MockGetMetadata.Error).Times(1)
		resp, err := phonebook.GetList(ctx, &model.ParamsPhoneBook{
			Name:  data.UsecaseParams.Name,
			Limit: data.UsecaseParams.Limit,
			Page:  data.UsecaseParams.Page,
		})
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp.Page).To(Equal(data.MockUsecase.Result.Page))
			Expect(resp.Total).To(Equal(data.MockUsecase.Result.Total))
		}
	}

	// sort all function names
	var unitTestLogic = map[string]map[string]interface{}{
		"GetList": {"func": GetListLogic, "test_case_count": len(testcases.GetPhoneBookData), "desc": testcases.Description()},
	}

	for _, val := range unitTestLogic {
		s := reflect.ValueOf(val["desc"])
		var arr []TableEntry
		for i := 0; i < val["test_case_count"].(int); i++ {
			fmt.Println(s.Index(i).String())
			arr = append(arr, Entry(s.Index(i).String(), i))
		}
		DescribeTable("Function ", val["func"], arr...)
	}
})
