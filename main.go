package main

import (
	"flag"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func main() {
	// Где файлы
	dir := flag.String("dir", "pic", "Name file")

	// Высота и ширина выходных файлов
	hh := flag.Int("h", 300, "Height")
	wh := flag.Int("w", 200, "Wighth")

	flag.Parse()

	// Укажите путь к каталогу
	dirin := *dir + "/"

	files, err := ioutil.ReadDir(dirin)
	if err != nil {
		log.Fatal(err)
	}

	// цикл по каталогу
	for _, file := range files {
		if file.IsDir() {

			// Пропускаем подкаталоги
			continue
		}

		filePath := filepath.Join(dirin, file.Name())
		Resizer(filePath, uint(*hh), uint(*wh))

		fmt.Println(filePath)
		// Здесь вы можете выполнить необходимые действия с файлом
		// Например, открыть, прочитать или изменить его

		// Пример чтения содержимого файла:
		// content, err := ioutil.ReadFile(filePath)
		// if err != nil {
		//     log.Println(err)
		// }
		// fmt.Println(string(content))
	}
}

func Resizer(namefile string, hh, wh uint) {

	// Тело взодного файла
	nmf := strings.Split(namefile, ".")[0]

	// Открываем исходный файл изображения
	file, err := os.Open(namefile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Декодируем изображение
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Изменяем размер изображения
	resizedImg := resize.Resize(hh, wh, img, resize.Lanczos3)
	nameout := fmt.Sprintf("-%v-%v-out.jpg", hh, wh)
	dirout := filepath.Join("out/", nmf+nameout)

	// Создаем новый файл для сохранения измененного изображения
	//output, err := os.Create(nmf + nameout)
	output, err := os.Create(dirout)

	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	fmt.Println("OUT file :", namefile)

	// Кодируем и сохраняем измененное изображение в формате JPEG
	jpeg.Encode(output, resizedImg, nil)

	fmt.Println("Изображение успешно изменено и сохранено в output.jpg")
}

func OneFile() {

	namefile := flag.String("name", "namefile.jpg", "Name file")

	hh := flag.Int("h", 300, "Height")
	wh := flag.Int("w", 200, "Wighth")

	flag.Parse()

	// Открываем исходный файл изображения
	file, err := os.Open(*namefile)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Декодируем изображение
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Изменяем размер изображения
	resizedImg := resize.Resize(uint(*hh), uint(*wh), img, resize.Lanczos3)

	// Создаем новый файл для сохранения измененного изображения
	nameout := fmt.Sprintf(" %v_%v_out.jpg", *hh, *wh)

	output, err := os.Create(*namefile + nameout)

	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	// Кодируем и сохраняем измененное изображение в формате JPEG
	jpeg.Encode(output, resizedImg, nil)
	fmt.Println("Изображение успешно изменено и сохранено в output.jpg")
}
