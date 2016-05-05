package main

import (
"fmt"
//"errors"
"strings"
)

func main() {
  displayMenu()

  var option string
  fmt.Scanln(&option)
  selectOption(option)
}


func displayMenu() {
  fmt.Println("1) Generate Power Plant Report")
  fmt.Println("2) Generate Power Grid Report")
  fmt.Println("Please select an option: ")
}

func selectOption(option string) {
  _, gridLoad := getTestData()
  switch option {
  case "1":
    gridLoad.generatePlantReport()
  case "2":
    gridLoad.generateGridReport()
  default:
    fmt.Println("Invalid option selected")
 }
}

func generatePlantCapacityReport(plantCapacities ...float64) {
  for index, cap := range plantCapacities {
    fmt.Println("Plant Capacity: ", index, cap)
  }
}

func getTestData()(plants []PowerPlant, gridLoad PowerGrid){
  plants = []PowerPlant {
    PowerPlant{hydro, 300, active},
    PowerPlant{wind, 30, active},
    PowerPlant{wind, 25, inactive},
    PowerPlant{wind, 35, active},
    PowerPlant{solar, 45, unavailable},
    PowerPlant{solar, 40, inactive},
  }

  gridLoad = PowerGrid{300, plants}

  return
}

type PlantType string

const (
  hydro PlantType = "Hydro"
  wind PlantType = "Wind"
  solar PlantType = "Solar"
)

type PlantStatus string

const(
  active PlantStatus = "Active"
   inactive PlantStatus = "Inactive"
  unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
  plantType PlantType
  capacity float64
  status PlantStatus
  }

  type PowerGrid struct {
    load float64
    plants []PowerPlant
  }

  func (pg * PowerGrid) generatePlantReport() {
    for index, plant := range pg.plants {
      label := fmt.Sprintf("","Plant #", index)
      fmt.Println(plant)
      fmt.Println(strings.Repeat("-", len(label)))
      fmt.Println()
    }
  }

  func (pg * PowerGrid) generateGridReport() {
    capacity := 0.
    for _, plant := range pg.plants {
      if plant.status == active{
        capacity += plant.capacity
      }
    }
    label := "Power Grid Report"
    fmt.Println(label)
    fmt.Println("",strings.Repeat("-", len(label)))
    fmt.Println("Capacity", capacity)
    fmt.Println("Load", pg.load)
    fmt.Println("Utilization", pg.load/capacity*100)
  }
