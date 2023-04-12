package yr

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/uia-worker/misc/conv"
	//"github.com/SiMENhol/funtemps/conv"
)

func ConvTemperature() {

	outputFilename := "kjevik-temp-fahr-20220318-20230318.csv"

	if _, err := os.Stat(outputFilename); err == nil {
		fmt.Printf("Fil '%s' finnes allerede. Vil du generere filen paa nytt? (j/n): ", outputFilename)
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			answer := scanner.Text()
			if strings.ToLower(answer) == "j" || strings.ToLower(answer) == "ja" {
				break
			} else if strings.ToLower(answer) == "n" || strings.ToLower(answer) == "nei" {
				fmt.Println("shutting down")
				return
			} else {
				fmt.Print("Invalid answer. Do you want to regenerate the file? (j/n): ")
			}
		}
	}

	inputFile := openInputFile()
	defer inputFile.Close()

	outputFile, err := createOutputFile()
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	if scanner.Scan() {
		fmt.Fprintln(outputFile, scanner.Text())
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ";")

		if fields[3] == "" {
			continue
		}

		celsius, err := strconv.ParseFloat(fields[3], 64)

		if err != nil {
			log.Fatal(err)
		}

		fahrenheit := conv.CelsiusToFahrenheit(celsius)
		fields[3] = fmt.Sprintf("%.2f", fahrenheit)
		line = strings.Join(fields, ";")
		fmt.Fprintln(outputFile, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	footer := []string{"Data er basert paa gyldig data (per 18.03.2023)(CC BY 4.0) fra Meteorologisk institutt (MET); endringen er gjort av Majd Saleh"}
	writer := csv.NewWriter(outputFile)
	err = writer.Write(footer)
	if err != nil {
		fmt.Println("Kunne ikke skrive endelig tekst:", err)
	}
	writer.Flush()

}

func openInputFile() *os.File {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func createOutputFile() (*os.File, error) {
	outputFilePath := "kjevik-temp-fahr-20220318-20230318.csv"
	if _, err := os.Stat(outputFilePath); err == nil {
		fmt.Printf("File %s already exists. Deleting...\n", outputFilePath)
		err := os.Remove(outputFilePath)
		if err != nil {
			return nil, fmt.Errorf("could not delete file: %v", err)
		}
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not create file: %v", err)
	}
	return outputFile, nil
}

func AverageTemp() {
	file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0.0
	count := 0.0

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) < 4 {
			continue
		}

		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			continue
		}

		sum += temperature
		count++
	}

	if count > 0 {
		var unit string
		fmt.Println("Vil du ha gjennomsnittstemperaturen i Celsius eller Fahrenheit? (celsius/fahrenheit)")
		fmt.Scanln(&unit)

		if strings.ToLower(unit) == "fahrenheit" {
			average := (sum/float64(count))*1.8 + 32
			fmt.Printf("Gjennomsnittstemperaturen i Fahrenheit er: %.2f\n", average)
		} else {
			average := sum / float64(count)
			fmt.Printf("Gjennomsnittstemperaturen i celsius er: %.2f\n", average)
		}
	}
}

func ProcessLine(line string) string {
	if line == "" {
		return ""
	}
	fields := strings.Split(line, ";")
	lastField := ""
	if len(fields) > 0 {
		lastField = fields[len(fields)-1]
	}
	convertedField := ""
	if lastField != "" {
		var err error
		convertedField, err = convertLastField(lastField)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			return ""
		}
	}
	if convertedField != "" {
		fields[len(fields)-1] = convertedField
	}
	if line[0:7] == "Data er" {
		return "Data er basert paa gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Majd Saleh"
	} else {
		return strings.Join(fields, ";")
	}
}

func convertLastField(lastField string) (string, error) {
	celsius, err := strconv.ParseFloat(lastField, 64)
	if err != nil {
		return "", err
	}

	fahrenheit := conv.CelsiusToFahrenheit(celsius)

	return fmt.Sprintf("%.1f", fahrenheit), nil
}

func CountLines(inputFile string) int {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	countedLines := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			countedLines++
		}
	}
	return countedLines
}

func Average(fileName string) (float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0.0
	count := 0.0

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ";")
		if len(fields) < 4 {
			continue
		}

		temperature, err := strconv.ParseFloat(fields[3], 64)
		if err != nil {
			continue
		}

		sum += temperature
		count++
	}

	if count > 0 {
		average := sum / count
		return average, nil
	}

	return 0, errors.New("No temperature data found")
}

/**
func CelsiusToFahrenheitString(celsius string) (string, error) {
	var fahrFloat float64
	var err error
	if celsiusFloat, err := strconv.ParseFloat(celsius, 64); err == nil {
		fahrFloat = conv.CelsiusToFahrenheit(celsiusFloat)
	}
	fahrString := fmt.Sprintf("%.1f", fahrFloat)
	return fahrString, err
}

// Forutsetter at vi kjenner strukturen i filen og denne implementasjon
// er kun for filer som inneholder linjer hvor det fjerde element
// på linjen er verdien for temperaturaaling i grader celsius
func CelsiusToFahrenheitLine(line string) (string, error) {

	dividedString := strings.Split(line, ";")
	var err error

	if len(dividedString) == 4 {
		dividedString[3], err = CelsiusToFahrenheitString(dividedString[3])
		if err != nil {
			return "", err
		}
	} else {
		return "", errors.New("linje har ikke forventet format")
	}
	return strings.Join(dividedString, ";"), nil

	return "Kjevik;SN39040;18.03.2022 01:50;42.8", err
}
**/
