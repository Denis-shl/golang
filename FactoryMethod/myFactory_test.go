package factory_method

import "testing"

func TestFactoryMethod(t *testing.T) {

	 str := []string{"Inmarco", "Russian Cold"}

	factory := NewCreator()
	products := []iCompany{
		factory.Create(str[0]),
		factory.Create(str[1]),
	}

	for i, company := range products {
		if name := company.GettingNameCompany(); name != str[i]{
			t.Errorf("Expect action to %s, but %s.\n", company.GettingNameCompany(), str[i])
		}else {
			company.CreateIceCream()
		}
	}
}
