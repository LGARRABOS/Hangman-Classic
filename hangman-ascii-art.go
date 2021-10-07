package piscine

import "fmt"

func Ascii_art_min(tabunderscore string) {
	var ligne1 string
	var ligne2 string
	var ligne3 string
	var ligne4 string
	var ligne5 string
	var ligne6 string
	var ligne7 string
	tabl1 := []string{"        ","        ","       ","        ","       ","      ","        ","        ","    ","   _  ","        ","    ","            ","        ","        ","        ","        ","       ","       ","      ","        ","         ","              ","         ","        ","      ","        "}
	tabl2 := []string{"        "," _      ","       ","     _  ","       ","  __  ","  __ _  "," _      "," _  ","  (_) ","        "," _  ","            ","        ","        "," _ __   ","  __ _  ","       ","       "," _    ","        ","         ","              ","         "," _   _  ","      ","        "}
	tabl3 := []string{"        ","| |     ","       ","    | | ","       "," / _| "," / _` | ","| |     ","(_) ","   _  "," _      ","| | ","            ","        ","        ","| '_ \\ "," / _` | ","       ","       ","| |   ","        ","         ","              ","         ","| | | | ","      ","        "}
	tabl4 := []string{"  __ _  ","| |__   ","  ___  ","  __| | ","  ___  ","| |_  ","| (_| | ","| |__   "," _  ","  | | ","| | _   ","| | "," _ __ ___   "," _ __   ","  ___   ","| |_) | ","| (_| | "," _ __  "," ___   ","| |_  "," _   _  ","__   __  ","__      __    ","__  __   ","| |_| | "," ____ ","        "}
	tabl5 := []string{" / _` | ","| '_ \\ "," / __| "," / _` | "," / _ \\","|  _| "," \\__, |","|  _ \\ ","| | ","  | | ","| |/ /  ","| | ","| '_ ` _ \\ ","| '_ \\ "," / _ \\ ","| .__/  "," \\__, |","| '__| ","/ __|  ","| __| ","| | | | ","\\ \\ / /","\\ \\ /\\ / / ","\\ \\/ / "," \\__, |","|_  / ","        "}
	tabl6 := []string{"| (_| | ","| |_) | ","| (__  ","| (_| | ","|  __/ ","| |   ","  __/ | ","| | | | ","| | "," _/ | ","|   <   ","| | ","| | | | | | ","| | | | ","| (_) | ","| |     ","    | | ","| |    ","\\__ \\","\\ |_ ","| |_| | "," \\ V /  "," \\ V  V /    "," >  <    "," __/ /  "," / /  "," ______ "}
	tabl7 := []string{" \\__,_|","|_.__/  "," \\___|"," \\__,_|"," \\___|","|_|   "," |___/  ","|_| |_| ","|_| ","|__/  ","|_|\\_\\","|_| ","|_| |_| |_| ","|_| |_| "," \\___/ ","|_|     ","    |_| ","|_|    ","|___/  "," \\__|"," \\__,_|","  \\_/   ","  \\_/\\_/    ","/_/\\_\\ ","|___/   ","/___| ","|______|"}

	for j := 0; j < len(tabunderscore); j++ {
		for i := 97; i <= 122; i++ {
			if tabunderscore[j] == i {
				ligne1 = ligne1 + tab1[i-97]
				ligne2 = ligne2 + tab2[i-97]
				ligne3 = ligne3 + tab3[i-97]
				ligne4 = ligne4 + tab4[i-97]
				ligne5 = ligne5 + tab5[i-97]
				ligne6 = ligne6 + tab6[i-97]
				ligne7 = ligne7 + tab7[i-97]
			} else if tabunderscore[] == '_' {
				ligne1 = ligne1 + tab1[26]
				ligne2 = ligne2 + tab2[26]
				ligne3 = ligne3 + tab3[26]
				ligne4 = ligne4 + tab4[26]
				ligne5 = ligne5 + tab5[26]
				ligne6 = ligne6 + tab6[26]
				ligne7 = ligne7 + tab7[26]
			}
		}
	}
	fmt.Println(ligne1)
	fmt.Println(ligne2)
	fmt.Println(ligne3)
	fmt.Println(ligne4)
	fmt.Println(ligne5)
	fmt.Println(ligne6)
	fmt.Println(ligne7)
}
