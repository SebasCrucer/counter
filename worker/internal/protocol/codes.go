package protocol


const GENDCODE string = "::::END::::"


const WINFOCODE string = "::::WI%d::::"

type WCODE int

const (
	_ WCODE = iota
	WOK
	WGENDCODE
	WERROR
)


const RCOUNT string = "::::RC%d::::"

type RCODE int

const (
	_ RCODE = iota
	ROK
	RERROR
)