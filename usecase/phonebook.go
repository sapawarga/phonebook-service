package usecase

import (
	"context"
	"fmt"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/sapawarga/phonebook-service/model"
	"github.com/sapawarga/phonebook-service/repository"
)

// PhoneBook ...
type PhoneBook struct {
	repo   repository.PhoneBookI
	logger kitlog.Logger
}

// NewPhoneBook ...
func NewPhoneBook(repo repository.PhoneBookI, logger kitlog.Logger) *PhoneBook {
	return &PhoneBook{
		repo:   repo,
		logger: logger,
	}
}

// GetList ...
func (pb *PhoneBook) GetList(ctx context.Context, params *model.ParamsPhoneBook) (*model.PhoneBookWithMeta, error) {
	logger := kitlog.With(pb.logger, "method", "GetList")
	resp, err := pb.repo.GetListPhoneBook(ctx, &model.GetListRequest{
		Name:  params.Name,
		Limit: params.Limit,
		Page:  params.Page,
	})
	fmt.Println(resp)
	if err != nil {
		level.Error(logger).Log("error", err)
		return nil, err
	}

	total, err := pb.repo.GetMetaDataPhoneBook(ctx, &model.GetListRequest{
		Name:  params.Name,
		Limit: params.Limit,
		Page:  params.Page,
	})

	return &model.PhoneBookWithMeta{
		PhoneBooks: resp,
		Page:       params.Page,
		Total:      total,
	}, nil
}
