package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
    return k - 273.15
}

func celsiusToKelvin(c float64) float64 {
    return c + 273.15
}

func main() {
    k := 273.15
    c := kelvinToCelsius(k)
    fmt.Printf("%.2f Kelvin = %.2f Celsius\n", k, c)

    c = 0
    k = celsiusToKelvin(c)
    fmt.Printf("%.2f Celsius = %.2f Kelvin\n", c, k)
}
