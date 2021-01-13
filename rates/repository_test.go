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
		// mocked := new(RepositoryMock)
		// m := NewMockRepository(mocked)

		// app := RepositoryMock{Mock: mocked}

		// _, err := app.LatestRates()

		// assert.Error(t, err)
		// assert.Equal(t, nil, err)

	})
}
