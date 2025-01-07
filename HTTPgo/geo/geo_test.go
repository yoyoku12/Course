package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {

	// Arrange - подготовка теста, ожидаемый результат, данные для функции
	city := "Moscow"
	expected := geo.GeoData{
		City: "Moscow",
	}

	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Assert - проверка результата с ожиданиями
	if err != nil {
		t.Error("Ошибка получения города")
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получено %v", expected, got)
	}

}
