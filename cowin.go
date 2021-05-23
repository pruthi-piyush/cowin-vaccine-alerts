package main

import (
  "fmt"
  "os"
  "flag"
  "time"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gen2brain/beeep"
)

type Filters struct {
  age int
  vaccine string
  pin int
  dose int
}

var filters Filters
   
func init() {

  flag.IntVar(&filters.age, "age", 18, `Age - Enter 18 if you want 18+ alerts, 45 if you want 45+ alerts.
  If you don't enter any value - default will be 18`)
  flag.StringVar(&filters.vaccine, "vaccine", "", `Vaccine ( Can be either COVISHIELD or COVAXIN ). 
  You can leave this empty if you want alerts for all the vaccines`)
  flag.IntVar(&filters.pin, "pincode", 0, "Pin Code")
  flag.IntVar(&filters.dose, "dose", 1, "Dose No. for which you would like the alerts - Default is Dose 1")
  
  flag.Parse()

}


func validateFilters() {

  if filters.pin == 0 {
    fmt.Println("Please provide a pincode for which you would like to find a vaccine using --pincode flag")
    os.Exit(0)
  }

  if filters.dose != 1 && filters.dose != 2 {
    fmt.Println("Invalid dose : can be only 1 or 2")
    os.Exit(0)
  }

  if filters.vaccine != "" {
    if filters.vaccine != "COVISHIELD" && filters.vaccine != "COVAXIN" {
      fmt.Println("Vaccine must be either COVISHIELD or COVAXIN")
      os.Exit(0)
    }
  }

  if filters.age != 18 && filters.age != 45 {
    fmt.Println("Age can be either 18 or 45")
    os.Exit(0)
  }

}

func main() {

  fmt.Println("Hello ! Please Stay Safe for Yourself and your FAMILY!")
  fmt.Println("Filters Provided ", filters)
 
  validateFilters()

  for {
    fmt.Println("------------------------------------------\n\n")
    request()
    // to honor rate limit
    time.Sleep(6 * time.Second)
  }
}

func request() {

  // get date in dd-mm-yyyy format
  date := time.Now().Format("02-01-2006")
  fmt.Println("****************************")
  fmt.Println("Today's Date => ", date)
  fmt.Println("****************************")

  url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/calendarByPin?pincode=%v&date=%v", 
                      filters.pin, date)
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Add("accept", "application/json")
  req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36 Edg/90.0.818.51")

  client := &http.Client{}
  resp, err := client.Do(req)
	if err != nil {
    fmt.Println("ERROR - in making request : ", err)
		return
	}

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

  if resp.StatusCode != http.StatusOK {
    fmt.Println("ERROR - API did not return 200(OK) response : will retry after some time")
    fmt.Println(resp)
    return
  }
  var i interface{}

  err = json.Unmarshal(body, &i)
  if err != nil {
    panic(err)
  }

  checkForSlotsAvailability(i)
}

func checkForSlotsAvailability(i interface{}) {

  centers := i.(map[string]interface{})
  centersArray := centers["centers"]
  var content string

  for _,center := range(centersArray.([]interface{})) {
    c := center.(map[string]interface{})

    sessions := c["sessions"].([]interface{})
    for _, sess := range(sessions) {
      session := sess.(map[string]interface{})
      ageLimit := session["min_age_limit"].(float64)
      vaccine := session["vaccine"]

      // discard undesired alerts

      if int(ageLimit) != filters.age {
        continue
      }

      if filters.vaccine != "" && vaccine != filters.vaccine {
        continue
      }

      date := session["date"]
      availableDoses := session["available_capacity"].(float64)
      availableDose1 := session["available_capacity_dose1"].(float64)
      availableDose2 := session["available_capacity_dose2"].(float64)

      fmt.Println(c["address"])
      fmt.Println(c["name"])
      fmt.Println("DATE => ", date)
      fmt.Println("VACCINE => ",  vaccine)
      fmt.Println("TOTAL DOSES =>", availableDoses)
      fmt.Println("DOSE 1 =>", availableDose1)
      fmt.Println("DOSE 2 =>", availableDose2)
      fmt.Println("AGE LIMIT =>", ageLimit)
      fmt.Println("***********************************")

      if availableDose1 > 0  && filters.dose == 1 {
        fmt.Println("***********************************")
        fmt.Println("***********************************")
        fmt.Println("DOSE 1 AVAILABILITY")
        fmt.Println("ALERTTTTTTT")
        fmt.Println("ALERTTTTTTT")
        fmt.Println("ALERTTTTTTT")
        content = fmt.Sprintf("Date: %v | Dose No: 1 | Age: %v+\nAvailable: %v %v\n%v-%v", 
        date, ageLimit, availableDose1, vaccine, c["name"], c["address"])
        alert(content)
        fmt.Println("***********************************")
        fmt.Println("***********************************")
      }

      if availableDose2 > 0  && filters.dose == 2 {
        fmt.Println("***********************************")
        fmt.Println("***********************************")
        fmt.Println("DOSE 2 AVAILABILITY")
        fmt.Println("ALERTTTTTTT")
        fmt.Println("ALERTTTTTTT")
        fmt.Println("ALERTTTTTTT")
        content = fmt.Sprintf("Date: %v | Dose No: 2 | Age: %v+\nAvailable: %v %v\n%v-%v",
        date, ageLimit, availableDose2, vaccine, c["name"], c["address"])
        alert(content)
        fmt.Println("***********************************")
        fmt.Println("***********************************")
      }
    }
}

}

func alert(content string) {

  err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
  if err != nil {
    panic(err)
  }

  fmt.Println(content)
  err = beeep.Notify("Vaccine Available", content, "")
  if err != nil {
    panic(err)
  }

  err = beeep.Alert("Vaccine Available", content,  "")
  if err != nil {
    panic(err)
  }

}


