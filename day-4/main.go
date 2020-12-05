package main

import (
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const inputFilePath = "input.txt"

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	log.Println("Loading:", inputFilePath)
	byteArray, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		panic(err)
	}
	data := string(byteArray)

	passportSamples := strings.Split(data, "\n\n")
	log.Println("Found", len(passportSamples), "passport samples")

	passports := parsePassports(passportSamples)
	log.Println(len(passports), "passports parsed")

	validPassports := 0
	for _, p := range passports {
		if p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" &&
			p.pid != "" {
			validPassports++
		}
	}
	log.Println("Valid passports:", validPassports)

	passports2 := parsePassports2(passportSamples)
	validPassports2 := 0
	for _, p := range passports2 {
		if p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" &&
			p.pid != "" {
			validPassports2++
		} else {
		}
	}
	log.Println("Valid passports2:", validPassports2)
}

func parsePassports(passportSample []string) []passport {

	var passports []passport
	for _, sample := range passportSample {
		passport := passport{}
		byr := regexp.MustCompile(`byr:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(byr) > 1 {
			passport.byr = byr[1]
		}
		iyr := regexp.MustCompile(`iyr:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(iyr) > 1 {
			passport.iyr = iyr[1]
		}
		eyr := regexp.MustCompile(`eyr:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(eyr) > 1 {
			passport.eyr = eyr[1]
		}
		hgt := regexp.MustCompile(`hgt:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(hgt) > 1 {
			passport.hgt = hgt[1]
		}
		hcl := regexp.MustCompile(`hcl:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(hcl) > 1 {
			passport.hcl = hcl[1]
		}
		ecl := regexp.MustCompile(`ecl:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(ecl) > 1 {
			passport.ecl = ecl[1]
		}
		pid := regexp.MustCompile(`pid:([\#a-z0-9]+)`).FindStringSubmatch(sample)
		if len(pid) > 1 {
			passport.pid = pid[1]
		}

		passports = append(passports, passport)
	}
	return passports
}

func parsePassports2(passportSample []string) []passport {

	var passports []passport
	for _, sample := range passportSample {
		passport := passport{}

		byr := regexp.MustCompile(`byr:(\d{4})`).FindStringSubmatch(sample)
		if len(byr) == 2 {
			value, _ := strconv.Atoi(byr[1])
			if value >= 1920 && value <= 2002 {
				passport.byr = byr[1]
			}
		}
		iyr := regexp.MustCompile(`iyr:(\d{4})`).FindStringSubmatch(sample)
		if len(iyr) == 2 {
			value, _ := strconv.Atoi(iyr[1])
			if value >= 2010 && value <= 2020 {
				passport.iyr = iyr[1]
			}
		}
		eyr := regexp.MustCompile(`eyr:(\d{4})`).FindStringSubmatch(sample)
		if len(eyr) == 2 {
			value, _ := strconv.Atoi(eyr[1])
			if value >= 2020 && value <= 2030 {
				passport.eyr = eyr[1]
			}
		}
		hgt := regexp.MustCompile(`hgt:(\d+)(cm|in){1}`).FindStringSubmatch(sample)
		if len(hgt) == 3 {
			value, _ := strconv.Atoi(hgt[1])
			if hgt[2] == "cm" && (value >= 150 && value <= 193) {
				passport.hgt = hgt[1]
			}

			if hgt[2] == "in" && (value >= 59 && value <= 76) {
				passport.hgt = hgt[1]
			}
		}
		hcl := regexp.MustCompile(`hcl:\#([a-f0-9]{6})`).FindStringSubmatch(sample)
		if len(hcl) == 2 {
			passport.hcl = hcl[1]
		}
		ecl := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`).FindStringSubmatch(sample)
		if len(ecl) == 2 {
			passport.ecl = ecl[1]
		}
		pid := regexp.MustCompile(`pid:(\d+)`).FindStringSubmatch(sample)
		if len(pid) == 2 && len(pid[1]) == 9 {
			passport.pid = pid[1]
		}

		passports = append(passports, passport)
	}
	return passports
}
