package rates

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"
)

func Test_Handler_GetRates_Latest(t *testing.T) {
	t.Run("Test handler get latest rates success", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/rates/latest", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mockUseCase := new(UsecaseMock)
		mockUseCase.On("GetRates", mock.Anything).Return(nil, nil)
		h := NewHandler(mockUseCase)
		handler := http.HandlerFunc(h.GetRates)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("Test handler get analyze rates success", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/rates/analyze", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mockUseCase := new(UsecaseMock)
		mockUseCase.On("GetRates", mock.Anything).Return(nil, nil)
		h := NewHandler(mockUseCase)
		handler := http.HandlerFunc(h.GetRates)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("Test handler get rates by date success", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/rates/2019-08-19", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mockUseCase := new(UsecaseMock)
		mockUseCase.On("GetRates", mock.Anything).Return(nil, nil)
		h := NewHandler(mockUseCase)
		handler := http.HandlerFunc(h.GetRates)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusOK, rr.Code)
	})
	t.Run("Test handler get rates by date invalid", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/rates/2019-23-19", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mockUseCase := new(UsecaseMock)
		mockUseCase.On("GetRates", mock.Anything).Return(nil, nil)
		h := NewHandler(mockUseCase)
		handler := http.HandlerFunc(h.GetRates)
		handler.ServeHTTP(rr, req)
		body := rr.Body.String()
		assert.Equal(t, `{"Code":400,"Message":"GetRatesDate: date invalid format"}
`, body)
	})
	t.Run("Test handler get rates fail 500", func(t *testing.T) {
		t.Parallel()
		req, err := http.NewRequest("GET", "/rates/latest", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		mockUseCase := new(UsecaseMock)
		mockUseCase.On("GetRates", mock.Anything).Return(nil, errors.New("Server error"))
		h := NewHandler(mockUseCase)
		handler := http.HandlerFunc(h.GetRates)
		handler.ServeHTTP(rr, req)
		body := rr.Body.String()
		assert.Equal(t, `{"Code":500,"Message":"GetRates: error server"}
`, body)
	})
}
