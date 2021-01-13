package rates

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UsecaseMock struct {
	mock.Mock
}

func (u *UsecaseMock) GetRates(prefix string) (interface{}, error) {
	args := u.Called()
	return nil, args.Error(1)
}
func Test_UseCase_GetLatestRates(t *testing.T) {
	t.Run("Test usecase get rate with prefix latest success", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "latest"
		mockRepository.On("LatestRates").Return(nil, nil)
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("Test usecase get rate with prefix latest fail", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "latest"
		mockRepository.On("LatestRates").Return(nil, errors.New(""))
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
	t.Run("Test usecase get rate with analyze success", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "analyze"
		mockRepository.On("RatesAnalyze").Return(nil, nil)
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("Test usecase get rate with analyze fail", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "analyze"
		mockRepository.On("RatesAnalyze").Return(nil, errors.New(""))
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
	t.Run("Test usecase get rate by date success", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "2019-08-19"
		mockRepository.On("DateRates", prefix).Return(nil, nil)
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.Equal(t, nil, err)
	})
	t.Run("Test usecase get rate by date fail", func(t *testing.T) {
		t.Parallel()
		mockRepository := new(RepositoryMock)
		prefix := "2019-08-19"
		mockRepository.On("DateRates", prefix).Return(nil, errors.New(""))
		u := NewUsecase(mockRepository, &http.Client{})
		_, err := u.GetRates(prefix)
		assert.NotEqual(t, nil, err)
	})
}
