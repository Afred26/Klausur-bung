package Projekt_Euler

//If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.
//If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
//Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters.
//The use of "and" when writing out numbers is in compliance with British usage.

func NumberString6Digits(n int) string {

	result := ""
	if n == 0 {
		return "null"
	}
	low := n % 1000
	high := n / 1000
	if high != 0 {
		result += NumberString3Digits(high) + "tausend"
	}
	if low != 0 {
		result += NumberString3Digits(low)
	}
	return result
}

func NumberString3Digits(n int) string {
	if n <= 19 {
		return NumberStringBelow20(n)
	}
	return NumberStringGreater20(n)
}

func NumberStringBelow20(n int) string {
	switch n {
	case 0:
		return "zero"
	case 1:
		return "one"
	case 2:
		return "tow"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	}
	return ""
}

func NumberStringGreater20(n int) string {
	if n%100 < 20 && n%100 > 0 {
		return DigitString100(n/100) + NumberStringBelow20(n%100)
	}
	return DigitString100(n/100) + DigitString1(n%10) + DigitString10((n%100)/10)
}

func DigitString100(digit int) string {
	return []string{"", "onehunderd", "towhunderd", "threehunderd", "fourhunderd", "fivehunderd", "sichshunderd", "sevenhunderd", "eighthunderd", "ninehunderd"}[digit]
}

func DigitString10(digit int) string {
	return []string{"", "", "zwent", "dreißig", "vierzig", "fünfzig", "sechzig", "siebzig", "achtzig", "neunzig"}[digit]

}

func DigitString1(digit int) string {
	return []string{"", "einund", "zweiund", "dreiund", "vierund", "fünfund", "sechsund", "siebenund", "achtund", "neunund"}[digit]
}