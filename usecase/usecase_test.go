package usecase_test

import (
	"context"
	"fmt"
	"os"
	"reflect"

	mocks "github.com/sapawarga/phonebook-service/mocks"
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
		mockPhoneBookRepo.EXPECT().GetCategoryNameByID(ctx, data.GetCategoryNameParams).Return(data.MockCategorydata.Result, data.MockCategorydata.Error).Times(len(data.MockCategorydata.Result) * 2)
		mockPhoneBookRepo.EXPECT().GetListPhonebookByLongLat(ctx, gomock.Any()).Return(data.MockGetList.Result, data.MockGetList.Error).Times(1)
		mockPhoneBookRepo.EXPECT().GetListPhonebookByLongLatMeta(ctx, gomock.Any()).Return(data.MockGetMetadata.Result, data.MockGetMetadata.Error).Times(1)
		resp, err := phonebook.GetList(ctx, &model.ParamsPhoneBook{
			Search:    data.UsecaseParams.Search,
			Limit:     data.UsecaseParams.Limit,
			Page:      data.UsecaseParams.Page,
			Longitude: data.UsecaseParams.Longitude,
			Latitude:  data.UsecaseParams.Latitude,
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
		mockPhoneBookRepo.EXPECT().GetPhonebookDetailByID(ctx, data.GetDetailRequest).Return(data.MockPhonebookDetail.Result, data.MockPhonebookDetail.Error)
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

	// InsertPhonebookLogic ...
	var InsertPhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.InsertPhonebookTestcases[idx]
		mockPhoneBookRepo.EXPECT().GetCategoryNameByID(ctx, data.GetCategoryNameRequest).Return(data.MockCategory.Result, data.MockCategory.Error).Times(1)
		mockPhoneBookRepo.EXPECT().Insert(ctx, &data.RepositoryRequest).Return(data.RepositoryResponse).Times(1)
		err := phonebook.Insert(ctx, &data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
		}
	}

	// UpdatePhonebookLogic ...
	var UpdatePhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.UpdatePhonebookTestcases[idx]
		mockPhoneBookRepo.EXPECT().GetPhonebookDetailByID(ctx, data.GetDetailRepositoryRequest).Return(data.MockDetailRepository.Result, data.MockDetailRepository.Error).Times(1)
		mockPhoneBookRepo.EXPECT().Update(ctx, &data.UpdateRepositoryRequest).Return(data.MockUpdateRepository).Times(1)
		err := phonebook.Update(ctx, &data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	// DeletePhonebookLogic ...
	var DeletePhonebookLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.DeletePhonebookTestcases[idx]
		mockPhoneBookRepo.EXPECT().GetPhonebookDetailByID(ctx, data.GetDetailRepositoryRequest).Return(data.MockDetailRepository.Result, data.MockDetailRepository.Error).Times(1)
		mockPhoneBookRepo.EXPECT().Delete(ctx, data.DeleteRepositoryRequest).Return(data.MockDeleteRepository).Times(1)
		err := phonebook.Delete(ctx, data.UsecaseRequest)
		if err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	var CheckReadinessLogic = func(idx int) {
		ctx := context.Background()
		data := testcases.CheckReadinessData[idx]
		mockPhoneBookRepo.EXPECT().CheckHealthReadiness(ctx).Return(data.MockCheckReadiness).Times(1)
		if err := phonebook.CheckHealthReadiness(ctx); err != nil {
			Expect(err).NotTo(BeNil())
		} else {
			Expect(err).To(BeNil())
		}
	}

	// sort all function names
	var unitTestLogic = map[string]map[string]interface{}{
		"GetList":        {"func": GetListLogic, "test_case_count": len(testcases.GetPhoneBookData), "desc": testcases.ListPhonebookDescription()},
		"GetDetail":      {"func": GetDetailPhonebookLogic, "test_case_count": len(testcases.GetDetailPhonebookData), "desc": testcases.DetailPhonebookDescription()},
		"Insert":         {"func": InsertPhonebookLogic, "test_case_count": len(testcases.InsertPhonebookTestcases), "desc": testcases.InsertPhonebookDescription()},
		"Update":         {"func": UpdatePhonebookLogic, "test_case_count": len(testcases.UpdatePhonebookTestcases), "desc": testcases.UpdatePhonebookDescription()},
		"Delete":         {"func": DeletePhonebookLogic, "test_case_count": len(testcases.DeletePhonebookTestcases), "desc": testcases.DeletePhonebookDescription()},
		"CheckReadiness": {"func": CheckReadinessLogic, "test_case_count": len(testcases.CheckReadinessData), "desc": testcases.CheckReadinessDescription()},
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
