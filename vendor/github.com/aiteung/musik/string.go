package musik

import (
	"math/rand"
	"strings"
	"time"
)

func NumberStringToHari(number string) (namahari string) {
	replacer := strings.NewReplacer("1", "Senin", "2", "Selasa", "3", "Rabu", "4", "Kamis", "5", "Jumat", "6", "Sabtu", "7", "Minggu")
	namahari = replacer.Replace(number)
	return namahari

}

func KelasNumberToAbjad(number string) (kelas string) {
	nomorkelas := number[len(number)-1:]
	angkatan := number[0:1]
	replacer := strings.NewReplacer("1", "A", "2", "B", "3", "C", "4", "D", "5", "E", "6", "F", "7", "G")
	abjadkelas := replacer.Replace(nomorkelas)
	kelas = angkatan + abjadkelas
	return kelas
}

func NormalizeString(message string) (m string) {
	m = strings.ToLower(message)
	return m
}

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
