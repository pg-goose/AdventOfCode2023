package main

import (
	"bufio"
	"os"
	"testing"
)

func charIsNum(c byte) bool {
	return c >= '0' && c <= '9'
}

func part1() int {
	file, _ := os.Open("day1.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var i int = 0
		var j int = len(line) - 1
		num := [2]byte{0x00, 0x00}

		for !(charIsNum(num[0]) && charIsNum(num[1])) && i <= j {
			if charIsNum(line[i]) {
				num[0] = line[i]
			} else {
				i++
			}
			if charIsNum(line[j]) {
				num[1] = line[j]
			} else {
				j--
			}
		}
		if num[0] == 0x00 {
			num[0] = num[1]
		}
		if num[1] == 0x00 {
			num[1] = num[0]
		}
		total += int(10*(num[0]-'0') + num[1] - '0')
	}
	return total
}

func part2() int {
	numbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	file, _ := os.Open("day1.txt")
	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		var i int = 0
		var j int = len(line) - 1
		num := [2]byte{0x00, 0x00}
		str := [2]string{"", ""}

		for !(charIsNum(num[0]) && charIsNum(num[1])) && i <= j {
			str[0] = line[i : i+5]
			str[1] = line[j-5 : j]

			// usar algo d'aixo
			// strings.Contains(line[i:i+5], numbers[0])
			// strings.Contains(line[j-5:j], numbers[0])
			// strings.Index()
		}
	}
	return total
}

func Test_day1(t *testing.T) {
	t.Log(part1())
	t.Log(part2())
}

/*
--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to
take a look. The Elves have even given you a map; on it, they've used stars to
mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you
need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each
day in the Advent calendar; the second puzzle is unlocked when you complete the
first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough")
and where they're even sending you ("the sky") and why your map looks mostly
blank ("you sure ask a lot of questions") and hang on did you just say the sky
("of course, where do you think snow comes from") when you realize that the
Elves are already loading you into a trebuchet ("please hold still, we need to
strap you in").

As they're making the final adjustments, they discover that their calibration
document (your puzzle input) has been amended by a very young Elf who was
apparently just excited to show off her art skills. Consequently, the Elves are
having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line
originally contained a specific calibration value that the Elves now need to
recover. On each line, the calibration value can be found by combining the
first digit and the last digit (in that order) to form a single two-digit
number.

For example:
1abc2		-> 12
pqr3stu8vwx	-> 38
a1b2c3d4e5f	-> 15
treb7uchet	-> 77
			-> 142

The calibration values of these four lines are 12, 38, 15, and 77.
Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

--- Part Two ---

Your calculation isn't quite right. It looks like some of the digits are
actually spelled out with letters: one, two, three, four, five, six, seven,
eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and
last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76.
Adding these together produces 281.

What is the sum of all of the calibration values?

*/
