package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type Key struct {
	rune
}

var WheelWiring = map[string]map[string]map[Key]Key{
	"M3": {
		"I": {
			Key{'A'}: Key{'E'}, Key{'B'}: Key{'K'}, Key{'C'}: Key{'M'}, Key{'D'}: Key{'F'}, Key{'E'}: Key{'L'}, Key{'F'}: Key{'G'},
			Key{'G'}: Key{'D'}, Key{'H'}: Key{'Q'}, Key{'I'}: Key{'V'}, Key{'J'}: Key{'Z'}, Key{'K'}: Key{'N'}, Key{'L'}: Key{'T'},
			Key{'M'}: Key{'O'}, Key{'N'}: Key{'W'}, Key{'O'}: Key{'Y'}, Key{'P'}: Key{'H'}, Key{'Q'}: Key{'X'}, Key{'R'}: Key{'U'},
			Key{'S'}: Key{'S'}, Key{'T'}: Key{'P'}, Key{'U'}: Key{'A'}, Key{'V'}: Key{'I'}, Key{'W'}: Key{'B'}, Key{'X'}: Key{'R'},
			Key{'Y'}: Key{'C'}, Key{'Z'}: Key{'J'},
		},
		"II": {
			Key{'A'}: Key{'A'}, Key{'B'}: Key{'J'}, Key{'C'}: Key{'D'}, Key{'D'}: Key{'K'}, Key{'E'}: Key{'S'}, Key{'F'}: Key{'I'},
			Key{'G'}: Key{'R'}, Key{'H'}: Key{'U'}, Key{'I'}: Key{'X'}, Key{'J'}: Key{'B'}, Key{'K'}: Key{'L'}, Key{'L'}: Key{'H'},
			Key{'M'}: Key{'W'}, Key{'N'}: Key{'T'}, Key{'O'}: Key{'M'}, Key{'P'}: Key{'C'}, Key{'Q'}: Key{'Q'}, Key{'R'}: Key{'G'},
			Key{'S'}: Key{'Z'}, Key{'T'}: Key{'N'}, Key{'U'}: Key{'P'}, Key{'V'}: Key{'Y'}, Key{'W'}: Key{'F'}, Key{'X'}: Key{'V'},
			Key{'Y'}: Key{'O'}, Key{'Z'}: Key{'E'},
		},
		"III": {
			Key{'A'}: Key{'B'}, Key{'B'}: Key{'D'}, Key{'C'}: Key{'F'}, Key{'D'}: Key{'H'}, Key{'E'}: Key{'J'}, Key{'F'}: Key{'L'},
			Key{'G'}: Key{'C'}, Key{'H'}: Key{'P'}, Key{'I'}: Key{'R'}, Key{'J'}: Key{'T'}, Key{'K'}: Key{'X'}, Key{'L'}: Key{'V'},
			Key{'M'}: Key{'Z'}, Key{'N'}: Key{'N'}, Key{'O'}: Key{'Y'}, Key{'P'}: Key{'E'}, Key{'Q'}: Key{'I'}, Key{'R'}: Key{'W'},
			Key{'S'}: Key{'G'}, Key{'T'}: Key{'A'}, Key{'U'}: Key{'K'}, Key{'V'}: Key{'M'}, Key{'W'}: Key{'U'}, Key{'X'}: Key{'S'},
			Key{'Y'}: Key{'Q'}, Key{'Z'}: Key{'O'},
		},
		"IV": {
			Key{'A'}: Key{'E'}, Key{'B'}: Key{'S'}, Key{'C'}: Key{'O'}, Key{'D'}: Key{'V'}, Key{'E'}: Key{'P'}, Key{'F'}: Key{'Z'},
			Key{'G'}: Key{'J'}, Key{'H'}: Key{'A'}, Key{'I'}: Key{'Y'}, Key{'J'}: Key{'Q'}, Key{'K'}: Key{'U'}, Key{'L'}: Key{'I'},
			Key{'M'}: Key{'R'}, Key{'N'}: Key{'H'}, Key{'O'}: Key{'X'}, Key{'P'}: Key{'L'}, Key{'Q'}: Key{'N'}, Key{'R'}: Key{'F'},
			Key{'S'}: Key{'T'}, Key{'T'}: Key{'G'}, Key{'U'}: Key{'K'}, Key{'V'}: Key{'D'}, Key{'W'}: Key{'C'}, Key{'X'}: Key{'M'},
			Key{'Y'}: Key{'W'}, Key{'Z'}: Key{'B'},
		},
		"V": {
			Key{'A'}: Key{'V'}, Key{'B'}: Key{'Z'}, Key{'C'}: Key{'B'}, Key{'D'}: Key{'R'}, Key{'E'}: Key{'G'}, Key{'F'}: Key{'I'},
			Key{'G'}: Key{'T'}, Key{'H'}: Key{'Y'}, Key{'I'}: Key{'U'}, Key{'J'}: Key{'P'}, Key{'K'}: Key{'S'}, Key{'L'}: Key{'D'},
			Key{'M'}: Key{'N'}, Key{'N'}: Key{'H'}, Key{'O'}: Key{'L'}, Key{'P'}: Key{'X'}, Key{'Q'}: Key{'A'}, Key{'R'}: Key{'W'},
			Key{'S'}: Key{'M'}, Key{'T'}: Key{'J'}, Key{'U'}: Key{'Q'}, Key{'V'}: Key{'O'}, Key{'W'}: Key{'F'}, Key{'X'}: Key{'E'},
			Key{'Y'}: Key{'C'}, Key{'Z'}: Key{'K'},
		},
		"VI": {
			Key{'A'}: Key{'J'}, Key{'B'}: Key{'P'}, Key{'C'}: Key{'G'}, Key{'D'}: Key{'V'}, Key{'E'}: Key{'O'}, Key{'F'}: Key{'U'},
			Key{'G'}: Key{'M'}, Key{'H'}: Key{'F'}, Key{'I'}: Key{'Y'}, Key{'J'}: Key{'Q'}, Key{'K'}: Key{'B'}, Key{'L'}: Key{'E'},
			Key{'M'}: Key{'N'}, Key{'N'}: Key{'H'}, Key{'O'}: Key{'Z'}, Key{'P'}: Key{'R'}, Key{'Q'}: Key{'D'}, Key{'R'}: Key{'K'},
			Key{'S'}: Key{'A'}, Key{'T'}: Key{'S'}, Key{'U'}: Key{'X'}, Key{'V'}: Key{'L'}, Key{'W'}: Key{'I'}, Key{'X'}: Key{'C'},
			Key{'Y'}: Key{'T'}, Key{'Z'}: Key{'W'},
		},
		"VII": {
			Key{'A'}: Key{'N'}, Key{'B'}: Key{'Z'}, Key{'C'}: Key{'J'}, Key{'D'}: Key{'H'}, Key{'E'}: Key{'G'}, Key{'F'}: Key{'R'},
			Key{'G'}: Key{'C'}, Key{'H'}: Key{'X'}, Key{'I'}: Key{'M'}, Key{'J'}: Key{'Y'}, Key{'K'}: Key{'S'}, Key{'L'}: Key{'W'},
			Key{'M'}: Key{'B'}, Key{'N'}: Key{'O'}, Key{'O'}: Key{'U'}, Key{'P'}: Key{'F'}, Key{'Q'}: Key{'A'}, Key{'R'}: Key{'I'},
			Key{'S'}: Key{'V'}, Key{'T'}: Key{'L'}, Key{'U'}: Key{'P'}, Key{'V'}: Key{'E'}, Key{'W'}: Key{'K'}, Key{'X'}: Key{'Q'},
			Key{'Y'}: Key{'D'}, Key{'Z'}: Key{'T'},
		},
		"VIII": {
			Key{'A'}: Key{'F'}, Key{'B'}: Key{'K'}, Key{'C'}: Key{'Q'}, Key{'D'}: Key{'H'}, Key{'E'}: Key{'T'}, Key{'F'}: Key{'L'},
			Key{'G'}: Key{'X'}, Key{'H'}: Key{'O'}, Key{'I'}: Key{'C'}, Key{'J'}: Key{'B'}, Key{'K'}: Key{'J'}, Key{'L'}: Key{'S'},
			Key{'M'}: Key{'P'}, Key{'N'}: Key{'D'}, Key{'O'}: Key{'Z'}, Key{'P'}: Key{'R'}, Key{'Q'}: Key{'A'}, Key{'R'}: Key{'M'},
			Key{'S'}: Key{'E'}, Key{'T'}: Key{'W'}, Key{'U'}: Key{'N'}, Key{'V'}: Key{'I'}, Key{'W'}: Key{'U'}, Key{'X'}: Key{'Y'},
			Key{'Y'}: Key{'G'}, Key{'Z'}: Key{'V'},
		},
	},
	"M4": {
		"Beta": {
			Key{'A'}: Key{'L'}, Key{'B'}: Key{'E'}, Key{'C'}: Key{'Y'}, Key{'D'}: Key{'J'}, Key{'E'}: Key{'V'}, Key{'F'}: Key{'C'},
			Key{'G'}: Key{'N'}, Key{'H'}: Key{'I'}, Key{'I'}: Key{'X'}, Key{'J'}: Key{'W'}, Key{'K'}: Key{'P'}, Key{'L'}: Key{'B'},
			Key{'M'}: Key{'Q'}, Key{'N'}: Key{'M'}, Key{'O'}: Key{'D'}, Key{'P'}: Key{'R'}, Key{'Q'}: Key{'T'}, Key{'R'}: Key{'A'},
			Key{'S'}: Key{'K'}, Key{'T'}: Key{'Z'}, Key{'U'}: Key{'G'}, Key{'V'}: Key{'F'}, Key{'W'}: Key{'U'}, Key{'X'}: Key{'H'},
			Key{'Y'}: Key{'O'}, Key{'Z'}: Key{'S'},
		},
		"Gamma": {
			Key{'A'}: Key{'F'}, Key{'B'}: Key{'S'}, Key{'C'}: Key{'O'}, Key{'D'}: Key{'K'}, Key{'E'}: Key{'A'}, Key{'F'}: Key{'N'},
			Key{'G'}: Key{'U'}, Key{'H'}: Key{'E'}, Key{'I'}: Key{'R'}, Key{'J'}: Key{'H'}, Key{'K'}: Key{'M'}, Key{'L'}: Key{'B'},
			Key{'M'}: Key{'T'}, Key{'N'}: Key{'I'}, Key{'O'}: Key{'Y'}, Key{'P'}: Key{'C'}, Key{'Q'}: Key{'W'}, Key{'R'}: Key{'L'},
			Key{'S'}: Key{'Q'}, Key{'T'}: Key{'P'}, Key{'U'}: Key{'Z'}, Key{'V'}: Key{'X'}, Key{'W'}: Key{'V'}, Key{'X'}: Key{'G'},
			Key{'Y'}: Key{'J'}, Key{'Z'}: Key{'D'},
		},
	},
}

func (key *Key) UnmarshalJSON(bytes []byte) (err error) {
	key.rune = rune(bytes[1])

	fmt.Println(key.rune)

	return
}

type Wheel struct {
	Number string `json:"number"`
	RingSetting Key `json:"ring setting"`
	GroundSetting Key `json:"ground setting"`
}

type PlugboardMap map[Key]Key

/*func (plugboardMap *PlugboardMap) UnmarshalJSON(bytes []byte) (err error) {
	

	return
}*/

type Enigma struct {
	Model string `json:"modelx"`
	Reflector string `json:"reflectorx"`
	Wheels []Wheel `json:"wheelsx"`
	Plugboard PlugboardMap `json:"plugboard"`
}

func (enigma *Enigma) Key(inKey rune) (outKey rune) {
	// plugboard
	if plugboardKey, ok := enigma.Plugboard[Key{inKey}]; ok {
		outKey = plugboardKey.rune
	}

	return
}

func (enigma *Enigma) Encrypt(plainText string, wheelSettings string) (cipherText string, err error) {
	cipherText = wheelSettings

	encryptedWheelSettings := strings.Map(enigma.Key, wheelSettings)
	//for _, wheelSetting := range wheelSettings {
	//	encryptedWheelSettings += enigma.Key(wheelSetting)
	//}

	for i, encryptedWheelSetting := range encryptedWheelSettings {
		enigma.Wheels[i].GroundSetting = Key{encryptedWheelSetting}
	}

	cipherText = strings.Map(enigma.Key, plainText)

	//for _, ch := range plainText {
	//	cipherText += enigma.Key(ch)
	//}

	return
}

func (enigma *Enigma) Decrypt(cipherText string) (plainText string, err error) {
	wheelSettings := cipherText[: 3]
	cipherText = cipherText[3 :]
	decryptedWheelSettings := ""
	decryptedWheelSettings = strings.Map(enigma.Key, wheelSettings)

	//for _, wheelSetting := range wheelSettings {
	//	decryptedWheelSettings += enigma.Key(wheelSetting)
	//}

	for i, decryptedWheelSetting := range decryptedWheelSettings {
		enigma.Wheels[i].GroundSetting = Key{decryptedWheelSetting}
	}

	plainText = strings.Map(enigma.Key, cipherText)

	//for _, ch := range cipherText {
	//	plainText += enigma.Key(ch)
	//}

	return
}

func main() {
	enigma := Enigma{}
	settings, err := ioutil.ReadFile("settings.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(settings, &enigma)

	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Fatal("")

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
