package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

var WheelWiring = map[string]map[string]map[string]string{
	"M3": {
		"I": {
			"A": "E", "B": "K", "C": "M", "D": "F", "E": "L", "F": "G",
			"G": "D", "H": "Q", "I": "V", "J": "Z", "K": "N", "L": "T",
			"M": "O", "N": "W", "O": "Y", "P": "H", "Q": "X", "R": "U",
			"S": "S", "T": "P", "U": "A", "V": "I", "W": "B", "X": "R",
			"Y": "C", "Z": "J",
		},
		"II": {
			"A": "A", "B": "J", "C": "D", "D": "K", "E": "S", "F": "I",
			"G": "R", "H": "U", "I": "X", "J": "B", "K": "L", "L": "H",
			"M": "W", "N": "T", "O": "M", "P": "C", "Q": "Q", "R": "G",
			"S": "Z", "T": "N", "U": "P", "V": "Y", "W": "F", "X": "V",
			"Y": "O", "Z": "E",
		},
		"III": {
			"A": "B", "B": "D", "C": "F", "D": "H", "E": "J", "F": "L",
			"G": "C", "H": "P", "I": "R", "J": "T", "K": "X", "L": "V",
			"M": "Z", "N": "N", "O": "Y", "P": "E", "Q": "I", "R": "W",
			"S": "G", "T": "A", "U": "K", "V": "M", "W": "U", "X": "S",
			"Y": "Q", "Z": "O",
		},
		"IV": {
			"A": "E", "B": "S", "C": "O", "D": "V", "E": "P", "F": "Z",
			"G": "J", "H": "A", "I": "Y", "J": "Q", "K": "U", "L": "I",
			"M": "R", "N": "H", "O": "X", "P": "L", "Q": "N", "R": "F",
			"S": "T", "T": "G", "U": "K", "V": "D", "W": "C", "X": "M",
			"Y": "W", "Z": "B",
		},
		"V": {
			"A": "V", "B": "Z", "C": "B", "D": "R", "E": "G", "F": "I",
			"G": "T", "H": "Y", "I": "U", "J": "P", "K": "S", "L": "D",
			"M": "N", "N": "H", "O": "L", "P": "X", "Q": "A", "R": "W",
			"S": "M", "T": "J", "U": "Q", "V": "O", "W": "F", "X": "E",
			"Y": "C", "Z": "K",
		},
		"VI": {
			"A": "J", "B": "P", "C": "G", "D": "V", "E": "O", "F": "U",
			"G": "M", "H": "F", "I": "Y", "J": "Q", "K": "B", "L": "E",
			"M": "N", "N": "H", "O": "Z", "P": "R", "Q": "D", "R": "K",
			"S": "A", "T": "S", "U": "X", "V": "L", "W": "I", "X": "C",
			"Y": "T", "Z": "W",
		},
		"VII": {
			"A": "N", "B": "Z", "C": "J", "D": "H", "E": "G", "F": "R",
			"G": "C", "H": "X", "I": "M", "J": "Y", "K": "S", "L": "W",
			"M": "B", "N": "O", "O": "U", "P": "F", "Q": "A", "R": "I",
			"S": "V", "T": "L", "U": "P", "V": "E", "W": "K", "X": "Q",
			"Y": "D", "Z": "T",
		},
		"VIII": {
			"A": "F", "B": "K", "C": "Q", "D": "H", "E": "T", "F": "L",
			"G": "X", "H": "O", "I": "C", "J": "B", "K": "J", "L": "S",
			"M": "P", "N": "D", "O": "Z", "P": "R", "Q": "A", "R": "M",
			"S": "E", "T": "W", "U": "N", "V": "I", "W": "U", "X": "Y",
			"Y": "G", "Z": "V",
		},
	},
	"M4": {
		"Beta": {
			"A": "L", "B": "E", "C": "Y", "D": "J", "E": "V", "F": "C",
			"G": "N", "H": "I", "I": "X", "J": "W", "K": "P", "L": "B",
			"M": "Q", "N": "M", "O": "D", "P": "R", "Q": "T", "R": "A",
			"S": "K", "T": "Z", "U": "G", "V": "F", "W": "U", "X": "H",
			"Y": "O", "Z": "S",
		},
		"Gamma": {
			"A": "F", "B": "S", "C": "O", "D": "K", "E": "A", "F": "N",
			"G": "U", "H": "E", "I": "R", "J": "H", "K": "M", "L": "B",
			"M": "T", "N": "I", "O": "Y", "P": "C", "Q": "W", "R": "L",
			"S": "Q", "T": "P", "U": "Z", "V": "X", "W": "V", "X": "G",
			"Y": "J", "Z": "D",
		},
	},
}

type Wheel struct {
	Number string `json:"number"`
	RingSetting string `json:"ring setting"`
	GroundSetting string `json:"ground setting"`
}

type PlugboardMap map[string]string

type Enigma struct {
	Model string `json:"model"`
	Reflector string `json:"reflector"`
	Wheels []Wheel `json:"wheels"`
	Plugboard PlugboardMap `json:"plugboard"`
}

func (enigma *Enigma) Encrypt(plainText string, wheelSettings string) (cipherText string, err error) {
	cipherText = wheelSettings
	encryptedWheelSettings := ""

	for _, wheelSetting := range wheelSettings {
		encryptedWheelSettings += enigma.Key(wheelSetting)
	}

	for i, encryptedWheelSetting := range encryptedWheelSettings {
		enigma.Wheels[i].GroundSetting = encryptedWheelSetting
	}

	for _, ch := range plainText {
		cipherText += enigma.Key(ch)
	}

	return
}

func (enigma *Enigma) Decrypt(cipherText string) (plainText string, err error) {
	decryptedWheelSettings := ""

	for _, wheelSetting := range wheelSettings {
		decryptedWheelSettings += enigma.Key(wheelSetting)
	}

	for i, decryptedWheelSetting := range decryptedWheelSettings {
		enigma.Wheels[i].GroundSetting = decryptedWheelSetting
	}

	for _, ch := range cipherText {
		plainText += enigma.Key(ch)
	}

	return
}

func main() {
	enigma := Enigma{}
	settings, err := ioutil.ReadFile("settings.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(settings, &enigma)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(enigma)

	plainText := "Hello Enigma"

	cipherText, err := enigma.Encrypt(plainText, "XKW")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cipherText)

	plainText, err = enigma.Decrypt(cipherText)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(plainText)
}
