package service_test

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/labstack/gommon/log"
	. "github.com/panoplymedia/sales_admin/cmd/server/internal/service"
	"github.com/panoplymedia/sales_admin/internal/mock"
	"github.com/panoplymedia/sales_admin/internal/sales"
)

func TestUpload(t *testing.T) {
	tests := map[string]struct {
		request              *http.Request
		expectedResponseCode int
		salesService         *mock.SalesService
		addInvoked           bool
	}{
		"should update records and return status ok": {
			request:              uploadRequest(t, UploadFileFormField, "./testdata/good.csv"),
			expectedResponseCode: http.StatusSeeOther,
			salesService: &mock.SalesService{
				AddFn: func(sales ...sales.Sale) error {
					// TODO(wh): Check that we get the expected sales.
					return nil
				},
			},
			addInvoked: true,
		},
		"should return BadRequest if form does not have the file field": {
			request:              uploadRequest(t, "badFieldName", "./testdata/good.csv"),
			expectedResponseCode: http.StatusBadRequest,
			salesService:         &mock.SalesService{},
			addInvoked:           false,
		},
		"should return BadRequest for invalid CSV": {
			request:              uploadRequest(t, UploadFileFormField, "./testdata/invalid.csv"),
			expectedResponseCode: http.StatusBadRequest,
			salesService:         &mock.SalesService{},
			addInvoked:           false,
		},
		"should return InternalServerError if db is down": {
			request:              uploadRequest(t, UploadFileFormField, "./testdata/good.csv"),
			expectedResponseCode: http.StatusInternalServerError,
			salesService: &mock.SalesService{
				AddFn: func(sales ...sales.Sale) error {
					return errors.New("DB down")
				},
			},
			addInvoked: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			logger := log.New("test")
			logger.SetOutput(ioutil.Discard)

			s, err := New(
				WithSalesService(test.salesService),
				WithTemplates(templatesGlob),
				WithLogger(logger),
			)
			if err != nil {
				t.Fatalf("Unexpected error returned when creating the new service: %s", err)
			}

			recorder := httptest.NewRecorder()
			s.HTTPHandler().ServeHTTP(recorder, test.request)

			response := recorder.Result()
			if response.StatusCode != test.expectedResponseCode {
				t.Fatalf("Unexpected response code: Expected %v, but got %v", test.expectedResponseCode, response.StatusCode)
			}

			if test.salesService.AddInvoked != test.addInvoked {
				t.Fatalf("Expected the sales service Add invoked does not match expected: Expected %v, but got %v", test.addInvoked, test.salesService.AddInvoked)
			}
		})
	}
}

func uploadRequest(t *testing.T, fieldname, filepath string) *http.Request {
	t.Helper()
	body, contentType := uploadCSVRequestBody(t, fieldname, filepath)
	r := httptest.NewRequest(http.MethodPost, "/sales/upload", body)
	r.Header.Set("Content-Type", contentType)
	return r
}

func uploadCSVRequestBody(t *testing.T, fieldname, path string) (body *bytes.Buffer, requestType string) {
	t.Helper()
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("Unexpected error returned when attempting to open '%s': %s", path, err)

	}
	defer file.Close()

	body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fieldname, filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Fatalf("Unexpected error returned when attempting to create the form file: %s", err)
	}

	io.Copy(part, file)
	err = writer.Close()
	if err != nil {
		t.Fatalf("Unexpected error returned when attempting to close the writer: %s", err)
	}

	return body, writer.FormDataContentType()
}

func TestListSales(t *testing.T) {
	tests := map[string]struct {
		request      *http.Request
		salesService *mock.SalesService

		expectedResponseCode int
		getAllInvoked        bool
	}{
		"should return success if anything is pulled from the database": {
			request: httptest.NewRequest("GET", "/sales", nil),
			salesService: &mock.SalesService{
				GetAllFn: func() ([]sales.Sale, error) {
					// Could return something here so that we can check the
					// sales that are returned, but we would checking the
					// template and would be a brittle test.
					return nil, nil
				},
			},

			expectedResponseCode: http.StatusOK,
			getAllInvoked:        true,
		},
		"should return InternalServerError if the database is down": {
			request: httptest.NewRequest("GET", "/sales", nil),
			salesService: &mock.SalesService{
				GetAllFn: func() ([]sales.Sale, error) {
					// Could return something here so that we can check the
					// sales that are returned, but we would checking the
					// template and would be a brittle test.
					return nil, errors.New("DB error")
				},
			},

			expectedResponseCode: http.StatusInternalServerError,
			getAllInvoked:        true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			logger := log.New("test")
			logger.SetOutput(ioutil.Discard)

			s, err := New(
				WithSalesService(test.salesService),
				WithTemplates(templatesGlob),
				WithLogger(logger),
			)
			if err != nil {
				t.Fatalf("Unexpected error returned when creating the new service: %s", err)
			}

			recorder := httptest.NewRecorder()
			s.HTTPHandler().ServeHTTP(recorder, test.request)

			response := recorder.Result()
			if response.StatusCode != test.expectedResponseCode {
				t.Fatalf("Unexpected response code: Expected %v, but got %v", test.expectedResponseCode, response.StatusCode)
			}

			if test.salesService.GetAllInvoked != test.getAllInvoked {
				t.Fatalf("Expected the sales service Add invoked does not match expected: Expected %v, but got %v", test.getAllInvoked, test.salesService.GetAllInvoked)
			}
		})
	}
}
