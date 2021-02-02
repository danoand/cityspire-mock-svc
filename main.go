package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var cityCodeMap = map[string]int{
	"New_York_City":    0,
	"Los_Angeles":      0,
	"Chicago":          0,
	"Houston":          0,
	"Phoenix":          0,
	"Philadelphia":     0,
	"San_Antonio":      0,
	"San_Diego":        0,
	"Dallas":           0,
	"Austin":           0,
	"San_Jose":         0,
	"Fort_Worth":       0,
	"Jacksonville":     0,
	"Columbus":         0,
	"Charlotte":        0,
	"Indianapolis":     0,
	"San_Francisco":    0,
	"Seattle":          0,
	"Denver":           0,
	"Washington":       0,
	"Boston":           0,
	"El_Paso":          0,
	"Nashville":        0,
	"Oklahoma_City":    0,
	"Las_Vegas":        0,
	"Detroit":          0,
	"Portland":         0,
	"Memphis":          0,
	"Louisville":       0,
	"Milwaukee":        0,
	"Baltimore":        0,
	"Albuquerque":      0,
	"Tucson":           0,
	"Mesa":             0,
	"Fresno":           0,
	"Sacramento":       0,
	"Atlanta":          0,
	"Kansas_City":      0,
	"Colorado_Springs": 0,
	"Raleigh":          0,
	"Omaha":            0,
	"Miami":            0,
	"Long_Beach":       0,
	"Virginia_Beach":   0,
	"Oakland":          0,
	"Minneapolis":      0,
	"Tampa":            0,
	"Tulsa":            0,
	"Arlington":        0,
	"Wichita":          0,
	"Bakersfield":      0,
	"Aurora":           0,
	"New_Orleans":      0,
	"Cleveland":        0,
	"Anaheim":          0,
	"Henderson":        0,
	"Honolulu":         0,
	"Riverside":        0,
	"Santa_Ana":        0,
	"Corpus_Christi":   0,
	"Lexington":        0,
	"San_Juan":         0,
	"Stockton":         0,
	"St_Paul":          0,
	"Cincinnati":       0,
	"Greensboro":       0,
	"Pittsburgh":       0,
	"Irvine":           0,
	"St_Louis":         0,
	"Lincoln":          0,
	"Orlando":          0,
	"Durham":           0,
	"Plano":            0,
	"Anchorage":        0,
	"Newark":           0,
	"Chula_Vista":      0,
	"Fort_Wayne":       0,
	"Chandler":         0,
	"Toledo":           0,
	"St_Petersburg":    0,
	"Reno":             0,
	"Laredo":           0,
	"Scottsdale":       0,
	"North_Las_Vegas":  0,
	"Lubbock":          0,
	"Madison":          0,
	"Gilbert":          0,
	"Jersey_City":      0,
	"Glendale":         0,
	"Buffalo":          0,
	"WinstonSalem":     0,
	"Chesapeake":       0,
	"Fremont":          0,
	"Norfolk":          0,
	"Irving":           0,
	"Garland":          0,
	"Paradise":         0,
	"Richmond":         0,
	"Hialeah":          0,
	"Boise":            0,
	"Spokane":          0,
	"Frisco":           0,
	"Moreno_Valley":    0,
	"Tacoma":           0,
	"Fontana":          0,
	"Modesto":          0,
	"Baton_Rouge":      0,
	"Port_St_Lucie":    0,
	"San_Bernardino":   0,
	"McKinney":         0,
	"Fayetteville":     0,
	"Santa_Clarita":    0,
	"Des_Moines":       0,
	"Oxnard":           0,
	"Birmingham":       0,
	"Spring_Valley":    0,
	"Huntsville":       0,
	"Rochester":        0,
	"Cape_Coral":       0,
	"Tempe":            0,
	"Grand_Rapids":     0,
	"Yonkers":          0,
	"Overland_Park":    0,
	"Salt_Lake_City":   0,
	"Amarillo":         0,
	"Augusta":          0,
	"Tallahassee":      0,
	"Montgomery":       0,
	"Huntington_Beach": 0,
	"Akron":            0,
	"Little_Rock":      0,
	"Grand_Prairie":    0,
	"Sunrise_Manor":    0,
	"Ontario":          0,
	"Sioux_Falls":      0,
	"Knoxville":        0,
	"Vancouver":        0,
	"Mobile":           0,
	"Worcester":        0,
	"Chattanooga":      0,
	"Brownsville":      0,
	"Peoria":           0,
	"Fort_Lauderdale":  0,
	"Shreveport":       0,
	"Newport_News":     0,
	"Providence":       0,
	"Elk_Grove":        0,
	"Rancho_Cucamonga": 0,
	"Salem":            0,
	"Pembroke_Pines":   0,
	"Santa_Rosa":       0,
	"Eugene":           0,
	"Oceanside":        0,
	"Cary":             0,
	"Fort_Collins":     0,
	"Corona":           0,
	"Enterprise":       0,
	"Garden_Grove":     0,
	"Springfield":      0,
	"Clarksville":      0,
	"Bayamon":          0,
	"Lakewood":         0,
	"Alexandria":       0,
	"Hayward":          0,
	"Murfreesboro":     0,
	"Killeen":          0,
	"Hollywood":        0,
	"Lancaster":        0,
	"Salinas":          0,
	"Jackson":          0,
	"Midland":          0,
	"Macon_County":     0,
	"Palmdale":         0,
	"Sunnyvale":        0,
	"Escondido":        0,
	"Pomona":           0,
	"Bellevue":         0,
	"Surprise":         0,
	"Naperville":       0,
	"Pasadena":         0,
	"Denton":           0,
	"Roseville":        0,
	"Joliet":           0,
	"Thornton":         0,
	"McAllen":          0,
	"Paterson":         0,
	"Rockford":         0,
	"Carrollton":       0,
	"Bridgeport":       0,
	"Miramar":          0,
	"Round_Rock":       0,
	"Metairie":         0,
	"Olathe":           0,
	"Waco":             0}

//RetStruct models the return object for each request
type RetStruct struct {
	Msg   string `json:"msg"`
	Score int    `json:"score"`
}

func main() {
	// Iterate through the city map and generate random values
	for key := range cityCodeMap {
		cityCodeMap[key] = rand.Intn(5) + 1
	}

	r := gin.Default()
	r.GET("/crime_scr/:city", func(c *gin.Context) {
		var ok bool
		var city string
		var retVal RetStruct
		var tmpScr int

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		tmpScr, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "crime score"
		retVal.Score = tmpScr
		c.JSON(200, retVal)
	})

	r.GET("/walk_scr/:city", func(c *gin.Context) {
		var ok bool
		var city string
		var retVal RetStruct
		var tmpScr int

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		tmpScr, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "walkability score"
		retVal.Score = tmpScr
		c.JSON(200, retVal)
	})

	r.GET("/rent_avg/:city", func(c *gin.Context) {
		var ok bool
		var city string
		var retVal RetStruct

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "average rent"
		max := 3000
		min := 1500
		retVal.Score = rand.Intn(max-min) + min
		c.JSON(200, retVal)
	})

	// Assume running on Heroku. Grab the appropriate port
	prt := os.Getenv("PORT")
	log.Printf("INFO: starting web server on port: %v\n", prt)

	r.Run(fmt.Sprintf(":%v", prt))
}
