package utils

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(num int64) string {
	if num == 0 {
		return "0"
	}

	result := ""
	for num > 0 {
		rem := num % 62
		result = string(alphabet[rem]) + result
		num = num / 62
	}
	return result
}
