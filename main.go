package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type User struct {
	Session string `yaml:"session"`
	XSRF    string `yaml:"xsrf"`
}

type Response struct {
	Username   string `json:"username"`
	Available  bool   `json:"available"`
	Message    string `json:"message"`
	Cost       int    `json:"cost"`
	CostString string `json:"costString"`
}

func (c *User) loadConf() *User {
	exist := doesExist("config.yaml")
	if exist {
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Printf("yamlFile.Get err   #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, c)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		return c
	} else {
		genConfig("config.yaml", "session:\nxsrf:")
	}
	return nil
}

func (c *User) checkUser(user string) {

	name := fmt.Sprintf("username=%v", user)

	payload := strings.NewReader(name)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://osu.ppy.sh/users/check-username-availability", payload)
	if err != nil {
		log.Fatal(err)
	}

	Cookie := fmt.Sprintf("XSRF-TOKEN=%v; osu_session=%v", c.XSRF, c.Session)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:95.0) Gecko/20100101 Firefox/95.0")
	req.Header.Set("Cookie", Cookie)
	req.Header.Set("X-CSRF-Token", c.XSRF)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored when sending request to the server")
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data Response
	erro := json.Unmarshal(responseBody, &data)
	if erro != nil {
		log.Fatal(erro, "", data)
	}

	if data.Available {
		fmt.Printf("%v is available!\n", data.Username)
	}

	switch resp.StatusCode {
	case 200:
		return
	default:
		fmt.Printf("Different status code wtf | %v", resp.StatusCode)
		os.Exit(0)
	}
}

func main() {
	var c User
	c.loadConf()

	exist := doesExist("list.txt")

	if exist {
		f, _ := os.Open("list.txt")

		scanner := bufio.NewScanner(f)

		items := 0
		for scanner.Scan() {
			line := scanner.Text()
			c.checkUser(line)
			items += 1
			if items >= 60 {
				fmt.Println("Sleeping for 60 seconds to prevent ratelimiting...")
				time.Sleep(60 * time.Second)
				items = 0
			}
		}
	} else {
		_, err := os.Create("list.txt")
		if err != nil {
			panic(err)
		}
		main()
	}

}

func doesExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func genConfig(path string, data string) {
	f, e := os.Create(path)
	if e != nil {
		panic(e)
	}
	fmt.Fprint(f, data)
	main()
}
