package yr

import (
	"bufio"
	"os"
	"testing"
)

func testNumberOfLines(t *testing.T, filename string, expected int) {
	count, err := GetLastLine(filename)
	if err != nil {
		t.Fatalf("Feil ved telling av linjer: %v", err)
	}

	if count != expected {
		t.Errorf("Uventet antall linjer i filen %s: Forventa %d, Fikk %d", filename, expected, count)
	}
}

const inputFile = "../kjevik-temp-celsius-20220318-20230318.csv"
const outputFile = "../kjevik-temp-fahr-20220318-20230318.csv"

func TestNumberOfLines(t *testing.T) {
	//Tester hvor mange linjer, både på input(cels) filen og output(fahr) filen
	testNumberOfLines(t, inputFile, 16756)

	testNumberOfLines(t, outputFile, 16756)

}

func TestCelsiusToFahrenheitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
		{input: "-11", want: "12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("Test mislykkes, forventa %s, Fikk: %s", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheitFull(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{

		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if tc.want != got {
			t.Errorf("Test mislykkes, forventa: %s, Fikk: %s", tc.want, got)
		}
	}
}

func TestLastLineOfFiles(t *testing.T) {
	// Map of file names and expected last lines
	expected := map[string]string{
		inputFile:  "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
		outputFile: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av SiMENhol",
	}

	for filename, want := range expected {
		// Open the file
		file, err := os.Open(filename)
		if err != nil {
			t.Fatalf("Feil ved åpning av filen %q: %v", filename, err)
		}
		defer file.Close()

		// Create a scanner and set its split function to ScanLines
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		var lastLine string
		// Iterate over the lines in the file and store the last line in a variable
		for scanner.Scan() {
			lastLine = scanner.Text()
		}

		// Check that the last line matches the expected value using want
		if lastLine != want {
			t.Errorf("%q: last line = %q, want %q", filename, lastLine, want)
		}
	}
}

//en test som sjekker om gjennomsnittstemperatur er 8.56?
