// Package enigma provides a simple library for German WWII Enigma encryption/decryption.
package enigma

import (
	"encoding/json"
	"io/ioutil"
	"unicode"
	"errors"
	"os"
)

// Rune represents a character to be encrypted/decrypted.
type Rune rune

// WheelTurnoverPoints represents the points at which the left sibling wheel
// should be stepped.
var WheelTurnoverPoints = map[string]map[Rune]struct{}{
	"I": {'Q': struct{}{}},
	"II": {'E': struct{}{}},
	"III": {'V': struct{}{}},
	"IV": {'J': struct{}{}},
	"V": {'Z': struct{}{}},
	"VI": {'Z': struct{}{}, 'M': struct{}{}},
	"VII": {'Z': struct{}{}, 'M': struct{}{}},
	"VIII": {'Z': struct{}{}, 'M': struct{}{}},
}

// ReflectorWiring represents the wiring configuration for each reflector rotor.
var ReflectorWiring = map[string]map[Rune]Rune{
	"A": {
		'A': 'E', 'B': 'J', 'C': 'M', 'D': 'Z', 'E': 'A', 'F': 'L',
		'G': 'Y', 'H': 'X', 'I': 'V', 'J': 'B', 'K': 'W', 'L': 'F',
		'M': 'C', 'N': 'R', 'O': 'Q', 'P': 'U', 'Q': 'O', 'R': 'N',
		'S': 'T', 'T': 'S', 'U': 'P', 'V': 'I', 'W': 'K', 'X': 'H',
		'Y': 'G', 'Z': 'D',
	},
	"B": {
		'A': 'Y', 'B': 'R', 'C': 'U', 'D': 'H', 'E': 'Q', 'F': 'S',
		'G': 'L', 'H': 'D', 'I': 'P', 'J': 'X', 'K': 'N', 'L': 'G',
		'M': 'O', 'N': 'K', 'O': 'M', 'P': 'I', 'Q': 'E', 'R': 'B',
		'S': 'F', 'T': 'Z', 'U': 'C', 'V': 'W', 'W': 'V', 'X': 'J',
		'Y': 'A', 'Z': 'T',
	},
	"C": {
		'A': 'F', 'B': 'V', 'C': 'P', 'D': 'J', 'E': 'I', 'F': 'A',
		'G': 'O', 'H': 'Y', 'I': 'E', 'J': 'D', 'K': 'R', 'L': 'Z',
		'M': 'X', 'N': 'W', 'O': 'G', 'P': 'C', 'Q': 'T', 'R': 'K',
		'S': 'U', 'T': 'Q', 'U': 'S', 'V': 'B', 'W': 'N', 'X': 'M',
		'Y': 'H', 'Z': 'L',
	},
	"B Thin": {
		'A': 'E', 'B': 'N', 'C': 'K', 'D': 'Q', 'E': 'A', 'F': 'U',
		'G': 'Y', 'H': 'W', 'I': 'J', 'J': 'I', 'K': 'C', 'L': 'O',
		'M': 'P', 'N': 'B', 'O': 'L', 'P': 'M', 'Q': 'D', 'R': 'X',
		'S': 'Z', 'T': 'V', 'U': 'F', 'V': 'T', 'W': 'H', 'X': 'R',
		'Y': 'G', 'Z': 'S',
	},
	"C Thin": {
		'A': 'R', 'B': 'D', 'C': 'O', 'D': 'B', 'E': 'J', 'F': 'N',
		'G': 'T', 'H': 'K', 'I': 'V', 'J': 'E', 'K': 'H', 'L': 'M',
		'M': 'L', 'N': 'F', 'O': 'C', 'P': 'W', 'Q': 'Z', 'R': 'A',
		'S': 'X', 'T': 'G', 'U': 'Y', 'V': 'I', 'W': 'P', 'X': 'S',
		'Y': 'U', 'Z': 'Q',
	},
}

// WheelWiring represents the wiring configuration for each rotor.
var WheelWiring = map[string]map[Rune]Rune{
	"I": {
		'A': 'E', 'B': 'K', 'C': 'M', 'D': 'F', 'E': 'L', 'F': 'G',
		'G': 'D', 'H': 'Q', 'I': 'V', 'J': 'Z', 'K': 'N', 'L': 'T',
		'M': 'O', 'N': 'W', 'O': 'Y', 'P': 'H', 'Q': 'X', 'R': 'U',
		'S': 'S', 'T': 'P', 'U': 'A', 'V': 'I', 'W': 'B', 'X': 'R',
		'Y': 'C', 'Z': 'J',
	},
	"II": {
		'A': 'A', 'B': 'J', 'C': 'D', 'D': 'K', 'E': 'S', 'F': 'I',
		'G': 'R', 'H': 'U', 'I': 'X', 'J': 'B', 'K': 'L', 'L': 'H',
		'M': 'W', 'N': 'T', 'O': 'M', 'P': 'C', 'Q': 'Q', 'R': 'G',
		'S': 'Z', 'T': 'N', 'U': 'P', 'V': 'Y', 'W': 'F', 'X': 'V',
		'Y': 'O', 'Z': 'E',
	},
	"III": {
		'A': 'B', 'B': 'D', 'C': 'F', 'D': 'H', 'E': 'J', 'F': 'L',
		'G': 'C', 'H': 'P', 'I': 'R', 'J': 'T', 'K': 'X', 'L': 'V',
		'M': 'Z', 'N': 'N', 'O': 'Y', 'P': 'E', 'Q': 'I', 'R': 'W',
		'S': 'G', 'T': 'A', 'U': 'K', 'V': 'M', 'W': 'U', 'X': 'S',
		'Y': 'Q', 'Z': 'O',
	},
	"IV": {
		'A': 'E', 'B': 'S', 'C': 'O', 'D': 'V', 'E': 'P', 'F': 'Z',
		'G': 'J', 'H': 'A', 'I': 'Y', 'J': 'Q', 'K': 'U', 'L': 'I',
		'M': 'R', 'N': 'H', 'O': 'X', 'P': 'L', 'Q': 'N', 'R': 'F',
		'S': 'T', 'T': 'G', 'U': 'K', 'V': 'D', 'W': 'C', 'X': 'M',
		'Y': 'W', 'Z': 'B',
	},
	"V": {
		'A': 'V', 'B': 'Z', 'C': 'B', 'D': 'R', 'E': 'G', 'F': 'I',
		'G': 'T', 'H': 'Y', 'I': 'U', 'J': 'P', 'K': 'S', 'L': 'D',
		'M': 'N', 'N': 'H', 'O': 'L', 'P': 'X', 'Q': 'A', 'R': 'W',
		'S': 'M', 'T': 'J', 'U': 'Q', 'V': 'O', 'W': 'F', 'X': 'E',
		'Y': 'C', 'Z': 'K',
	},
	"VI": {
		'A': 'J', 'B': 'P', 'C': 'G', 'D': 'V', 'E': 'O', 'F': 'U',
		'G': 'M', 'H': 'F', 'I': 'Y', 'J': 'Q', 'K': 'B', 'L': 'E',
		'M': 'N', 'N': 'H', 'O': 'Z', 'P': 'R', 'Q': 'D', 'R': 'K',
		'S': 'A', 'T': 'S', 'U': 'X', 'V': 'L', 'W': 'I', 'X': 'C',
		'Y': 'T', 'Z': 'W',
	},
	"VII": {
		'A': 'N', 'B': 'Z', 'C': 'J', 'D': 'H', 'E': 'G', 'F': 'R',
		'G': 'C', 'H': 'X', 'I': 'M', 'J': 'Y', 'K': 'S', 'L': 'W',
		'M': 'B', 'N': 'O', 'O': 'U', 'P': 'F', 'Q': 'A', 'R': 'I',
		'S': 'V', 'T': 'L', 'U': 'P', 'V': 'E', 'W': 'K', 'X': 'Q',
		'Y': 'D', 'Z': 'T',
	},
	"VIII": {
		'A': 'F', 'B': 'K', 'C': 'Q', 'D': 'H', 'E': 'T', 'F': 'L',
		'G': 'X', 'H': 'O', 'I': 'C', 'J': 'B', 'K': 'J', 'L': 'S',
		'M': 'P', 'N': 'D', 'O': 'Z', 'P': 'R', 'Q': 'A', 'R': 'M',
		'S': 'E', 'T': 'W', 'U': 'N', 'V': 'I', 'W': 'U', 'X': 'Y',
		'Y': 'G', 'Z': 'V',
	},
	"Beta": {
		'A': 'L', 'B': 'E', 'C': 'Y', 'D': 'J', 'E': 'V', 'F': 'C',
		'G': 'N', 'H': 'I', 'I': 'X', 'J': 'W', 'K': 'P', 'L': 'B',
		'M': 'Q', 'N': 'M', 'O': 'D', 'P': 'R', 'Q': 'T', 'R': 'A',
		'S': 'K', 'T': 'Z', 'U': 'G', 'V': 'F', 'W': 'U', 'X': 'H',
		'Y': 'O', 'Z': 'S',
	},
	"Gamma": {
		'A': 'F', 'B': 'S', 'C': 'O', 'D': 'K', 'E': 'A', 'F': 'N',
		'G': 'U', 'H': 'E', 'I': 'R', 'J': 'H', 'K': 'M', 'L': 'B',
		'M': 'T', 'N': 'I', 'O': 'Y', 'P': 'C', 'Q': 'W', 'R': 'L',
		'S': 'Q', 'T': 'P', 'U': 'Z', 'V': 'X', 'W': 'V', 'X': 'G',
		'Y': 'J', 'Z': 'D',
	},
}

// Wheel is an Enigma rotor.
type Wheel struct {
	Number string `json:"number"`
	RingSetting Rune `json:"ring"`
	GroundSetting Rune `json:"ground"`
}

// Enigma is the type encapsulating encryption/decryption methods.
type Enigma struct {
	Model string `json:"model"`
	Reflector string `json:"reflector"`
	Wheels []Wheel `json:"wheels"`
	Plugboard [][]Rune `json:"plugboard"`
}

// UnmarshalJSON unmarshals a json string representation of a letter into a Rune.
func (c *Rune) UnmarshalJSON(data []byte) (err error) {
	*c = Rune(data[1])

	return
}

// GetEntryContact gets the entry contact for a given letter.
func (wheel *Wheel) GetEntryContact(letter Rune) (result Rune) {
	letterNum := int(letter)
	groundNum := int(wheel.GroundSetting)
	ringNum := int(wheel.RingSetting)
	offset := (ringNum - 65) + (65 - groundNum)
	resultNum := letterNum - offset

	if resultNum > 90 {
		resultNum -= 26
	} else if resultNum < 65 {
		resultNum += 26
	}

	result = Rune(resultNum)

	return
}

// GetExitContact gets the exit contact for a given letter.
func (wheel *Wheel) GetExitContact(letter Rune) (result Rune) {
	letterNum := int(letter)
	groundNum := int(wheel.GroundSetting)
	ringNum := int(wheel.RingSetting)
	offset := (ringNum - 65) + (65 - groundNum)
	resultNum := letterNum + offset

	if resultNum > 90 {
		resultNum -= 26
	} else if resultNum < 65 {
		resultNum += 26
	}

	result = Rune(resultNum)

	return
}

// Step the rotor, returning boolean value representing if turnover is required.
func (wheel *Wheel) Step() (turnover bool) {
	groundNum := int(wheel.GroundSetting) + 1

	if groundNum > 90 {
		groundNum -= 26
	}

	_, turnover = WheelTurnoverPoints[wheel.Number][wheel.GroundSetting]
	wheel.GroundSetting = Rune(groundNum)

	return 
}

// New creates a new instance of an Enigma from a settings JSON file.
func New(settingsPath string) (enigma *Enigma, err error) {
	enigma = new(Enigma)
	settings, err := ioutil.ReadFile(settingsPath)

	if err != nil {
		return
	}

	err = json.Unmarshal(settings, &enigma)

	return
}

// Step each wheel (if required), accounting also for doublestep occurances.
func (enigma *Enigma) Step() {
	for wheelIndex := len(enigma.Wheels) - 1; wheelIndex >= 0; wheelIndex-- {
		// We don't want to rotate these wheels.
		if enigma.Wheels[wheelIndex].Number == "Beta" {
			break
		} else if enigma.Wheels[wheelIndex].Number == "Gamma" {
			break
		}

		if !enigma.Wheels[wheelIndex].Step() {
			// Doubestepping.
			if wheelIndex > 0 {
				doublestepWheel := enigma.Wheels[wheelIndex - 1]
				_, doublestep := WheelTurnoverPoints[doublestepWheel.Number][doublestepWheel.GroundSetting]

				if doublestep {
					continue
				}
			}

			break
		}
	}
}

// MapToPlugboard maps letters through the Enigma plugboard.
func (enigma *Enigma) MapToPlugboard(letter Rune) (result Rune) {
	result = letter

	for _, plugPair := range enigma.Plugboard {
		if plugPair[0] == result {
			result = plugPair[1]

			break
		} else if plugPair[1] == result {
			result = plugPair[0]

			break
		}
	}

	return
}

// Key inputs a letter into the Enigma machine, returning the ciphertext.
func (enigma *Enigma) Key(letter Rune) (result Rune, err error) {
	enigma.Step()

	result = enigma.MapToPlugboard(letter)

	for wheelIndex := len(enigma.Wheels) - 1; wheelIndex >= 0; wheelIndex-- {
		result = enigma.Wheels[wheelIndex].GetEntryContact(result)
		result = WheelWiring[enigma.Wheels[wheelIndex].Number][result]
		result = enigma.Wheels[wheelIndex].GetExitContact(result)
	}

	if _, ok := ReflectorWiring[enigma.Reflector]; !ok {
		err = errors.New("Invalid reflector.")

		return
	}

	result = ReflectorWiring[enigma.Reflector][result]

	for wheelIndex := 0; wheelIndex < len(enigma.Wheels); wheelIndex++ {
		result = enigma.Wheels[wheelIndex].GetEntryContact(result)

		for key, value := range WheelWiring[enigma.Wheels[wheelIndex].Number] {
			if value == result {
				result = key

				break
			}
		}

		result = enigma.Wheels[wheelIndex].GetExitContact(result)
	}

	result = enigma.MapToPlugboard(result)

	return
}

// Encrypt a string, returning the ciphertext.
func (enigma *Enigma) Encrypt(plainText string) (cipherText string, err error) {
	var result Rune

	for _, ch := range plainText {
		upperCh := unicode.ToUpper(ch)

		if upperCh < 65 || upperCh > 90 {
			cipherText += string(ch)
		} else {
			result, err = enigma.Key(Rune(upperCh))

			if err != nil {
				return
			}

			if unicode.IsLower(ch) {
				result = Rune(unicode.ToLower(rune(result)))
			}

			cipherText += string(result)
		}
	}

	return
}

// Takes letters from a byte slice and writes encrypted ciphertext result to STDOUT.
func (enigma *Enigma) Write(data []byte) (numBytes int, err error) {
	result, err := enigma.Encrypt(string(data))

	if err != nil {
		return
	}

	numBytes, err = os.Stdout.Write([]byte(result))

	return
}
