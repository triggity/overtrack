package main

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/triggity/overtrack/models"
)

type Vechile interface {
	VE() string
}
type Car struct {
	A string `json:"a"`
	B string `json:"b"`
}

func (c *Car) VE() string {
	return fmt.Sprintf("a: %s, b: %s", c.A, c.B)
}

type Holder struct {
	Name  string  `json:"name"`
	Stats Vechile `json:"stats"`
}

type Character struct {
	Name  string          `json:"name"`
	Stats json.RawMessage `json:"stats"`
}

/*
func (h *Holder) UnmarshalJSON(d []byte) error {
	var c Character
	err := json.Unmarshal(d, &c)
	if err != nil {
		return err
	}
	h.Name = c.Name
	switch c.Name {
	case "car":
		var car Car
		err := json.Unmarshal(c.Stats, &car)
		if err != nil {
			return err
		}
		h.Stats = &car
		return nil
	}
	return errors.New("adfasdfas")

}
*/

/*
func (h *Holder) MarshalJSON() ([]byte, error) {
	var c Character
	err := json.Unmarshal(d, &c)
	if err != nil {
		return err
	}
	h.Name = c.Name
	switch c.Name {
	case "car":
		var car Car
		err := json.Unmarshal(c.Stats, &car)
		if err != nil {
			return err
		}
		h.Stats = &car
		return nil
	}
	return errors.New("adfasdfas")

}
*/

func main() {
	/*
		fmt.Print("time to unmarshal\n")
		d := "{\"name\": \"car\", \"stats\": {\"a\": \"first\", \"b\": \"second\"} }"
		var h Holder
		err := json.Unmarshal([]byte(d), &h)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("holder: %+v\n", h)
		fmt.Printf("please work: %s\n", h.Stats.VE())
		fmt.Printf("car: %+v\n", h.Stats)

		fmt.Print("time to marshal\n")
		b, err := json.Marshal(h)
		if err != nil {
			fmt.Print("err!")
			log.Fatal(err)
		}
		fmt.Printf("marshalled: %s\n", string(b))
	*/

	fmt.Print("time to unmarshal\n")
	d := "{\"name\": \"orisa\", \"stats\": {\"a\": \"first\", \"b\": \"second\", \"damage_blocked\": 45, \"eliminations\": 10} }"
	var h models.CharacterResult
	err := json.Unmarshal([]byte(d), &h)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("holder: %+v\n", h)
	fmt.Printf("iface function: %+v\n", h.Stats)
	fmt.Printf("stats(): %+v\n", h.Stats.CoreStats())

	fmt.Print("time to marshal\n")
	b, err := json.Marshal(h)
	if err != nil {
		fmt.Print("err!")
		log.Fatal(err)
	}
	fmt.Printf("marshalled: %s\n", string(b))

}
