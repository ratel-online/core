package consts

import "errors"

const IS = "INTERACTIVE_SIGNAL"

type FacesType int

const (
	_                   FacesType = iota
	FacesBomb                     = 1
	FacesSingle                   = 2
	FacesDouble                   = 3
	FacesTriple                   = 4
	FacesUnion3                   = 5
	FacesUnion4                   = 6
	FacesStraight                 = 7
	FacesUnion3Straight           = 8
	FacesInvalid                  = 9
)

var (
	ErrsInvalidFaces = errors.New("Invalid faces. ")
)
