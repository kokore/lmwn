package covid

import (
	"covid/handler/model"
	"covid/internal/config"
	"covid/internal/response"
	"encoding/json"
	"net/http"
)

func GetCovidCasesAPI() (*model.CovidData, *response.Response) {
	result, err := http.Get(config.Config.CovidAPI.CovidCaseURL)
	if err != nil {
		return nil, response.Err(response.UnableGetCovidCase, http.StatusInternalServerError, "can't get covid case api")
	}
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		return nil, response.Err(response.UnableGetCovidCase, http.StatusInternalServerError, "can't get covid case api")
	}

	var apiResponse model.CovidData
	err = json.NewDecoder(result.Body).Decode(&apiResponse)
	if err != nil {
		return nil, response.Err(response.UnableMapCovidCaseToObject, http.StatusInternalServerError, "can't map object")
	}

	return &apiResponse, nil
}
