package service

import (
	"BannerMan-Server/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetWidgetsFromNpm() (error, []*model.Widgets) {
	resp, err := http.Get(viper.GetString("packages.api"))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	var res []*model.Widgets
	err = json.Unmarshal(body, &res)
	if err != nil {
		return err, nil
	}
	return nil, res
}