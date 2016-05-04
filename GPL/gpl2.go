package main

import "fmt"

func main() {
  displayMenu()
}


func displayMenu() {
  fmt.Println("1) Generate Power Plant Report")
  fmt.Println("2) Generate Power Grid Report")
  fmt.Println("Please select an option: ")

  var option string

  fmt.Scanln(&option)

  selectOption(option)
}

func selectOption(option string) {
  plantCapacities, activePlants, gridLoad := getTestData()

  switch option {
  case "1":
    generatePlantCapacityReport(plantCapacities...)
  case "2":
    generatePowerGridReport(activePlants, plantCapacities, gridLoad)
  default:
    fmt.Println("Invalid option selected")
}

}

func generatePlantCapacityReport(plantCapacities ...float64) {
  for index, cap := range plantCapacities {
    fmt.Println("Plant Capacity: ", index, cap)

  }
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64){
  capacity := 0.
  for _, plantId := range activePlants {
    capacity += plantCapacities[plantId]
  }
  fmt.Println("Capacity: ", capacity)
  fmt.Println("Load: ", gridLoad)
  fmt.Println("utilization: ", gridLoad/capacity*100)
}

func getTestData()(plantCapacities []float64, activePlants []int, gridLoad float64){

      plantCapacities = []float64{30, 30, 30, 60, 60, 100}

      activePlants = []int{0, 1}

      gridLoad = 75.

      return
}
