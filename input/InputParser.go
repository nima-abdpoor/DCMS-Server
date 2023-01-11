package input

import (
	db2 "DCMS/db/postgresql/sqlc"
	"DCMS/util"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

func ReadFromFile() (configs []db2.AddConfigTxParams, err error) {
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
		secondUrl := make([]db2.UrlSecondTx, len(cluster.SecondURL))
		if len(cluster.Ids) != 0 {
			for _, id := range cluster.Ids {
				for i, url := range cluster.SecondURL {
					secondUrl[i] = db2.UrlSecondTx{
						UniqueID: id,
						UrlHash:  strconv.Itoa(int(util.GenerateCR32(url.URL))),
						Regex:    url.Regex,
					}
				}
				configs = append(configs, db2.AddConfigTxParams{
					ID:                     id,
					NetworkType:            cluster.NetworkType,
					IsLive:                 cluster.IsLive,
					SaveRequest:            cluster.SaveRequest,
					SaveResponse:           cluster.SaveResponse,
					SaveError:              cluster.SaveError,
					SaveSuccess:            cluster.SaveSuccess,
					RepeatInterval:         cluster.RepeatInterval,
					RepeatIntervalTimeUnit: cluster.RepeatIntervalTimeUnit,
					RequiresBatteryNotLow:  cluster.RequiresBatteryNotLow,
					RequiresStorageNotLow:  cluster.RequiresStorageNotLow,
					RequiresCharging:       cluster.RequiresCharging,
					RequiresDeviceIdl:      cluster.RequiresDeviceIdl,
					UrlHashFirst:           firstUrls,
					UrlSecond:              secondUrl,
					RequestUrl:             cluster.RequestURL,
				})
				index++
			}
		} else {
			for i, url := range cluster.SecondURL {
				secondUrl[i] = db2.UrlSecondTx{
					UniqueID: cluster.Name,
					UrlHash:  strconv.Itoa(int(util.GenerateCR32(url.URL))),
					Regex:    url.Regex,
				}
			}
			configs = append(configs, db2.AddConfigTxParams{
				ID:                     cluster.Name,
				NetworkType:            cluster.NetworkType,
				IsLive:                 cluster.IsLive,
				SaveRequest:            cluster.SaveRequest,
				SaveResponse:           cluster.SaveResponse,
				SaveError:              cluster.SaveError,
				SaveSuccess:            cluster.SaveSuccess,
				RepeatInterval:         cluster.RepeatInterval,
				RepeatIntervalTimeUnit: cluster.RepeatIntervalTimeUnit,
				RequiresBatteryNotLow:  cluster.RequiresBatteryNotLow,
				RequiresStorageNotLow:  cluster.RequiresStorageNotLow,
				RequiresCharging:       cluster.RequiresCharging,
				RequiresDeviceIdl:      cluster.RequiresDeviceIdl,
				UrlHashFirst:           firstUrls,
				UrlSecond:              secondUrl,
				RequestUrl:             cluster.RequestURL,
			})
		}
	}
	return
}

type Config struct {
	Cluster []Cluster `json:"cluster"`
}

type Cluster struct {
	Name                   int64       `json:"name"`
	Ids                    []int64     `json:"ids"`
	NetworkType            string      `json:"networkType"`
	IsLive                 bool        `json:"isLive"`
	SaveRequest            bool        `json:"saveRequest"`
	SaveResponse           bool        `json:"saveResponse"`
	SaveError              bool        `json:"saveError"`
	SaveSuccess            bool        `json:"saveSuccess"`
	RepeatInterval         int64       `json:"repeatInterval"`
	RepeatIntervalTimeUnit string      `json:"repeatIntervalTimeUnit"`
	RequiresBatteryNotLow  bool        `json:"requiresBatteryNotLow"`
	RequiresStorageNotLow  bool        `json:"requiresStorageNotLow"`
	RequiresCharging       bool        `json:"requiresCharging"`
	RequiresDeviceIdl      bool        `json:"requiresDeviceIdl"`
	FirstURL               []string    `json:"firstUrl"`
	RequestURL             []string    `json:"requestUrl"`
	SecondURL              []SecondUrl `json:"secondUrl"`
}

type SecondUrl struct {
	URL   string      `json:"url"`
	Regex []db2.Regex `json:"regex"`
}
