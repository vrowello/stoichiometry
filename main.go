package main

import (
  "net/http"
  "html/template"
  "strconv"
  "fmt"
  "os"
  "log"
)

var (
  Test1 float64
  Test2 float64
  GC float64
  GD float64
  LR float64
  ER float64
)

func main() {

  http.HandleFunc("/", stoichiometry)
  http.ListenAndServe(GetPort(), nil)
}

/*type InputData struct {
	MMA float64
	MMB float64
	MMC float64
  MMD float64
  NA float64
  NB float64
  NC float64
  ND float64
  GA float64
  GB float64
} */

type Stoichiometry struct {
	Message string
	Ans  float64
}

type StoichiometryData struct {
  Stoics []Stoichiometry
  Success bool
}

func stoichiometry(w http.ResponseWriter, r *http.Request) {
  page := template.Must(template.ParseFiles("stoichiometry.html"))
  if r.Method != http.MethodPost {
    page.Execute(w, nil)
      return
  }
  r.ParseForm()
  /*input := InputData{
    MMA:   strconv.ParseFloat(r.FormValue("MMA"), 64),
  	MMB:   r.FormValue("MMB"),
    MMC:   r.FormValue("MMC"),
    MMD:   r.FormValue("MMD"),
    NA:   r.FormValue("NA"),
    NB:   r.FormValue("NB"),
    NC:   r.FormValue("NC"),
    ND:   r.FormValue("ND"),
    GA:   r.FormValue("GA"),
    GB:   r.FormValue("GB"),
	} */

  MMA, err := strconv.ParseFloat(r.FormValue("MMA"), 64)
  if err != nil {
    log.Fatal(err)
  }
  
  MMB, err := strconv.ParseFloat(r.FormValue("MMB"), 64)
  if err != nil {
    log.Fatal(err)
  }

  MMC, err := strconv.ParseFloat(r.FormValue("MMC"), 64)
  if err != nil {
    log.Fatal(err)
  }

  MMD, err := strconv.ParseFloat(r.FormValue("MMD"), 64)
  if err != nil {
    log.Fatal(err)
  }

  NA, err := strconv.ParseFloat(r.FormValue("NA"), 64)
  if err != nil {
    log.Fatal(err)
  }

  NB, err := strconv.ParseFloat(r.FormValue("NB"), 64)
  if err != nil {
    log.Fatal(err)
  }

  NC, err := strconv.ParseFloat(r.FormValue("NC"), 64)
  if err != nil {
    log.Fatal(err)
  }

  ND, err := strconv.ParseFloat(r.FormValue("ND"), 64)
  if err != nil {
    log.Fatal(err)
  }

  GA, err := strconv.ParseFloat(r.FormValue("GA"), 64)
  if err != nil {
    log.Fatal(err)
  }

  GB, err := strconv.ParseFloat(r.FormValue("GB"), 64)
  if err != nil {
    log.Fatal(err)
  }

  Test1 = (((GA / MMA) * (NC / NA)) * MMC)
  Test2 = (((GB / MMB) * (NC / NB)) * MMC)
  if Test1 < Test2 {
    LR = 1
    ER = 2
    GC = Test1
    GD = (((GA / MMA) * (ND/ NA)) * MMD)
  } else if Test2 < Test1 {
    LR = 2
    ER = 1
    GC = Test2
    GD = (((GB / MMB) * (ND/ NB)) * MMD)
  } else {
    LR = 0
    ER = 0
    GC = Test1
    GD = (((GA / MMA) * (ND/ NA)) * MMD)
  }


  output := StoichiometryData{
    Stoics: []Stoichiometry{
		  {Message: "Grams of C Produced", Ans: GC},
			{Message: "Grams of D Produced", Ans: GD},
			{Message: "Limiting Reagant", Ans: LR},
      {Message: "Excess Reagent", Ans: ER},
		},
    Success: true,
  }

  page.Execute(w, output)
}

func GetPort() string {
	var port = os.Getenv("PORT")
 	// Set a default port if there is nothing in the environment
 	if port == "" {
 		port = "4747"
 		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
 	}
 	return ":" + port
}
