package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

func main123() {
	fmt.Println("hello")
	//s := []string{"000", "1110", "01", "001", "110", "11"}
	//s := []string{"100110", "1001", "1001111"}
	s := []string{"1", "10", "11010"}
	xyz := autocomplete(s)
	fmt.Println(xyz)

}
func main123() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	//resp, err := client.Get("http://rmmitsapi.qaitsupport247.local/rmmitsapi/v1/partners/50000006/sites/50000809/resources/50066388/ValidatePasswordVault")
	resp, _ := client.Get("https://qa.agent.exec.itsupport247.net/agent//v1/endpoint/52635175-7aef-4228-8632-3abd47773bc6/mapping")
	defer resp.Body.Close()

	// decoder := json.NewDecoder(resp.Body)
	// legacyRegIDSt := new(LegacyRegIDSt)
	// err = decoder.Decode(legacyRegIDSt)
	// fmt.Println(err)
	// fmt.Println(legacyRegIDSt)

	body, err := ioutil.ReadAll(resp.Body)
	l := LegacyRegIDSt{}
	json.Unmarshal(body, &l)
	fmt.Println(string(body), err, l)
}

// LegacyRegIDSt legacy reg ID decode
type LegacyRegIDSt struct {
	LegacyRegID string `json:"legacy_reg_id,omitempty"`
}

func main12() {

	s := []int32{10, 20, 7}
	xyz := minSum(s, 4)
	fmt.Println(xyz)

}

func minSum(num []int32, k int32) int32 {
	// Write your code here

	var minSum int32
	var i int32

	sort.Slice(num, func(i, j int) bool { return num[i] < num[j] })
	fmt.Println(num)
	for i < k {
		index := findMax(num)
		num[index] = performOperation(num[index])
		i++
	}

	for _, v := range num {
		minSum = minSum + v
	}
	return minSum

}

func performOperation(input int32) int32 {
	if input%2 == 0 {
		input = input / 2
	} else {
		input = input/2 + 1
	}

	return input
}

func findMax123(num []int32) int {
	var max int32
	max = math.MinInt32
	index := 0

	for i, v := range num {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func findMax(num []int32) int {
	var max int32
	max = num[len(num)-1]
	index := 0

	for i := len(num) - 2; i >= 0; i-- {
		if max >= num[i] {
			return i + 1
		} else {

		}
	}

	return index
}

func autocomplete(command []string) []int32 {
	prefixes := make(map[string]int32)
	var result []int32
	result = append(result, 0)

	for i, v := range command {
		k := 1
		for k <= len(v) {
			prefixes[v[0:k]+"-"+strconv.Itoa(i)] = int32(i)
			k++
		}
	}

	for i, v := range command {
		if i == 0 {
			continue
		}
		k := 0
		var res int32
		for k < len(v) {
			j := 0
			for j < i {
				key := v[0:k] + "-" + strconv.Itoa(j)
				if value, ok := prefixes[key]; ok {
					res = value
				}
				j++
			}
			k++
		}
		res = res + 1
		result = append(result, res)
	}
	return result
}

//insert
