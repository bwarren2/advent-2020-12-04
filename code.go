package advent

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// Record stores data about passports
type Record struct {
	Byr, Iyr, Eyr, Hgt, Hcl, Ecl, Pid, Cid string
}

var hclValid = regexp.MustCompile(`^#[a-f0-9]{6}$`)
var pidValid = regexp.MustCompile(`^[0-9]{9}$`)
var pidValidLoose = regexp.MustCompile(`[0-9]{9}`)

// IsValidv1 decides whether a record is valid
func (r *Record) IsValidv1() bool {
	return r.Byr != "" &&
		r.Iyr != "" &&
		r.Eyr != "" &&
		r.Hgt != "" &&
		r.Hcl != "" &&
		r.Ecl != "" &&
		r.Pid != ""
}

// IsValidv2 decides whether a record is valid for part 2
func (r *Record) IsValidv2() bool {
	return r.ByrIsValid() &&
		r.IyrIsValid() &&
		r.EyrIsValid() &&
		r.HgtIsValid() &&
		r.HclIsValid() &&
		r.EclIsValid() &&
		r.PidIsValid()
}

// ByrIsValid decides whether a birth year is valid
func (r *Record) ByrIsValid() bool {
	parsedYr, err := strconv.ParseInt(r.Byr, 10, 64)
	if err != nil {
		return false
	}
	return 1920 <= parsedYr && parsedYr <= 2002
}

// IyrIsValid decides whether an issue year is valid
func (r *Record) IyrIsValid() bool {
	parsedYr, err := strconv.ParseInt(r.Iyr, 10, 64)
	if err != nil {
		return false
	}
	return 2010 <= parsedYr && parsedYr <= 2020
}

// EyrIsValid decides whether an expiration year is valid
func (r *Record) EyrIsValid() bool {
	parsedYr, err := strconv.ParseInt(r.Eyr, 10, 64)
	if err != nil {
		return false
	}
	return parsedYr >= 2020 && parsedYr <= 2030
}

// HgtIsValid decides whether an expiration year is valid
func (r *Record) HgtIsValid() (result bool) {
	if len(r.Hgt) <= 2 {
		return false
	}
	parsedHeight, err := strconv.ParseInt(r.Hgt[:len(r.Hgt)-2], 10, 64)
	if err != nil {
		fmt.Println(r.Hgt)
		return false
	}
	units := r.Hgt[len(r.Hgt)-2:]
	if units == "in" {
		result = 59 <= parsedHeight && parsedHeight <= 76
	} else if units == "cm" {
		result = 150 <= parsedHeight && parsedHeight <= 193
	} else {
		fmt.Println(r.Hgt)
		result = false
	}

	return
}

// HclIsValid decides whether an expiration year is valid
func (r *Record) HclIsValid() bool {
	return hclValid.MatchString(r.Hcl)
}

var eclOptions = []string{
	"amb",
	"blu",
	"brn",
	"gry",
	"grn",
	"hzl",
	"oth",
}

// EclIsValid decides whether an expiration year is valid
func (r *Record) EclIsValid() bool {

	for _, choice := range eclOptions {
		if choice == r.Ecl {
			return true
		}
	}
	return false
}

// PidIsValid decides whether an expiration year is valid
func (r *Record) PidIsValid() bool {
	if !pidValid.MatchString(r.Pid) && pidValidLoose.MatchString(r.Pid) {
		fmt.Println(r.Pid, len(r.Pid))
	}
	return pidValid.MatchString(r.Pid)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// LinesFromFile gets all the lines from a file
func LinesFromFile(filename string) (result []string) {
	file, err := os.Open(filename)
	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return
}

// RecordFromString turns space-separated colon-values into a Record
func RecordFromString(input string) *Record {
	output := new(Record)
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		token := scanner.Text()
		if strings.Contains(token, ":") {
			pieces := strings.Split(token, ":")
			field := strings.Title(pieces[0])
			value := pieces[1]
			reflect.ValueOf(output).Elem().FieldByName(field).SetString(value)
		}
	}
	return output
}

// RecordIterator returns records
func RecordIterator(filename string) <-chan *Record {
	iterator := make(chan *Record)
	lines := LinesFromFile(filename)

	buildString := new(strings.Builder)
	go func(lines []string) {
		for idx, line := range lines {
			buildString.WriteString(" ")
			buildString.WriteString(line)
			if line == "" || idx == len(lines)-1 {
				iterator <- RecordFromString(buildString.String())
				buildString = new(strings.Builder)
			}
		}
		close(iterator)
	}(lines)
	return iterator
}

// Part1 answers part 1
func Part1(filename string) (count int) {
	c1 := RecordIterator(filename)
	for record := range c1 {
		if record.IsValidv1() {
			count++
		}
	}
	return
}

// Part2 answers part 1
func Part2(filename string) (count int) {
	c1 := RecordIterator(filename)
	for record := range c1 {
		// fmt.Println(record.IsValidv2(), record.HclIsValid())
		if record.IsValidv2() {
			count++
		}
	}
	return
}
