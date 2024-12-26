package helpers

type symbol interface {
	getSymbol() rune
}

type identifiable interface {
	getId() Id
}
