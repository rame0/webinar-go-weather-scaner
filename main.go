package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type WeatherInfo struct {
	XMLName xml.Name `xml:"MMWEATHER"`
	Text    string   `xml:",chardata"`
	REPORT  struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
		TOWN struct {
			Text      string `xml:",chardata"`
			Index     string `xml:"index,attr"`
			Sname     string `xml:"sname,attr"`
			Latitude  string `xml:"latitude,attr"`
			Longitude string `xml:"longitude,attr"`
			FORECAST  []struct {
				Text      string `xml:",chardata"`
				Day       string `xml:"day,attr"`
				Month     string `xml:"month,attr"`
				Year      string `xml:"year,attr"`
				Hour      string `xml:"hour,attr"`
				Tod       string `xml:"tod,attr"`
				Predict   string `xml:"predict,attr"`
				Weekday   string `xml:"weekday,attr"`
				PHENOMENA struct {
					Text          string `xml:",chardata"`
					Cloudiness    string `xml:"cloudiness,attr"`
					Precipitation string `xml:"precipitation,attr"`
					Rpower        string `xml:"rpower,attr"`
					Spower        string `xml:"spower,attr"`
				} `xml:"PHENOMENA"`
				PRESSURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"PRESSURE"`
				TEMPERATURE struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"TEMPERATURE"`
				WIND struct {
					Text      string `xml:",chardata"`
					Min       string `xml:"min,attr"`
					Max       string `xml:"max,attr"`
					Direction string `xml:"direction,attr"`
				} `xml:"WIND"`
				RELWET struct {
					Text string `xml:",chardata"`
					Max  string `xml:"max,attr"`
					Min  string `xml:"min,attr"`
				} `xml:"RELWET"`
				HEAT struct {
					Text string `xml:",chardata"`
					Min  string `xml:"min,attr"`
					Max  string `xml:"max,attr"`
				} `xml:"HEAT"`
			} `xml:"FORECAST"`
		} `xml:"TOWN"`
	} `xml:"REPORT"`
}

func main() {
	resp, err := http.Get("https://www.meteoservice.ru/en/export/gismeteo?point=37")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	byteVal, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var weather WeatherInfo

	err = xml.Unmarshal(byteVal, &weather)
	if err != nil {
		log.Fatal(err)
	}

	forcast := weather.REPORT.TOWN.FORECAST
	sname, _ := url.PathUnescape(weather.REPORT.TOWN.Sname)
	fmt.Println(sname)
	for i := 0; i < len(forcast); i++ {
		fmt.Print(
			forcast[i].Day + "/" + forcast[i].Month + "/" + forcast[i].Year +
				forcast[i].Hour)
		var daytime = "  "
		switch forcast[i].Tod {
		case "0":
			daytime += "Ночью:"
			break
		case "1":
			daytime += "Утром:"
			break
		case "2":
			daytime += "Днем:"
			break
		case "3":
			daytime += "Вечером:"
			break
		}

		fmt.Println(daytime)
		if forcast[i].TEMPERATURE.Min == forcast[i].TEMPERATURE.Max {
			fmt.Println(forcast[i].TEMPERATURE.Min + "°C")
		} else {
			fmt.Println("От " + forcast[i].TEMPERATURE.Min + "°C до " + forcast[i].TEMPERATURE.Max + "°C")
		}
		if forcast[i].HEAT.Min == forcast[i].HEAT.Max {
			fmt.Println("Ощущается как " + forcast[i].HEAT.Min + "°C")
		} else {
			fmt.Println("Ощущается как " + forcast[i].HEAT.Min + "°C - " + forcast[i].HEAT.Max + "°C")
		}
	}
}
