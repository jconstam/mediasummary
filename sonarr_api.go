package mediasummary

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type sonarrSeriesData []struct {
	Title             string        `json:"title"`
	AlternateTitles   []interface{} `json:"alternateTitles"`
	SortTitle         string        `json:"sortTitle"`
	SeasonCount       int           `json:"seasonCount"`
	TotalEpisodeCount int           `json:"totalEpisodeCount"`
	EpisodeCount      int           `json:"episodeCount"`
	EpisodeFileCount  int           `json:"episodeFileCount"`
	SizeOnDisk        int64         `json:"sizeOnDisk"`
	Status            string        `json:"status"`
	Overview          string        `json:"overview,omitempty"`
	PreviousAiring    time.Time     `json:"previousAiring,omitempty"`
	Network           string        `json:"network,omitempty"`
	AirTime           string        `json:"airTime,omitempty"`
	Images            []struct {
		CoverType string `json:"coverType"`
		URL       string `json:"url"`
	} `json:"images"`
	Seasons []struct {
		SeasonNumber int  `json:"seasonNumber"`
		Monitored    bool `json:"monitored"`
		Statistics   struct {
			EpisodeFileCount  int     `json:"episodeFileCount"`
			EpisodeCount      int     `json:"episodeCount"`
			TotalEpisodeCount int     `json:"totalEpisodeCount"`
			SizeOnDisk        int     `json:"sizeOnDisk"`
			PercentOfEpisodes float64 `json:"percentOfEpisodes"`
		} `json:"statistics"`
	} `json:"seasons"`
	Year              int           `json:"year"`
	Path              string        `json:"path"`
	ProfileID         int           `json:"profileId"`
	SeasonFolder      bool          `json:"seasonFolder"`
	Monitored         bool          `json:"monitored"`
	UseSceneNumbering bool          `json:"useSceneNumbering"`
	Runtime           int           `json:"runtime"`
	TvdbID            int           `json:"tvdbId"`
	TvRageID          int           `json:"tvRageId"`
	TvMazeID          int           `json:"tvMazeId"`
	FirstAired        time.Time     `json:"firstAired,omitempty"`
	LastInfoSync      time.Time     `json:"lastInfoSync"`
	SeriesType        string        `json:"seriesType"`
	CleanTitle        string        `json:"cleanTitle"`
	ImdbID            string        `json:"imdbId,omitempty"`
	TitleSlug         string        `json:"titleSlug"`
	Certification     string        `json:"certification,omitempty"`
	Genres            []string      `json:"genres"`
	Tags              []interface{} `json:"tags"`
	Added             time.Time     `json:"added"`
	Ratings           struct {
		Votes int     `json:"votes"`
		Value float64 `json:"value"`
	} `json:"ratings"`
	QualityProfileID int       `json:"qualityProfileId"`
	ID               int       `json:"id"`
	NextAiring       time.Time `json:"nextAiring,omitempty"`
}

func getData(baseURI string, apiKey string) sonarrSeriesData {
	uri := baseURI + "/api/series?apikey=" + apiKey

	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		log.Printf("REQUEST FAILED: %v", err)
		return nil
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("DO FAILED: %v", err)
		return nil
	}

	defer resp.Body.Close()

	var data sonarrSeriesData

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Printf("DECODE FAILED: %v", err)
		return nil
	}

	return data
}

func printSeriesInfo(data sonarrSeriesData) {
	for _, value := range data {
		fmt.Printf("%v\n", value.Title)
	}
}
