package rates

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	LatestRates() (*RatesResponse, error)
	DateRates(date string) (*RatesResponse, error)
	RatesAnalyze() (*[]ValuePerCurrency, error)
}

// Repository struct
type Repository struct {
	DB *mongo.Collection
}

func NewRepository(DB *mongo.Collection) RepositoryInterface {
	return &Repository{DB}
}

func (r *Repository) LatestRates() (*RatesResponse, error) {
	var (
		Cubes     []Cubes
		dateCubes RatesResponse
	)
	// dao, _ := db.LoadConfig()
	data, _ := r.DB.Find(context.Background(), bson.M{})

	defer data.Close(context.Background())
	err := data.All(context.Background(), &Cubes)
	// fmt.Println(Cubes)
	if err == nil {
		rateResult := make(map[string]float64)
		allCubes := make([]RatesResponse, 0)

		for _, cubes := range Cubes {
			dateCubes.Date = cubes.Time
			for _, cube := range cubes.Cubes {
				rateResult[cube.Currency] = cube.Rate
				dateCubes.Rates = rateResult
			}
			allCubes = append(allCubes, dateCubes)
		}
		return &allCubes[0], nil
	}
	return &RatesResponse{}, nil
}
func (r *Repository) DateRates(date string) (*RatesResponse, error) {
	if date != "" {
		var (
			Cubes     []Cubes
			DateCubes RatesResponse
		)
		// dao, _ := db.LoadConfig()
		data, _ := r.DB.Find(context.TODO(), bson.M{"time": date})
		defer data.Close(context.TODO())
		err := data.All(context.TODO(), &Cubes)
		if err == nil {
			rateResult := make(map[string]float64)
			for _, cubes := range Cubes {
				DateCubes.Date = cubes.Time
				for _, cube := range cubes.Cubes {
					rateResult[cube.Currency] = cube.Rate
					DateCubes.Rates = rateResult
				}
			}
		}
		return &DateCubes, nil
	}
	return &RatesResponse{}, nil
}

func (r *Repository) RatesAnalyze() (*[]ValuePerCurrency, error) {
	var getDataCube []ValuePerCurrency
	matchStage := bson.M{"$unwind": "$Cube"}
	groupStage := bson.M{
		"$group": bson.M{
			"_id": "$Cube.currency",
			"max": bson.M{
				"$max": "$Cube.rate",
			},
			"min": bson.M{
				"$min": "$Cube.rate",
			},
			"avg": bson.M{
				"$avg": "$Cube.rate",
			}},
	}
	// dao, _ := db.LoadConfig()
	getDataCubeCusor, err := r.DB.Aggregate(context.Background(), []bson.M{matchStage, groupStage})
	if err != nil {
		panic(err)
	}
	if err = getDataCubeCusor.All(context.Background(), &getDataCube); err != nil {
		return &[]ValuePerCurrency{}, nil
	}
	return &getDataCube, nil
}
