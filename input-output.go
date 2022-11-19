/************************************************************
 * Miscellaneous Go functions                               *
 * Time-stamp: <2022-11-06 11:43:16 christophe@pallier.org> *
 ************************************************************/

package goutils

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"encoding/csv"
)




// ReadFromCSV reads a .csv file and returns it as a [][]string
func ReadCSV(fname string) (dataframe [][]string, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	dataframe, err = reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return dataframe, nil
}

// WriteCSV takes a [][]string as saves it in a text file.
// Comma is the field delimiter.
func WriteCSV(table [][]string, fname string) error {
	f, err := os.Create(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	
	writer := csv.NewWriter(f)
	err = f.WriteAll(table)
	return err
}

// ReadLines reads lines from text input, returns as a []string.
// Example: 
//
//      xs, err := ReadLines(bufio.NewScanner(os.Stdin))
//
func ReadLines(s *bufio.Scanner) ([]string, error) {
	data := make([]string, 0, 1024)
	
	for s.Scan() {
		data = append(data, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}



// ReadWords reads words (split at whitespaces) from a text input, returns a []string
// Example: 
//
//      xs, err := ReadWords(bufio.NewScanner(os.Stdin))
//
func ReadWords(s *bufio.Scanner) ([]string, error) {
	data := make([]string, 0, 1024)
	
	s.Split(bufio.ScanWords)
	
	for s.Scan() {
		data = append(data, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}




// ReadFloats reads float64 numbers from text input, return as []float64
// Example: to read floats from standard input
//
//      xs, err := ReadFloats(bufio.NewScanner(os.Stdin))
//
func ReadFloats(s *bufio.Scanner) ([]float64, error) {
	data := make([]float64, 0, 1024)
	for s.Scan() {
		txt := s.Text()
		for _, tok := range strings.Fields(txt) {
			if x, err := strconv.ParseFloat(tok, 64); err != nil {
				return nil, err
			} else {
				data = append(data, x)
			}
		}

	}
	return data, nil
}
