package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/danoand/utils"
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

// RetStructAvgRent models the return object for each request
type RetStructAvgRent struct {
	Msg     string      `json:"msg"`
	Score   interface{} `json:"score"`
	AvgRent interface{} `json:"avg_rent"`
}

// RetStructWalk models the return object for each request
type RetStructWalk struct {
	Msg      string      `json:"msg"`
	Score    interface{} `json:"score"`
	RawScore interface{} `json:"raw_score"`
}

// RetStructGeneric models the return object for each request
type RetStructGeneric struct {
	Msg   string      `json:"msg"`
	Score interface{} `json:"score"`
}

func main() {
	var err error
	var fle *os.File
	var tmpRows [][]string
	var slcCities []interface{}
	var cols []string

	// Iterate through the city map and generate random values
	for key := range cityCodeMap {
		cityCodeMap[key] = rand.Intn(5) + 1
	}

	// Read city csv data
	log.Printf("INFO: reading in list of supported cities")
	fle, err = os.Open("city_list.csv")
	if err != nil {
		// error opening up the cities file
		log.Fatalf("FATAL: %v - error opening up the cities file - see: %v\n",
			utils.FileLine(),
			err)
	}

	// Set up a csv reader used to read csv data
	rdr := csv.NewReader(fle)

	// Read in all of the csv data
	rdr.FieldsPerRecord = 4
	tmpRows, err = rdr.ReadAll()
	if err != nil {
		// error reading in the csv file data
		log.Fatalf("FATAL: %v - error reading in the csv file data - see: %v\n",
			utils.FileLine(),
			err)
	}

	// Iterate through the read csv data and
	//   construct a slice of city maps
	for i := 0; i < len(tmpRows); i++ {
		var tmpMap = make(map[string]string)

		if i == 0 {
			// iterating over the column "headers"
			cols = tmpRows[i]
			continue
		}

		// iterating over a data row
		tmpMap[cols[0]] = tmpRows[i][0]
		tmpMap[cols[1]] = tmpRows[i][1]
		tmpMap[cols[2]] = tmpRows[i][2]
		tmpMap[cols[3]] = tmpRows[i][3]

		// apend the map of city data to our running slice of cities
		slcCities = append(slcCities, tmpMap)
	}
	log.Printf("INFO: finished processing the list of supported cities")

	r := gin.Default()

	// Return an array of officially supported cities
	r.GET("/cities", func(c *gin.Context) {
		c.JSON(200, slcCities)
	})

	r.GET("/rent_avg/:city", func(c *gin.Context) {
		var (
			ok     bool
			city   string
			retVal RetStructAvgRent
		)

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			retVal.Score = nil
			retVal.AvgRent = nil
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			retVal.Score = nil
			retVal.AvgRent = nil
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "average rent"
		retVal.Score = rand.Intn(5) + 1
		max := 3000
		min := 1500
		retVal.AvgRent = rand.Intn(max-min) + min
		c.JSON(200, retVal)
	})

	r.GET("/walk_scr/:city", func(c *gin.Context) {
		var (
			ok     bool
			city   string
			retVal RetStructWalk
		)

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			retVal.Score = nil
			retVal.RawScore = nil
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			retVal.Score = nil
			retVal.RawScore = nil
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "walkability score"
		retVal.Score = rand.Intn(5) + 1
		retVal.RawScore = rand.Intn(100) + 1
		c.JSON(200, retVal)
	})

	r.GET("/crime_scr/:city", func(c *gin.Context) {
		var (
			ok     bool
			city   string
			retVal RetStructGeneric
		)

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			retVal.Score = nil
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			retVal.Score = nil
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "crime score"
		retVal.Score = rand.Intn(5) + 1
		c.JSON(200, retVal)
	})

	r.GET("/air_qual_scr/:city", func(c *gin.Context) {
		var (
			ok     bool
			city   string
			retVal RetStructGeneric
		)

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			retVal.Score = nil
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			retVal.Score = nil
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "air quality score"
		retVal.Score = rand.Intn(5) + 1
		c.JSON(200, retVal)
	})

	r.GET("/city_scr/:city", func(c *gin.Context) {
		var (
			ok     bool
			city   string
			retVal RetStructGeneric
		)

		// Extract city parameter
		city = c.Param("city")
		if len(city) == 0 {
			retVal.Msg = "missing city parameter"
			retVal.Score = nil
			c.JSON(http.StatusBadRequest, retVal)
			return
		}

		// Lookup the city code
		_, ok = cityCodeMap[city]
		if !ok {
			// city code not found
			retVal.Msg = "city not found"
			retVal.Score = nil
			c.JSON(http.StatusNotFound, retVal)
			return
		}

		retVal.Msg = "overall quality of life score"
		retVal.Score = rand.Intn(5) + 1
		c.JSON(200, retVal)
	})

	log.Printf("INFO: starting web server on port: %v\n", 8999)

	r.Run(fmt.Sprintf(":%v", 8999))
}
