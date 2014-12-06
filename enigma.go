package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Rune rune

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

var WheelWiring = map[string]map[string]map[Rune]Rune{
	"M3": {
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
	},
	"M4": {
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
	},
}

type Wheel struct {
	Number string `json:"number"`
	RingSetting Rune `json:"ring setting"`
	GroundSetting Rune `json:"ground setting"`
}

type Enigma struct {
	Model string `json:"model"`
	Reflector string `json:"reflector"`
	Wheels []Wheel `json:"wheels"`
	Plugboard [][]Rune `json:"plugboard"`
}

func (c *Rune) UnmarshalJSON(data []byte) (err error) {
	*c = Rune(data[1])

	return
}

func (wheel *Wheel) getEntryContact(letter Rune) (result Rune) {
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

func (wheel *Wheel) getExitContact(letter Rune) (result Rune) {
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

func (wheel *Wheel) step() (turnover bool) {
	groundNum := int(wheel.GroundSetting) + 1

	if groundNum > 90 {
		groundNum -= 26
	}

	_, turnover = WheelTurnoverPoints[wheel.Number][wheel.GroundSetting]
	wheel.GroundSetting = Rune(groundNum)

	return 
}

func New(settingsPath string) (enigma *Enigma, err error) {
	enigma = new(Enigma)
	settings, err := ioutil.ReadFile(settingsPath)

	if err != nil {
		return
	}

	err = json.Unmarshal(settings, &enigma)

	return
}

func (enigma *Enigma) step() {
	for wheelIndex := len(enigma.Wheels) - 1; wheelIndex >= 0; wheelIndex-- {
		if !enigma.Wheels[wheelIndex].step() {
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

func (enigma *Enigma) Key(letter Rune) (result Rune) {
	enigma.step()

	result = enigma.MapToPlugboard(letter)

	for wheelIndex := len(enigma.Wheels) - 1; wheelIndex >= 0; wheelIndex-- {
		result = enigma.Wheels[wheelIndex].getEntryContact(result)
		result = WheelWiring[enigma.Model][enigma.Wheels[wheelIndex].Number][result]
		result = enigma.Wheels[wheelIndex].getExitContact(result)
	}

	result = ReflectorWiring[enigma.Reflector][result]

	for wheelIndex := 0; wheelIndex < len(enigma.Wheels); wheelIndex++ {
		result = enigma.Wheels[wheelIndex].getEntryContact(result)

		for key, value := range WheelWiring[enigma.Model][enigma.Wheels[wheelIndex].Number] {
			if value == result {
				result = key

				break
			}
		}

		result = enigma.Wheels[wheelIndex].getExitContact(result)
	}

	result = enigma.MapToPlugboard(result)

	return
}

func (enigma *Enigma) Encrypt(plainText string) (cipherText string, err error) {
	for _, ch := range plainText {
		cipherText += string(enigma.Key(Rune(ch)))
	}

	return
}

func (enigma *Enigma) Decrypt(cipherText string) (plainText string, err error) {
	plainText, err = enigma.Encrypt(cipherText)

	return
}

func main() {
	enigma, err := New("settings.json")

	if err != nil {
		log.Fatal(err)
	}

	plainText := "KKK"

	cipherText, err := enigma.Encrypt(plainText)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cipherText)

	enigma, err = New("settings.json")

	if err != nil {
		log.Fatal(err)
	}

	plainText, err = enigma.Decrypt(cipherText)

	fmt.Println(plainText)
}
