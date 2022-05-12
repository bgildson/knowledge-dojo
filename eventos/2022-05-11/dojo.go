package dojo

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

type Address struct {
	ID       int
	City     string
	State    string
	PersonID int
}

type PersonAddress struct {
	FirstName string
	LastName  string
	City      string
	State     string
}

func CombineTwoTables1(person []Person, address []Address) []PersonAddress {
	res := []PersonAddress{}

	for _, p := range person {
		pa := PersonAddress{
			FirstName: p.FirstName,
			LastName:  p.LastName,
			City:      "Null",
			State:     "Null",
		}
		for _, a := range address {
			if a.PersonID == p.ID {
				pa.City = a.City
				pa.State = a.State
				break
			}
		}
		res = append(res, pa)
	}

	return res
}

func CombineTwoTables2(person []Person, address []Address) []PersonAddress {
	store := map[int]Address{}

	for i := range address {
		store[address[i].PersonID] = address[i]
	}

	var res []PersonAddress
	for _, p := range person {
		pa := PersonAddress{
			FirstName: p.FirstName,
			LastName:  p.LastName,
			City:      "Null",
			State:     "Null",
		}

		a, exists := store[p.ID]
		if exists {
			pa.City = a.City
			pa.State = a.State
		}

		res = append(res, pa)
	}

	return res
}
