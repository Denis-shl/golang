package main

import (
	"fmt"
)

/*
// not Mytex

type data struct {
	name string
}
func main (){
	maps := map[string]int{"india" : 10, "Russian" : 30, "USA" : 5}
	sl := make([]int, len(maps))

	for _, number := range maps {
		sl = append(sl, number)
	}

	fmt.Println(len(sl))
	sort.Ints(sl)
	fmt.Println(len(sl))

	for _,num := range sl{
		fmt.Println(num)
	}

}
*/
//-----------------------
//interface
type Payer interface {
	Pay(int) error
}

//------------------
// struct

type Wallet struct {
	name string
	many int
}

type Card struct {
	Balance    int
	ValidUntil string
	CardHolder string
	CVV        string
	number     string
}

type AppleId struct {
	Money   int
	AppleId string
}

//-----------------
// methods
func (s *Wallet) Pay(sum int) error {
	if sum > s.many {
		fmt.Errorf("not many")
	}
	s.many = s.many - sum
	return nil
}

func (s *Card) Pay(sum int) error {
	if sum > s.Balance {
		return fmt.Errorf("not many")
	}
	s.Balance = s.Balance - sum
	return nil
}

func (s *AppleId)Pay(sum int)error{
	if sum > s.Money {
		return fmt.Errorf("not many")
	}
	s.Money = s.Money - sum
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		fmt.Println("Ошибка.Не хватает денег")
		return
	}
	fmt.Println("good buy")
}

func main() {
	myWallet := &Wallet{"denis", 100}
	Buy(myWallet)


	var pay Payer

	pay = &Card{Balance	: 11, CardHolder: "denis"}
	Buy(pay)

	my := &AppleId{9, "denis"}
	Buy(my)
	fmt.Println(my.Money)
}
