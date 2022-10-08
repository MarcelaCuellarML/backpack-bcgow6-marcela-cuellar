package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

func NewStore(pathFile string, fileName string) Store {
	return &fileStore{fileName}
}

type Type string

const (
	FileType Type = "file"
	//MonoType Type = "mongo"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

// func leerJson() []byte {
// 	readFile, err := os.ReadFile("/Users/marcuellar/backpack-bcgow6-marcela-cuellar/GoWeb1/Clase2/internal/products/productos.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	//listado := fmt.Sprint(string(readFile))
// 	return readFile
// }

// func TransformData(listado []byte) (product []products) {
// 	err := json.Unmarshal(listado, &product)
// 	if err != nil {
// 		panic(err)
// 	}
// 	//fmt.Println("archivo prueba: ", product)
// 	return product
// }

// func GetList() (prod []products) {
// 	respuesta := leerJson()
// 	listadoProds := TransformData(respuesta)
// 	return listadoProds
// }
