package main


type chesseTopping struct{
	pizza pizza
}

func (ct *chesseTopping) getPrice() int{
	pizzaPrice := ct.pizza.getPrice()
	return pizzaPrice + 10
}

