package pointer

type Wallet struct{
	balance float64
}

func (w *Wallet) Deposit(dep float64){
	w.balance+=dep
}

func (w *Wallet) Balance() float64{
	return w.balance
}