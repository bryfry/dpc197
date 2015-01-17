package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/bryfry/amazing"
)

var (
	isbn           string
	checkNeighbors bool
)

func init() {
	flag.StringVar(&isbn, "i", "", "isbn number to validate")
	flag.BoolVar(&checkNeighbors, "n", false, "check neighbor isbns")
	flag.Parse()
}

func main() {
	fmt.Println("Checking for valid ISBN:", isbn)
	isbn := strings.Replace(isbn, "-", "", -1)

	valid, err := ValidISBN10(isbn)
	if err != nil || !valid {
		fmt.Println("Not Valid ISBN-10: ", err)
		isbn, err = CreateValidISBN10(isbn)
		if err != nil {
			fmt.Println("Failed to create a Valid ISBN10:", err)
		}
		fmt.Println("Looking up expected ISBN-10:", isbn)
	}
	url, err := ISBN10URL(isbn)
	if err != nil {
		fmt.Println("Amazon Error:", err)
	}
	fmt.Println("ISBN-10:", isbn, url)

	if checkNeighbors {
		ISBN10Neighbors(isbn)
	}

}

func ISBN10Neighbors(isbn string) {
	fmt.Println("-----   Neighbor ISBNs  -----")
	for i := 0; i < 10; i++ {
		isbn, err := CreateValidISBN10(isbn[:8] + strconv.Itoa(i))
		if err != nil {
			fmt.Println("ISBN-10:", isbn, err)
			continue
		}
		url, err := ISBN10URL(isbn)
		if err != nil {
			fmt.Println(isbn)
			continue
		}
		fmt.Println(isbn, url)
	}

	fmt.Println("----- ----- ----- ----- -----")

}

func ISBN10URL(isbn string) (string, error) {
	client, err := amazing.NewAmazingFromEnv("US")
	if err != nil {
		return "", err
	}

	params := url.Values{
		"IdType":      []string{"ISBN"},
		"ItemId":      []string{isbn},
		"Operation":   []string{"ItemLookup"},
		"SearchIndex": []string{"All"},
	}

	result, err := client.ItemLookup(params)
	if err != nil {
		return "", err
	}

	if len(result.AmazonItems.Request.Errors) > 0 {
		return "", errors.New(result.AmazonItems.Request.Errors[0].Error())
	}

	u, _ := url.Parse(result.AmazonItems.Items[0].DetailPageURL)
	return strings.Split((u.Host + u.Path), "?")[0], nil
}

var ISBN10LengthError = errors.New("Incorrect Length")

func ValidISBN10(isbn string) (bool, error) {
	if len(isbn) != 10 {
		return false, ISBN10LengthError
	}
	expected, err := CheckDigitISBN10(isbn)
	if err != nil {
		return false, err
	}
	recieved := strings.ToUpper(string(isbn[9]))
	if expected != recieved {
		return false, errors.New(fmt.Sprintf("Invalid check digit: expected (%s) received (%s)", expected, recieved))
	}
	return true, nil
}

func CheckDigitISBN10(isbn string) (string, error) {
	if len(isbn) != 9 && len(isbn) != 10 {
		return "", ISBN10LengthError
	} else if len(isbn) == 10 {
		isbn = isbn[:9] // Ignore check digit if provided
	}
	sum := 0
	for i, r := range isbn {
		d, err := isbnAtoi(string(r))
		if err != nil {
			return "", errors.New(fmt.Sprintf("Digit parsing failed (location %d: %s)", i, string(r)))
		}
		sum = sum + (10-i)*d
	}
	return isbnItoa(11 - (sum % 11)), nil
}

func CreateValidISBN10(isbn string) (string, error) {
	if len(isbn) > 10 {
		return "", ISBN10LengthError
	}
	cd, err := CheckDigitISBN10(isbn)
	if err != nil {
		return "", err
	}
	isbn = isbn[:9] + cd
	return isbn, nil
}

func isbnAtoi(digit string) (int, error) {
	if digit == "X" || digit == "x" {
		return 10, nil
	} else {
		return strconv.Atoi(digit)
	}
}

func isbnItoa(digit int) string {
	if digit == 10 {
		return "X"
	}
	if digit == 11 {
		return "0"
	}
	return strconv.Itoa(digit)
}
