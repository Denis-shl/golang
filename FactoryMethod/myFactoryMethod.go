package factory_method

import "log"

type Creator struct{
}

func NewCreator () *Creator{
	return &Creator{}
}

func (p *Creator)Create(name string)iCompany{

	var company iCompany
	
	switch name{
		case "Inmarco":
			company = &Inmarco{"Inmarco"}
		case "Russian Cold":
			company = &RussianCold{"Russian Cold"}
		default:
			log.Fatal("Unknown Action")
	}
	return company
}