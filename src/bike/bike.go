package bike

import (
  "station"
)

type Coordinate [2]float64

type Bike struct{
  id int
  lastUsers [10]string
  inUse bool
  stationId int
}

func init(id int, sId int) *Bike{
  bike := new(Bike)
  bike.id = id
  bike.stationId = sId
  bike.inUse = false
  return bike
}

func isInUse() /*bool*/{


}

func getLastUsers(s station.Station) /*map[Station]Coordinate*/{

}
