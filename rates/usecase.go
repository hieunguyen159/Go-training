package rates

import (
	"net/http"
)

type UsecaseInterface interface {
	GetRates(prefix string) (interface{}, error)
	// GetLatestDate() (string, error)
	// ImportDataInit(cubes *Cube) error
	// GetXML(url string) ([]byte, error)
}

// Usecase struct
type Usecase struct {
	repo   RepositoryInterface
	Client *http.Client
}

func NewUsecase(r RepositoryInterface, Client *http.Client) UsecaseInterface {
	return &Usecase{r, Client}
}
func (u *Usecase) GetRates(prefix string) (interface{}, error) {
	switch prefix {
	case "latest":
		resp, err := u.repo.LatestRates()
		if err != nil {
			return nil, err
		}
		return resp, err
	case "analyze":
		resp, err := u.repo.RatesAnalyze()

		if err != nil {
			return nil, err
		}
		return resp, err
	default:
		resp, err := u.repo.DateRates(prefix)
		if err != nil {
			return nil, err
		}
		return resp, err
	}
}
