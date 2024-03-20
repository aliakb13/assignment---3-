package handler

import (
	"assignment3/models"
	"assignment3/util"
	"fmt"
	"net/http"
	"text/template"
)

func GetWeather(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		water, wind := util.GenerateRandomValue()
		// water, wind := 9, 16
		wtrStatus, wndStatus := util.GenerateStatus(water, wind)

		res := models.Response{
			Water:       water,
			Wind:        wind,
			WaterStatus: wtrStatus,
			WindStatus:  wndStatus,
		}

		template, err := template.ParseFiles("templates/template.html")

		if err != nil {
			fmt.Println(err)
			return
		}

		template.Execute(w, res)

		// data, err := toJson(res)
		// if err != nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte(err.Error()))
		// }

		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not found of route"))
	}
}
