package main

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Page struct, projenin temel veri yapısını temsil eder.
type Page struct {
	Text       string
	Banners    []string
	Ascii      string
	CountSlice []struct{}
}

var errorMessage = "Not Found, if nothing is found, for example templates or banners."

// main fonksiyonu, projenin ana giriş noktasıdır.
func main() {
	print(" Server is running", "standard")
	http.HandleFunc("/index.html", HomePage)
	http.HandleFunc("/result", submitHandler)
	http.HandleFunc("/download", downloadHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8045", nil)
}

// HomePage fonksiyonu, ana sayfa için HTTP isteklerini işler.
func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html", "templates/result.html", "templates/download.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := Page{
		Banners:    []string{"standard", "shadow", "thinkertoy"},
		CountSlice: make([]struct{}, 130),
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// submitHandler fonksiyonu, form gönderimlerini işler.
func submitHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Form.Get("generate") != "" {
		GenerateAsciiArt(w, r)
		return
	} else if r.Form.Get("export") != "" {
		downloadHandler(w, r)
		return
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")
	outputFileName := r.Form.Get("export")
	for _, char := range text {
		if char >= 1 && char <= 127 {
			continue
		} else {
			text = "Invalid Character."
		}
	}
	tmpl, err := template.ParseFiles("templates/download.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !(strings.Contains(outputFileName, ".txt")) {
		outputFileName = outputFileName + ".txt"
	}
	data, err := os.Create(outputFileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := Page{
		Text:       fmt.Sprintf("Dosya başarıyla Yazdırıldı: %s", outputFileName),
		CountSlice: make([]struct{}, 130),
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		fmt.Println("file could not be created")
		return
	}
	defer data.Close()
	PrintFile(text, banner, data)
}

func PrintFile(text, banner string, data *os.File) {
	textList := strings.Split(text, "\n")
	for _, char := range textList {
		if char != "" {
			ascii, _ := GetAscii(char, banner, errorMessage)
			for i, r := range ascii {
				if i < 8 {
					_, err := data.WriteString(r + "\n")
					if err != nil {
						fmt.Println("Error writing to file")
						return
					}
				}
			}
		} else {
			_, err := data.WriteString("\n")
			if err != nil {
				fmt.Println("Error adding a line to a file")
				return
			}
		}
	}
}

func GenerateAsciiArt(w http.ResponseWriter, r *http.Request) {
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")
	for _, char := range text {
		if char >= 1 && char <= 127 {
			continue
		} else {
			text = "Invalid Character."
		}
	}
	asciiArt, err := getAsciiArt(text, banner, errorMessage)
	if err != nil {
		text = errorMessage
	}
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := Page{
		Text:       text,
		Ascii:      asciiArt,
		CountSlice: make([]struct{}, 130),
	}
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getAsciiArt fonksiyonu, metni ASCII sanatına dönüştürür.
func getAsciiArt(text, bannerName, errorMessage string) (string, error) {
	result := ""
	if len(text) == 0 {
		return result, nil
	}
	wordsList := strings.Split(text, "\n")
	for _, char := range wordsList {
		if char != "" {
			ascii, err := GetAscii(char, bannerName, errorMessage)
			if err != nil {
				if err != nil {
					return errorMessage, fmt.Errorf(errorMessage)
				}
			}
			for _, line := range ascii {
				result += line + "\n"
			}
		}
	}
	return result, nil
}

// GetAscii fonksiyonu, metnin ASCII karakterlerini oluşturur.
func GetAscii(asciStr, bannerName, errorMessage string) ([]string, error) {
	lines := make([]string, 9)
	for i := 0; i < len(asciStr); i++ {

		assciiValue := (((int(asciStr[i])) - 32) * 9) + 2
		values, err := linesRead(Banner(bannerName), errorMessage, assciiValue, assciiValue+8)
		if err != nil {
			if err != nil {
				return []string{errorMessage}, fmt.Errorf(errorMessage)
			}
		}
		for index, value := range values {
			lines[index] += value
		}

	}
	return lines, nil
}

// linesRead fonksiyonu, dosyadan belirli aralıktaki satırları okur.
func linesRead(fileName, errorMessage string, beginning, finish int) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []string{errorMessage}, fmt.Errorf(errorMessage)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	numberOfLines := 0
	for scanner.Scan() {
		numberOfLines++
		if numberOfLines >= beginning && numberOfLines <= finish {
			lines = append(lines, scanner.Text())
		}
	}
	return lines, nil
}

// print fonksiyonu, metni konsola yazdırır.
func print(text, bannerName string) {
	if len(text) == 0 {
		return
	}
	wordsList := strings.Split(text, "\\n")
	for _, char := range wordsList {
		if char != "" {
			ascii, _ := GetAscii(char, bannerName, errorMessage)
			for i, line := range ascii {
				if i < 8 {
					fmt.Println(line)

				}
			}
		} else {
			fmt.Println()
		}
	}
}

// Banner fonksiyonu, banner dosya yolunu döndürür.
func Banner(bannerName string) string {
	if bannerName == "standard" {
		return "standard.txt"
	} else if bannerName == "shadow" {
		return "shadow.txt"
	} else if bannerName == "thinkertoy" {
		return "thinkertoy.txt"
	} else {
		return ""
	}
}
