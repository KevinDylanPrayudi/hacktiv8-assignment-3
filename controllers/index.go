package controllers

import (
	"assignment-3/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"script": helpers.Script(`
		const water = document.querySelectorAll("span.text-stone-950.text-center.font-display.block>b")[0];
		const wind = document.querySelectorAll("span.text-stone-950.text-center.font-display.block>b")[1];
		let notifWater, notifWind;
		setInterval(function(){
			fetch("/ping")
			.then(response => {
				if (!response.ok) {
				throw new Error('Network response was not ok');
				}
				return response.json();
			})
			.then(data => {
				if (data.data.Status.Water > 8) {
					water.classList.add("text-red-600");
					notifWater = "danger";
				} else if (data.data.Status.Water >= 6 && data.data.Status.Water <= 8) {
					water.classList.remove("text-red-600");
					notifWater = "Standby";
				} else {
					water.classList.remove("text-red-600");
					notifWater = "Safe";
				}
				if (data.data.Status.Wind > 15) {
					wind.classList.add("text-red-600");
					notifWind = "danger";
				} else if (data.data.Status.Wind >= 7 && data.data.Status.Wind <= 15) {
					wind.classList.remove("text-red-600");
					notifWind = "Standby";
				} else {
					wind.classList.remove("text-red-600");
					notifWind = "Safe";
				}
				water.innerText = data.data.Status.Water + "m " + notifWater;
				wind.innerText = data.data.Status.Wind + "m/s " + notifWind;
				console.log(data.data.Status)
			})
			.catch(error => {
				console.error('Error:', error);
			});
		 }, 15000);`),
	})
}
