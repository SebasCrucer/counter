package protocol

type PROTOCOLCODE string

const GENDCODE PROTOCOLCODE = "::::GEND::::"


const WINFOCODE PROTOCOLCODE = "::::WI%d::::"

type WCODE int

const (
	_ WCODE = iota
	WOK
	WGENDCODE
	WERROR
)


const RCOUNT PROTOCOLCODE = "::::RC%d::::"

type RCODE int

const (
	_ RCODE = iota
	ROK
	RERROR
)