package rates

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (u *RepositoryMock) LatestRates() (*RatesResponse, error) {
	args := u.Called()
	return nil, args.Error(1)
}

func (u *RepositoryMock) DateRates(date string) (*RatesResponse, error) {
	args := u.Called(date)
	return nil, args.Error(1)
}
func (u *RepositoryMock) RatesAnalyze() (*[]ValuePerCurrency, error) {
	args := u.Called()
	return nil, args.Error(1)
}

func TestLatestRates(t *testing.T) {
	t.Run("Test get latest rates success", func(t *testing.T) {
		t.Parallel()
		mockMap := make(map[string]float64)
		models := RatesResponse{Date: "2020-09-29", Rates: mockMap}
		mocked := new(RepositoryMock)
		mocked.On("Insert", models)
	})
}
