package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type weatherData struct {
	Lat            float32 `json:"lat"`
	Lon            float32 `json:"lon"`
	TimeZone       string  `json:"timezone"`
	TimeZoneOffset int64   `json:"timezone_offset"`
	Current        struct {
		Sunrise    int64   `json:"sunrise"`
		Sunset     int64   `json:"sunset"`
		Temp       float32 `json:"temp"`
		FeelsLike  float32 `json:"feels_like"`
		Pressure   float32 `json:"pressure"`
		Humidity   float32 `json:"humidity"`
		DewPoint   float32 `json:"dew_point"`
		Uvi        float32 `json:"uvi"`
		Clouds     float32 `json:"clouds"`
		Visibility float32 `json:"visibility"`
		WindSpeed  float32 `json:"wind_speed"`
		WindDeg    float32 `json:"wind_deg"`
		WindGust   float32 `json:"wind_gust"`
		Weather    []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		}
	}
}

type geocodeData []struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func sessionRespond(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: message,
		},
	})
	if err != nil {
		log.Panicf("could not respond to interaction: %s", err)
	}
}

func getGeocode(query string, apiKey string) (string, string) {
	var geocodeDataRes geocodeData
	url := fmt.Sprintf(
		"http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s",
		query,
		apiKey,
	)

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", ""
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return "", ""
	}

	body, err := io.ReadAll(res.Body)
	if err != nil || body == nil {
		log.Println(err)
		return "", ""
	}

	err = json.Unmarshal(body, &geocodeDataRes)
	if err != nil {
		log.Println(err)
		return "", ""
	}

	if len(geocodeDataRes) == 0 {
		return "", ""
	}

	return strconv.FormatFloat(float64(geocodeDataRes[0].Lat), 'f', -1, 32),
		strconv.FormatFloat(float64(geocodeDataRes[0].Lon), 'f', -1, 32)
}

func getWeather(query string, apiKey string) weatherData {
	var weatherDataRes weatherData
	lat, lon := getGeocode(query, apiKey)
	if lat == "" || lon == "" {
		return weatherData{}
	}

	url := fmt.Sprintf(
		"%s?units=%s&exclude=%s&lat=%s&lon=%s&appid=%s",
		"https://api.openweathermap.org/data/3.0/onecall",
		"metric",
		"minutely,hourly,daily,alerts",
		lat,
		lon,
		apiKey,
	)

	res, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return weatherData{}
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Println(err)
		return weatherData{}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil || body == nil {
		log.Println(err)
		return weatherData{}
	}

	err = json.Unmarshal(body, &weatherDataRes)
	if err != nil {
		log.Println(err)
		return weatherData{}
	}

	return weatherDataRes
}

func HandleWeahter(
	s *discordgo.Session,
	i *discordgo.InteractionCreate,
	opts map[string]*discordgo.ApplicationCommandInteractionDataOption,
	apiKey string,
) {
	builder := new(strings.Builder)
	query, ok := opts["query"]
	if !ok {
		return
	}

	weatherDataRes := getWeather(query.StringValue(), apiKey)
	if len(weatherDataRes.Current.Weather) == 0 {
		sessionRespond(s, i, "Bot made by syafahr\n\nfailed to get weather data\nPossible reason:\n\n- Invalid query\n- API limit")
		return
	}

	res := fmt.Sprintf(
		"Bot made by syafahr\n\n"+
			"Lat: %.2f\nLon: %.2f\nTime Zone: %s\nTime Zone Offset: %d\n"+
			"- Sunrise: %d\n- Sunset: %d\n- Temp: %.2f °C\n- Feels Like: %.2f °C\n"+
			"- Pressure: %.2f hPa\n- Humidity: %.2f%%\n- Dew Point: %.2f °C\n- UVI: %.2f\n"+
			"- Clouds: %.2f %%\n- Visibility: %.2f\n- Wind Speed: %.2f m/s\n- Wind Deg: %.2f\n"+
			"- Wind Gust: %.2f m/s\n  - Main: %s\n  - Description: %s\n",
		weatherDataRes.Lat,
		weatherDataRes.Lon,
		weatherDataRes.TimeZone,
		weatherDataRes.TimeZoneOffset,
		weatherDataRes.Current.Sunrise,
		weatherDataRes.Current.Sunset,
		weatherDataRes.Current.Temp,
		weatherDataRes.Current.FeelsLike,
		weatherDataRes.Current.Pressure,
		weatherDataRes.Current.Humidity,
		weatherDataRes.Current.DewPoint,
		weatherDataRes.Current.Uvi,
		weatherDataRes.Current.Clouds,
		weatherDataRes.Current.Visibility,
		weatherDataRes.Current.WindSpeed,
		weatherDataRes.Current.WindDeg,
		weatherDataRes.Current.WindGust,
		weatherDataRes.Current.Weather[0].Main,
		weatherDataRes.Current.Weather[0].Description,
	)

	builder.WriteString(res)

	sessionRespond(s, i, builder.String())
}
