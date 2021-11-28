package consts

import "errors"

type FacesType int

const (
	_ FacesType = iota
	FacesBomb
	FacesSingle
	FacesDouble
	FacesTriple
	FacesUnion
	FacesStraight
	FacesUnionStraight
	FacesInvalid
)

var (
	ErrsInvalidFaces = errors.New("Invalid faces. ")
)
