package functions

/*
golang does not allow  to have same function name increaseMagicianCount because a function declared with pointer receiver
and a function with value receiver have same signature (input/output)
*/
func (i *illusion) increaseMagicianCountRef() {
	i.magicianCount++
}

func increaseMagicianCountRef(i *illusion) {
	i.magicianCount++
}
