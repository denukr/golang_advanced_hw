package json

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"strings"
)

// Data описывает структуру отдельной записи
type Data struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

// JsonStorage управляет чтением и записью данных в JSON-файл
type JsonStorage struct {
	Data     []Data         `json:"data"`
	index    map[string]int `json:"-"` // Внутренний индекс для быстрого поиска
	fileName string
}

// NewJsonStorage — конструктор. Принимает путь к файлу.
func NewJsonStorage(fileName string) *JsonStorage {
	return &JsonStorage{
		Data:     []Data{},
		index:    make(map[string]int),
		fileName: fileName,
	}
}

// buildIndex сканирует слайс Data и заполняет карту для быстрого поиска по Email.
// Используем strings.ToLower, чтобы поиск не зависел от регистра букв.
func (storage *JsonStorage) buildIndex() {
	storage.index = make(map[string]int)
	for i, item := range storage.Data {
		cleanEmail := strings.ToLower(strings.TrimSpace(item.Email))
		storage.index[cleanEmail] = i
	}
}

func (storage *JsonStorage) Read() error {
	file, err := os.Open(storage.fileName)
	if err != nil {
		if os.IsNotExist(err) {
			// Если файла нет — это не ошибка, просто данных пока 0
			return nil
		}
		return err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// ЕСЛИ ФАЙЛ ПУСТОЙ — инициализируем пустые данные и выходим без ошибки
	if len(data) == 0 {
		storage.Data = []Data{}
		storage.index = make(map[string]int)
		return nil
	}

	storage.Data = []Data{}
	err = json.Unmarshal(data, storage)
	if err != nil {
		return err // Здесь будет ошибка, если в файле мусор вместо JSON
	}

	storage.buildIndex()
	return nil
}

// Write добавляет новую запись или обновляет существующую (по Email).
func (storage *JsonStorage) Write(newData Data) {
	// 1. Читаем файл, чтобы иметь актуальные данные и индекс в памяти
	storage.Read()

	// 2. Приводим email к нижнему регистру для надежного сравнения
	cleanEmail := strings.ToLower(strings.TrimSpace(newData.Email))

	// 3. Проверяем наличие Email через индекс
	if idx, exists := storage.index[cleanEmail]; exists {
		// Если найден — обновляем хеш в существующей записи
		storage.Data[idx].Hash = newData.Hash
	} else {
		// Если не найден — добавляем в конец списка
		storage.Data = append(storage.Data, newData)
		// Сразу обновляем индекс для новой записи
		storage.index[cleanEmail] = len(storage.Data) - 1
	}

	// 4. Сериализуем и сохраняем
	fileData, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		return
	}

	err = os.WriteFile(storage.fileName, fileData, 0644)
	if err != nil {
		log.Printf("Ошибка при записи в файл: %v", err)
	}
}

func (storage *JsonStorage) WriteAll() {
	fileData, err := json.MarshalIndent(storage, "", "  ")
	if err != nil {
		log.Printf("Ошибка сериализации JSON: %v", err)
		return
	}
	err = os.WriteFile(storage.fileName, fileData, 0644)
	if err != nil {
		log.Printf("Ошибка при записи в файл: %v", err)
	}
}

func (storage *JsonStorage) Remove(hash string) {
	for _, v := range storage.Data {
		if v.Hash == hash {
			storage.Data = append(storage.Data[:storage.index[v.Email]], storage.Data[storage.index[v.Email]+1:]...)
			break
		}
	}
	storage.WriteAll()
}
