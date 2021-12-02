package functions

func increaseMagicCount(i illusion) {
	i.magicianCount++
}

/*
this function is a method with a receiver argument. This function is exactly same like function increaseMagicCount(i illusion)
Parameters for these functions are passed by value not by reference
*/
func (i illusion) increaseMagicianCount() {
	i.magicianCount++
}
