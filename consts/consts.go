package consts

import "errors"

type FacesType int

const (
	_                  FacesType = iota
	FacesBomb                    = 1
	FacesSingle                  = 2
	FacesDouble                  = 3
	FacesTriple                  = 4
	FacesUnion                   = 5
	FacesStraight                = 6
	FacesUnionStraight           = 7
	FacesInvalid                 = 8
)

var (
	ErrsInvalidFaces = errors.New("Invalid faces. ")
)
