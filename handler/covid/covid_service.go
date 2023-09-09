package covid

import (
	"covid/handler/model"
	"covid/internal/response"
)

func getAgeGroup(age int) string {
	switch {
	case age >= 0 && age <= 30:
		return "0-30"
	case age >= 31 && age <= 60:
		return "31-60"
	case age >= 61:
		return "61+"
	default:
		return "N/A"
	}
}

func GetCovidSummaryService() (*model.CovidSummary, *response.Response) {

	covidCase, err := GetCovidCasesAPI()
	if err != nil {
		return nil, err
	}

	provinceCount := make(map[string]int)
	ageGroupCount := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	for _, record := range covidCase.Data {
		provinceCount[record.Province]++
		ageGroup := getAgeGroup(record.Age)
		ageGroupCount[ageGroup]++
	}

	covidSum := &model.CovidSummary{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}

	return covidSum, nil
}
