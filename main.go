package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fungsi untuk mengubah urutan kata
func rearrangeSentence(input string) string {
	words := strings.Fields(input)

	// Cari merek dalam kalimat
	for i, word := range words {
		if brands[strings.ToLower(word)] {
			// Gabungkan kata sebelum merek sebagai jenis produk atau deskripsi
			rest := strings.Join(words[:i], " ")
			// Kembalikan urutan yang diinginkan
			return fmt.Sprintf("%s %s", word, rest)
		}
	}

	// Jika tidak ada merek yang ditemukan, kembalikan kalimat asli
	return input
}

// Fungsi untuk mengganti spasi dengan karakter tertentu dalam frasa tertentu
func replaceSpaces(input, replacement string) string {
	for phrase := range spasi {
		if strings.Contains(strings.ToLower(input), strings.ToLower(phrase)) {
			modifiedPhrase := strings.ReplaceAll(phrase, " ", replacement)
			input = strings.ReplaceAll(input, phrase, modifiedPhrase)
		}
	}
	return input
}

// Fungsi untuk mengganti awalan kata tertentu dengan =
func replacePrefixes(input string) string {
	words := strings.Fields(input)
	var result strings.Builder

	for i := 0; i < len(words); i++ {
		trimmedWord := words[i]

		if gantidepan[strings.ToLower(trimmedWord)] {
			// Tambahkan = dan kata yang diubah
			result.WriteString("=" + trimmedWord)
		} else {
			// Jika bukan kata yang diganti, tambahkan kata tersebut ke hasil
			result.WriteString(trimmedWord)
		}

		// Tambahkan spasi jika bukan kata terakhir
		if i < len(words)-1 {
			result.WriteString(" ")
		}
	}

	return result.String()
}

func main() {
	// Buat file output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Variable untuk karakter pengganti spasi
	replacementChar := "="

	// Proses setiap kalimat dalam sentences
	for _, sentence := range sentences {
		modified := rearrangeSentence(sentence)
		modifiedWithReplacement := replaceSpaces(modified, replacementChar)
		finalResult := replacePrefixes(modifiedWithReplacement)
		_, err := writer.WriteString(finalResult + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Pastikan semua data ditulis ke file
	writer.Flush()
}