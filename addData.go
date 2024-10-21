package addData

import "fmt"

type Data struct {
	ID   int
	Name string
}

func AddData(dataList *[]Data, id int, name string, done chan bool) {
	data := Data{
		ID:   id,
		Name: name,
	}

	*dataList = append(*dataList, data)
	done <- true
}

func LogData(data Data, done chan bool) {
	fmt.Printf("Data added: ID: %d, Name: %s\n", data.ID, data.Name)
	done <- true
}

func ProcessData(dataInput []Data) {
	var dataList []Data
	done := make(chan bool, len(dataInput))

	for _, data := range dataInput {
		go AddData(&dataList, data.ID, data.Name, done)
	}

	for range dataInput {
		<-done
	}

	for _, data := range dataList {
		go LogData(data, done)
	}

	for range dataList {
		<-done
	}

	fmt.Printf("Semua data berhasil ditambahkan.\nTotal Data: %d\n", len(dataInput))
}
