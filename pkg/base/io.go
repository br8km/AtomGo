// common file system operation
package atomgo

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
)

func WriteJson(filePath string, jsonData interface{}, indent bool, perm os.FileMode) (int, error) {
	//write json data as buffer to json encoder
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)

	if indent {
		encoder.SetIndent("", "\t")
	}

	err := encoder.Encode(jsonData)
	if err != nil {
		return 0, err
	}
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, perm)
	if err != nil {
		return 0, err
	}
	length, err := file.Write(buffer.Bytes())
	if err != nil {
		return 0, err
	}
	return length, nil
}

func WriteStruct(filePath string, jsonData interface{}, indent bool, perm os.FileMode) error {
	//write struct data as buffer to json encoder
	var b []byte
	var e error
	if indent {
		b, e = json.MarshalIndent(jsonData, "", "\t")
	} else {
		b, e = json.Marshal(jsonData)
	}
	if e != nil {
		return e
	}
	return ioutil.WriteFile(filePath, b, perm)
}

func ReadLines(filePath string) ([]string, error) {
	return ReadLinesLong(filePath, 0)
}

func ReadLinesLong(filePath string, lineLength int) ([]string, error) {
	const maxCapacity = 65536
	var lines []string
	file, err := os.Open(filePath)
	if err != nil {
		return lines, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if lineLength > maxCapacity {
		buf := make([]byte, lineLength)
		scanner.Buffer(buf, lineLength)
	}
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return lines, err
		}
		lines = append(lines, scanner.Text())
	}
	return lines, err
}

func WriteLines(filePath string, textLines []string, perm os.FileMode) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range textLines {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	return nil
}

func ReadJson(filePath string, jsonData interface{}) (interface{}, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil

}

func ReadStruct() {}
