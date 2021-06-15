package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sapawarga/phonebook-service/endpoint"
	"github.com/sapawarga/phonebook-service/helper"
	"github.com/sapawarga/phonebook-service/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

func MakeHealthyCheckHandler(ctx context.Context, fs usecase.Provider, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()
	r.Handle("/health/live", kithttp.NewServer(endpoint.MakeCheckHealthy(ctx), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
	r.Handle("/health/ready", kithttp.NewServer(endpoint.MakeCheckReadiness(ctx, fs), decodeNoRequest, encodeResponse, opts...)).Methods(helper.HTTP_GET)
	return r
}

func MakeHTTPHandler(ctx context.Context, fs usecase.Provider, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processGetList := kithttp.NewServer(endpoint.MakeGetList(ctx, fs), decodeGetListRequest, encodeResponse, opts...)
	processGetDetail := kithttp.NewServer(endpoint.MakeGetDetail(ctx, fs), decodeGetByID, encodeResponse, opts...)
	processAdd := kithttp.NewServer(endpoint.MakeAddPhonebook(ctx, fs), decodeCreateRequest, encodeResponse, opts...)
	processUpdate := kithttp.NewServer(endpoint.MakeUpdatePhonebook(ctx, fs), decodeUpdateRequest, encodeResponse, opts...)
	processDelete := kithttp.NewServer(endpoint.MakeDeletePhonebook(ctx, fs), decodeGetByID, encodeResponse, opts...)

	r := mux.NewRouter()

	// TODO: handle token middleware
	r.Handle("/phone-books/", processGetList).Methods(helper.HTTP_GET)
	r.Handle("/phone-books/{id}", processGetDetail).Methods(helper.HTTP_GET)
	r.Handle("/phone-books/", processAdd).Methods(helper.HTTP_POST)
	r.Handle("/phone-books/{id}", processUpdate).Methods(helper.HTTP_PUT)
	r.Handle("/phone-books/{id}", processDelete).Methods(helper.HTTP_DELETE)

	return r
}

func decodeGetListRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	search := r.URL.Query().Get("search")
	regIDString := r.URL.Query().Get("regency_id")
	distIDString := r.URL.Query().Get("district_id")
	vilIDString := r.URL.Query().Get("village_id")
	statusString := r.URL.Query().Get("status")
	limitString := r.URL.Query().Get("limit")
	pageString := r.URL.Query().Get("page")

	if limitString == "" {
		limitString = "10"
	}
	if pageString == "" || pageString == "0" {
		pageString = "1"
	}

	_, regID := helper.ConvertFromStringToInt64(regIDString)
	_, disID := helper.ConvertFromStringToInt64(distIDString)
	_, vilID := helper.ConvertFromStringToInt64(vilIDString)
	status, _ := helper.ConvertFromStringToInt64(statusString)
	_, limit := helper.ConvertFromStringToInt64(limitString)
	_, page := helper.ConvertFromStringToInt64(pageString)

	return &endpoint.GetListRequest{
		Search:     search,
		RegencyID:  regID,
		DistrictID: disID,
		VillageID:  vilID,
		Status:     status,
		Limit:      limit,
		Page:       page,
		Latitude:   r.URL.Query().Get("latitude"),
		Longitude:  r.URL.Query().Get("longitude"),
	}, nil
}

func decodeGetByID(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := helper.ConvertFromStringToInt64(params["id"])

	return &endpoint.GetDetailRequest{
		ID: id,
	}, nil
}

func decodeCreateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	reqBody := &endpoint.AddPhonebookRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}

	return reqBody, nil
}

func decodeUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	params := mux.Vars(r)
	_, id := helper.ConvertFromStringToInt64(params["id"])

	reqBody := &endpoint.UpdatePhonebookRequest{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		return nil, err
	}

	reqBody.ID = id
	return reqBody, nil
}

func decodeNoRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(err); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	status, ok := response.(*endpoint.StatusResponse)
	if ok && status.Code != helper.STATUS_OK {
		if status.Code == helper.STATUS_CREATED {
			w.WriteHeader(http.StatusCreated)
		} else if status.Code == helper.STATUS_UPDATED || status.Code == helper.STATUS_DELETED {
			w.WriteHeader(http.StatusNoContent)
			_ = json.NewEncoder(w).Encode(nil)
			return nil
		}
	}

	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}
