package rates

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)

		_, err := m.LatestRates()
		assert.Equal(t, nil, err)

	})
	t.Run("Test get ratest latest fail", func(t *testing.T) {
		t.Parallel()
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)

		_, err := m.LatestRates()

		assert.NotEqual(t, nil, err)
	})
}
func TestGetRatesByDate(t *testing.T) {
	t.Run("Test get rates by date success", func(t *testing.T) {
		t.Parallel()
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)

		_, err := m.DateRates("2019-08-19")
		assert.Equal(t, nil, err)
	})
	t.Run("Test get rates by date fail", func(t *testing.T) {
		t.Parallel()
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)

		_, err := m.DateRates("2019-08-19")
		assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, err)
	})
}

func TestGetRatesAnalyze(t *testing.T) {
	t.Run("Test get rates analyze success", func(t *testing.T) {
		t.Parallel()
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)
		actual, err := m.RatesAnalyze()
		assert.Equal(t, nil, err)
		assert.NotEqual(t, nil, actual)
	})
	t.Run("Test get rates analyze fail", func(t *testing.T) {
		t.Parallel()
		mocked := new(RepositoryMock)
		m := NewMockRepository(mocked)
		_, err := m.RatesAnalyze()
		assert.NotEqual(t, nil, err)
	})
}
