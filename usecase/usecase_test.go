package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"github.com/sapawarga/phonebook-service/mocks"
	"github.com/sapawarga/phonebook-service/mocks/testcases"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/golang/mock/gomock"
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
		mockPhoneBookRepo.EXPECT().GetListPhoneBook(ctx, gomock.Any()).Return(data.MockGetList.Result, data.MockGetList.Error).Times(1)
		mockPhoneBookRepo.EXPECT().GetMetaDataPhoneBook(ctx, gomock.Any()).Return(data.MockGetMetadata.Result, data.MockGetMetadata.Error).Times(1)
		resp, err := phonebook.GetList(ctx, &model.ParamsPhoneBook{
			Search: data.UsecaseParams.Search,
			Limit:  data.UsecaseParams.Limit,
			Page:   data.UsecaseParams.Page,
		})
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp.Page).To(Equal(data.MockUsecase.Result.Page))
			Expect(resp.Total).To(Equal(data.MockUsecase.Result.Total))
			Expect(resp).NotTo(BeNil())
		}
	}

	// GetDetailPhonebookLogic ...
	var GetDetailPhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.GetDetailPhonebookData[idx]
		mockPhoneBookRepo.EXPECT().GetCategoryNameByID(ctx, gomock.Any()).Return(data.MockCategory.Result, data.MockCategory.Error).Times(1)
		mockPhoneBookRepo.EXPECT().GetLocationNameByID(ctx, gomock.Any()).Return(data.MockLocation.Result, data.MockLocation.Error).Times(3)
		mockPhoneBookRepo.EXPECT().GetPhonebookDetailByID(ctx, data.GetDetailRequest).Return(data.MockPhonebookDetail.Result, data.MockPhonebookDetail.Error).Times(1)
		resp, err := phonebook.GetDetail(ctx, data.UsecaseParams)
		if err != nil {
			Expect(err).NotTo(BeNil())
			Expect(resp).To(BeNil())
		} else {
			Expect(err).To(BeNil())
			Expect(resp).NotTo(BeNil())
			Expect(resp).To(Equal(data.MockUsecase.Result))
		}
	}

	// sort all function names
	var unitTestLogic = map[string]map[string]interface{}{
		"GetList":   {"func": GetListLogic, "test_case_count": len(testcases.GetPhoneBookData), "desc": testcases.ListPhonebookDescription()},
		"GetDetail": {"func": GetDetailPhonebookLogic, "test_case_count": len(testcases.GetDetailPhonebookData), "desc": testcases.DetailPhonebookDescription()},
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
