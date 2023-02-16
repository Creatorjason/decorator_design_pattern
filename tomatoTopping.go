package main


type tomatoTopping struct{
	pizza pizza
}

func (tp *tomatoTopping) getPrice() int{
	pizzaPrice := tp.pizza.getPrice()
	return pizzaPrice + 12
}

