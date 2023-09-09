package covid

import (
	"covid/handler/model"
	"reflect"
	"testing"
)

func TestGetAgeGroups(t *testing.T) {
	tests := []struct {
		age      int
		expected string
	}{
		{15, "0-30"},
		{40, "31-60"},
		{65, "61+"},
		{70, "61+"},
		{-5, "N/A"},
	}

	for _, test := range tests {
		result := getAgeGroup(test.age)
		if result != test.expected {
			t.Errorf("returned %s, expected %s", result, test.expected)
		}
	}
}

func TestGetCovidSummaryService(t *testing.T) {
	covidData := &model.CovidData{
		Data: []model.CovidRecord{
			{Age: 25, Province: "A"},
			{Age: 40, Province: "B"},
			{Age: 70, Province: "A"},
		},
	}

	provinceCount := make(map[string]int)
	ageGroupCount := map[string]int{
		"0-30":  0,
		"31-60": 0,
		"61+":   0,
		"N/A":   0,
	}

	for _, record := range covidData.Data {
		provinceCount[record.Province]++
		ageGroup := getAgeGroup(record.Age)
		ageGroupCount[ageGroup]++
	}

	covidSum := &model.CovidSummary{
		Province: provinceCount,
		AgeGroup: ageGroupCount,
	}

	expectedSummary := &model.CovidSummary{
		Province: map[string]int{"A": 2, "B": 1},
		AgeGroup: map[string]int{"0-30": 1, "31-60": 1, "61+": 1, "N/A": 0},
	}

	if !reflect.DeepEqual(covidSum, expectedSummary) {
		t.Errorf("GetCovidSummaryService returned %v, expected %v", covidSum, expectedSummary)
	}
}
