package city

import (
  "fmt"
)

type Coordinate [2]float64

type City struct{
  id string
  name string
  stations_num int
  var station_pos map[string]Coordinate
}
