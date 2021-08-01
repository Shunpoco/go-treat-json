package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func encoder(v interface{}) error {
	err := json.NewEncoder(os.Stdout).Encode(v)
	if err != nil {
		return err
	}

	return nil
}

func decoder() (interface{}, error) {
	var v map[string]string
	err := json.NewDecoder(os.Stdin).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

func main() {
	m := map[string]string{"hoge": "fuga"}
	b := new(bytes.Buffer)
	err := json.NewEncoder(b).Encode(m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.String())

	encoder(m)

	const jsonStream = `
		{"name": "ed", "text": "knock knock."}
		{"name": "sam", "text": "who's there?"}
	`

	type Message struct {
		Name string `json:"name"`
		Text string `json:"text"`
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}

	m2, _ := decoder()

	fmt.Println(m2)

	m3 := Message{Name: "Smith", Text: "aaa"}

	b2, _ := json.Marshal(m3)

	var m4 Message
	json.Unmarshal(b2, &m4)

	fmt.Println(m4)
}
