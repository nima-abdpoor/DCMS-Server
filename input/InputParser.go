package input

import (
	db "DCMS/db/sqlc"
	"DCMS/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

func ReadFromFile() (configs []db.AddConfigTxParams, err error) {
	byteValue, err := ioutil.ReadFile("./public/single/config.txt")
	if err != nil {
		log.Fatal("error in Reading the config file...", err)
		return
	}
	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal("error in parsing the config file...", err)
		return
	}
	index := 0
	for _, cluster := range config.Cluster {
		firstUrls := util.MapStringArrayToHashArray(cluster.FirstURL, func(i string) uint32 {
			return util.GenerateCR32(i)
		})
		secondUrl := make([]db.UrlSecondTx, len(cluster.SecondURL))
		if len(cluster.Ids) != 0 {
			for _, id := range cluster.Ids {
				for i, url := range cluster.SecondURL {
					secondUrl[i] = db.UrlSecondTx{
						UniqueID: id,
						UrlHash:  strconv.Itoa(int(util.GenerateCR32(url.URL))),
						Regex:    url.Regex,
					}
				}
				configs = append(configs, db.AddConfigTxParams{
					ID:           id,
					SyncType:     cluster.SyncType,
					IsLive:       cluster.IsLive,
					SaveRequest:  cluster.SaveRequest,
					SaveResponse: cluster.SaveResponse,
					SaveError:    cluster.SaveError,
					SaveSuccess:  cluster.SaveSuccess,
					UrlHashFirst: firstUrls,
					UrlSecond:    secondUrl,
					RequestUrl:   cluster.RequestURL,
				})
				index++
			}
		} else {
			for i, url := range cluster.SecondURL {
				secondUrl[i] = db.UrlSecondTx{
					UniqueID: cluster.Name,
					UrlHash:  strconv.Itoa(int(util.GenerateCR32(url.URL))),
					Regex:    url.Regex,
				}
			}
			configs[0] = db.AddConfigTxParams{
				ID:           cluster.Name,
				SyncType:     cluster.SyncType,
				IsLive:       cluster.IsLive,
				SaveRequest:  cluster.SaveRequest,
				SaveResponse: cluster.SaveResponse,
				SaveError:    cluster.SaveError,
				SaveSuccess:  cluster.SaveSuccess,
				UrlHashFirst: firstUrls,
				UrlSecond:    secondUrl,
				RequestUrl:   cluster.RequestURL,
			}
		}
	}
	return
}

type Config struct {
	Cluster []Cluster `json:"cluster"`
}

type Cluster struct {
	Name         int64       `json:"name"`
	Ids          []int64     `json:"ids"`
	SyncType     string      `json:"syncType"`
	IsLive       bool        `json:"isLive"`
	SaveRequest  bool        `json:"saveRequest"`
	SaveResponse bool        `json:"saveResponse"`
	SaveError    bool        `json:"saveError"`
	SaveSuccess  bool        `json:"saveSuccess"`
	FirstURL     []string    `json:"firstUrl"`
	RequestURL   []string    `json:"requestUrl"`
	SecondURL    []SecondUrl `json:"secondUrl"`
}

type SecondUrl struct {
	URL   string     `json:"url"`
	Regex []db.Regex `json:"regex"`
}