package pref_test

import (
	"fmt"

	"github.com/goark/gimei-cli/pref"
)

func ExamplePrefCode() {
	pref := pref.PrefCode("34")
	fmt.Println(pref.Name.Name)
	// Output:
	// 広島県
}

func ExamplePrefName() {
	prefs := pref.PrefName("島")
	for _, pref := range prefs {
		fmt.Println(pref.Name.Name)
	}
	// Output:
	// 福島県
	// 島根県
	// 広島県
	// 徳島県
	// 鹿児島県
}

func ExampleCityCode() {
	city := pref.CityCode("34100")
	fmt.Println(city.FullName().Name)
	// Output:
	// 広島県広島市
}

func ExampleCityName() {
	cities := pref.CityName("広島")
	for _, city := range cities {
		fmt.Println(city.FullName().Name)
	}
	// Output:
	// 北海道北広島市
	// 広島県広島市
	// 広島県呉市
	// 広島県竹原市
	// 広島県三原市
	// 広島県尾道市
	// 広島県福山市
	// 広島県府中市
	// 広島県三次市
	// 広島県庄原市
	// 広島県大竹市
	// 広島県東広島市
	// 広島県廿日市市
	// 広島県安芸高田市
	// 広島県江田島市
	// 広島県府中町
	// 広島県海田町
	// 広島県熊野町
	// 広島県坂町
	// 広島県安芸太田町
	// 広島県北広島町
	// 広島県大崎上島町
	// 広島県世羅町
	// 広島県神石高原町
}
