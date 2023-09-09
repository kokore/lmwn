package covid

import (
	"covid/handler/model"
	"covid/internal/response"
)

func GetCovidSummaryService() (*model.CovidData, *response.Response) {

	covidCase, err := GetCovidCasesAPI()
	if err != nil {
		return nil, err
	}

	return covidCase, nil
}
