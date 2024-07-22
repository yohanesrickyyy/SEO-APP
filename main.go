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

func main() {
	// Buat file output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Proses setiap kalimat dalam sentences
	for _, sentence := range sentences {
		modified := rearrangeSentence(sentence)
		_, err := writer.WriteString(modified + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	// Pastikan semua data ditulis ke file
	writer.Flush()
}
