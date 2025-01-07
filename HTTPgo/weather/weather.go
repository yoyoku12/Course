package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

func GetWeather(geo geo.GeoData, format int) string {

	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err := http.Get(baseUrl.String()) //получаем ответ
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) //вычитываем тело в байты
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(body)
}
