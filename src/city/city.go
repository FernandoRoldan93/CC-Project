package city

import (
  "station"
)

type Coordinate [2]float64

type City struct{
  id string
  name string
  stations_num int
  stationPos map[station.Station]Coordinate
}

func getStationsByCity(s string) /*map[Station]Coordinate*/{


}

func getStationCoords(s station.Station) /*Coordinate*/{

}
