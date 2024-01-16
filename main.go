package main

import (
	"fmt"
)

// type Car struct{
// 	Brand string
// 	Model string
// 	Price float64
// }

// func getCarAsText(car Car) string {
// 	carform := fmt.Sprintf("%s %s Price: %0.2f USD",car.Brand, car.Model, car.Price)

// 	return carform

// }

// func main(){
// 	car := Car{
// 		Brand: "Audi",
// 		Model: "A5",
// 		Price: 51999.9999,
// 	}

// 	fullCar := getCarAsText(car)
// 	fmt.Println(fullCar)
// }

// type Movie struct {
// 	Name      string
// 	Price     float64
// 	SizeBytes int
// }

// func main(){
// 	 availableMovies := []Movie{
// 		{
// 			Name: "The Meg",
// 			Price: 10.50,
// 			SizeBytes: 20000,
// 		},
// 		 {
// 			Name: "The Green Mile",
// 			Price: 6.890,
// 			SizeBytes: 30000,
// 		},
// 		{
// 			Name: "Swat Team",
// 			Price:  12.50,
// 			SizeBytes: 25000,
// 		},
// 		{
// 			Name: "Barbie",
// 			Price: 15.00,
// 			SizeBytes: 30000,
// 		},
// 		{
// 			Name: "Flash",
// 			Price: 5.50,
// 			SizeBytes: 45000,
// 		},
// 	}

// 	movieSold := []Movie{availableMovies[3], availableMovies[4]}

// 	turnOver := getSumOfMovies(movieSold)
// 	list := getListOfSoldMovies(movieSold)
// 	byteSum := getTotalSizeOfBytes(availableMovies)
// 	allMovies := getAllMovieNames(availableMovies)

// 	//fmt.Printf("Movies list: %s\n",allMovies)
// 	fmt.Printf("Movies list: %v\n",strings.Join(allMovies, ","))
// 	fmt.Printf("Turnover: %.2f GBP, total sold count: %d, used disk size: %d bytes\n", turnOver, list,byteSum)
// }

// func getSumOfMovies(movieSold []Movie)float64{
// 	turnOver := 0.0
// 	for _, currentMovie := range movieSold {
// 		turnOver = turnOver + currentMovie.Price
// 	}
// 	return turnOver

// }

// func getListOfSoldMovies(availableMovies []Movie)int{
// 	return len(availableMovies)
// }

// func getTotalSizeOfBytes(movies []Movie)int{
// 	byteSum := 0
// 	for _, currentMovie := range movies{
// 		byteSum = byteSum + currentMovie.SizeBytes
// 	}
// 	return byteSum
//}

// func getAllMovieNames(availableMovies []Movie)[]string{
// 	allMovies := []string{}
// 	for _, currentMovie := range availableMovies {
// 		allMovies = append(allMovies, currentMovie.Name)

// 	}
// 	return allMovies
// }

//const mapItemCount = x

type Apartment struct {
	RoomsCount int
	Area       float64
	District   string
	isRented   bool
	Name       string

}

const comfortPrice = 5000

func main(){
	apartments := []Apartment{
		{
			Name: "lux",
			RoomsCount: 5,
			Area: 100.2,
			District: "area1",
			isRented: false,
		},
		{
			Name: "lux2",
			RoomsCount: 3,
			Area: 88.9,
			District: "area1",
			isRented: true,
		},
		{
			Name: "sm",
			RoomsCount: 1,
			Area: 25.3,
			District: "area2",
			isRented: true,
		},
		{
			Name: "md",
			RoomsCount: 2,
			Area: 44.3,
			District: "area3",
			isRented: false,
		},
		{
			Name: "md2",
			RoomsCount: 2,
			Area: 56.0,
			District: "area4",
			isRented: false,
		},
		{
			Name: "lux3",
			RoomsCount: 4,
			Area: 118.0,
			District: "area5",
			isRented: true,
		},

	}
	for _, apartment := range apartments {
		pricePerSqFt := apartmentPrice(apartment.District,apartment.RoomsCount)
		totalPriceBasedOnArea := calculateApartmentPrice(apartment.Area, pricePerSqFt)
		totalPriceBasedOnArea = reduceApartmentPrice(apartment.isRented, totalPriceBasedOnArea) 
		areaComfortIndex := calculateAreaComfortIndex(apartment.District)

		totalPriceBasedOnArea = areaComfortIndex * comfortPrice + totalPriceBasedOnArea
		fmt.Printf("Selling price of %q is %.2f\n",apartment.Name, totalPriceBasedOnArea,)
	}
	

}

func calculateAreaComfortIndex(districtName string)float64{
	// if districtName == "area1" {
	// 	return 1
	// }
	// if districtName == "area2" {
	// 	return 0.6
	// }
	// if districtName == "area3" {
	// 	return 0.5
	// }
	// if districtName == "area4" {
	// 	return 0.1
	// }

	// return 0	
	
	switch districtName {
	case "area1":
		return 1.0
	case "area2":
		return 0.6
	case "area3":
		return 0.5
	case "area4":
		return 0.1
	default:
		return 0
	}
}

func apartmentPrice(district string, roomCount int)float64 {
	if (district == "area1" || district == "area2") && roomCount <  3 {
		return 10000
	}
	if ( district == "area1" || district == "area2") && roomCount > 3 {
		return 11000
	}
	if district == "area3" {
		return 7000
	}
	if district == "area4" {
		return 5000
	}
	return 3000
   
}

func reduceApartmentPrice(isRented bool, price float64)float64{
	if isRented {
		return price * 0.2
	}
	return price
}

func calculateApartmentPrice(price float64, area float64) float64 {
	return price * area
}