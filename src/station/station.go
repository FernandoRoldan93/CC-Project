package station

import "fmt"

type Station struct {
	id             int
	adr            string
	city           string
	totalSlots     int
	availableBikes int
	freeSlots      int
}

func constructor(id int, adr string, city string, totalSlots int) *Station {
	s := new(Station)
	s.id = id
	s.adr = adr
	s.city = city
	s.totalSlots = totalSlots
	s.availableBikes = 0 //at creation, a station will be out of bikes
	s.freeSlots = totalSlots
	return s
}

func (s Station) String() string {
	return fmt.Sprintf("\nId: %v\n Adress: %v\n total_slots: %v\n Bikes: %v\n Empty slots: %v\n", s.id, s.adr, s.totalSlots, s.availableBikes, s.freeSlots)
}

func checkFreeSlots(s Station) bool {
	// result := false
	// if s.freeSlots > 0 {
	// 	result = true
	// }
	// return result
	return false
}

func checkBikes(s Station) bool {
	// result := false
	// if s.availableBikes > 0 {
	// 	result = true
	// }
	// return result
	return false
}

func rentBike(s Station) string {
	// result := ""
	// if checkBikes(s) == true {
	// 	s.availableBikes--
	// 	s.freeSlots++
	// 	result = "1: ok "
	// } else {
	// 	result = "0: ERROR, There are not any bike to rent"
	// }
	// return result
	return ""
}

func storeBike(s Station) string {
	// result := ""
	// if checkFreeSlots(s) == true {
	// 	s.availableBikes++
	// 	s.freeSlots--
	// 	result = "1: ok "
	// } else {
	// 	result = "0: ERROR, There are not any free slot"
	// }
	// return result
	return ""
}
