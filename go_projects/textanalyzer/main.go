package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filePath := flag.String("file", "", "путь к файлу для анализа")
	topN := flag.Int("top", 10, "количество слов в топе")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("использование: textanalyzer -file <путь>")
		os.Exit(1)
	}

	data, err := os.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		os.Exit(1)
	}

	text := string(data)
	freq := WordFrequency(text)
	top := TopWords(freq, *topN)

  fmt.Println()
	fmt.Printf("Строк:           %d\n", CountLines(text))
	fmt.Printf("Слов:            %d\n", CountWords(text))
	fmt.Printf("Символов:        %d\n", CountChars(text))
	fmt.Printf("Уникальных слов: %d\n", len(freq))
	fmt.Println()

	fmt.Printf("Топ-%d слов:\n", *topN)
	for i, wc := range top {
		fmt.Printf("%3d. %-10s — %d\n", i+1, wc.Word, wc.Count)
	}

	// Прочитай файл
	// Вызови функции анализа
	// Выведи результат
}
