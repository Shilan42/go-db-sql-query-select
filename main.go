package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// Sale представляет собой структуру, описывающую продажу товара
type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

// selectSales выполняет выборку всех продаж для указанного клиента
func selectSales(client int) ([]Sale, error) {
	// Открываем подключение к базе данных
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println("Ошибка при подключении к БД:", err)
		return nil, err
	}
	// Закрываем подключение после завершения работы
	defer db.Close()

	// Слайс для хранения результатов
	var sales []Sale

	// Выполняем SQL-запрос с параметрами product, volume, date из таблицы sales
	rows, err := db.Query("SELECT product, volume, date FROM sales WHERE client = :client", sql.Named("client", client))
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	// Закрываем результат запроса
	defer rows.Close()

	// Парсим результаты запроса и записываем в структуру sale
	for rows.Next() {
		sale := Sale{}
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
			fmt.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		// Добавляем результат в слайс
		sales = append(sales, sale)
	}

	return sales, nil
}

func main() {
	// ID клиента для выборки
	client := 208

	// Получаем продажи для указанного клиента
	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Выводим результаты
	for _, sale := range sales {
		fmt.Println(sale)
	}
}
