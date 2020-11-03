package station

type Station struct {
	id              string
	adr             string
	city            string
	total_slots     int
	available_bikes int
	free_slots      int
}

func (s Station) String() string{
  return fmt.Sprintf("\nId: %v\n Adress: %v\n total_slots: %v\n Bikes: %v\n Empty slots: %v\n", s.id, s.adr, s.total_slots, s.available_bikes, s.free_slots)
}

func (s Station) bool check_free_slots(){
  result := false
  if s.free_slots > 0 {
    result = true
  }
  return result
}

func (s Station) bool check_bikes(){
  result := false
  if s.available_bikes > 0 {
    result = true
  }
  return result
}
