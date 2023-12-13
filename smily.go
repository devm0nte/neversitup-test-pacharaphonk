package main

// * Show coding logic and unit test Count the smiley faces! (Choose your the most skilled)
// * Given an array (arr) as an argument complete the function countSmileys that should return the total number of smiling faces.

/*
Rules for a smiling face:
- Each smiley face must contain a valid pair of eyes. Eyes can be marked as : or ;
- A smiley face can have a nose but it does not have to. Valid characters for a nose are - or ~
- Every smiling face must have a smiling mouth that should be marked with either ) or D

No additional characters are allowed except for those mentioned.

Valid smiley face examples:  :) :D ;-D :~)
Invalid smiley faces:  ;( :> :} :]
*/
type theFace struct {
	Eye   string
	Nose  string
	Mouth string
}

var validSmilyEye = []string{":", ";"}
var validSmilyNose = []string{"-", "~"}
var validSmilyMouth = []string{")", "D"}

func CountSmilyFace(text []string) int {
	if text == nil {
		return 0
	}
	var result int
	for _, v := range text {
		face, ok := GetTheFace(v)

		if ok {
			if IsValidSmilyFace(face) {
				result++
			}
		}

	}
	return result
}

func GetTheFace(face string) (theFace, bool) {
	// eye must be face[0]
	// if len(face != 3) nose doesn't exist, can assume mouth should be face[1]
	// if full face (len(face) == 3) face will be : eye = face[0]; nose = face[1]; and mouth = face[2];

	switch faceSize := len(face); faceSize {
	case 2:
		// fmt.Println("THIS IS 2 PART of FACE")
		return theFace{Eye: string(face[0]), Mouth: string(face[1])}, true

	case 3:
		// fmt.Println("THIS IS 3 PART of FACE")
		return theFace{Eye: string(face[0]), Nose: string(face[1]), Mouth: string(face[2])}, true
	default:
		return theFace{}, false
	}
}

func IsValidSmilyFace(validFace theFace) bool {
	return Contains(validSmilyEye, validFace.Eye) &&
		Contains(validSmilyMouth, validFace.Mouth) &&
		((validFace.Nose == "") || Contains(validSmilyNose, validFace.Nose))
}

func Contains(slice []string, t string) bool {
	for _, v := range slice {
		if v == t {
			return true
		}
	}
	return false
}
