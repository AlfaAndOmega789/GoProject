package main

import "fmt"

func main() {
	//инициализация при создании
	var user map[string]string = map[string]string{
		"name":     "Vasily",
		"lastName": "Romanov",
	}
	fmt.Println(user)

	//сразу с нужной емкостью
	profile := make(map[string]string, 10)
	fmt.Println(profile)

	//кол-во элементов
	mapLength := len(user)

	fmt.Printf("%d %+v\n", mapLength, profile)

	//если ключа нет - вернет значение по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	//проверка на существование ключа
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	//пустая переменная - толькопроверяем что ключ есть
	_, mNameExist2 := user["middleName"]
	fmt.Println("mNameExist2", mNameExist2)

	//удаление ключа
	delete(user, "lastName")
	fmt.Printf("%#v\n", user)
}
