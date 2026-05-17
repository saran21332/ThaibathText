package converter

import (
	"strings"

	"github.com/shopspring/decimal"
)

func ThaiBahtText(amount decimal.Decimal) string {
	baht, satang := SplitBahtAndSatang(amount)
	bahtText := ConvertNumberToThaiText(baht)
	if satang == "00" {
		return bahtText + "บาทถ้วน"
	}
	satangText := ConvertNumberToThaiText(satang)
	return bahtText + "บาท" + satangText + "สตางค์"
}

func SplitBahtAndSatang(amount decimal.Decimal) (string, string) {
	amountText := amount.StringFixed(2)
	parts := strings.Split(amountText, ".")
	bahtPart := parts[0]
	satangPart := parts[1]
	return bahtPart, satangPart
}

func ConvertDigitsToThaiText(number string) string {
	thaiText := ""
	for i := 0; i < len(number); i++ {
		digit := int(number[i] - '0')
		if digit == 0 {
			continue
		}
		position := len(number) - i - 1
		if position == 1 && digit == 1 {
			thaiText += "สิบ"
			continue
		}
		if position == 1 && digit == 2 {
			thaiText += "ยี่สิบ"
			continue
		}
		if position == 0 && digit == 1 && len(number) > 1 {
			thaiText += "เอ็ด"
			continue
		}
		thaiText += digitTexts[digit]
		thaiText += positionTexts[position]
	}
	return thaiText
}

func ConvertNumberToThaiText(number string) string {
	if number == "0" {
		return "ศูนย์"
	}
	if len(number) <= 6 {
		return ConvertDigitsToThaiText(number)
	}
	left := number[:len(number)-6]
	right := number[len(number)-6:]
	leftText := ConvertNumberToThaiText(left)
	rightText := ConvertDigitsToThaiText(right)
	return leftText + "ล้าน" + rightText
}
