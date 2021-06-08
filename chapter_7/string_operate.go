package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Println(strings.Index("chiken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println("ba" + strings.Repeat("na", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))

	fmt.Printf("[%q]", strings.Trim(" !!! Achtung !!! ", "! "))

	fmt.Printf("Fields are: %q", strings.Fields("  foo bar baz  "))

	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	f, err := strconv.ParseBool("false")
	checkError(err)
	g, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	h, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	i, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	j, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(f, g, h, i, j)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
