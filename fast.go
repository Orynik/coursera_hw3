package main

import (
	"bufio"
	"fmt"
	"io"
	"json/jsonpattern"
	"os"
	"strings"

	"github.com/mailru/easyjson/jlexer"
)

// Отладочная функция
// func main() {
// 	FastSearch(ioutil.Discard)
// }

//FastSearch ...
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	seenBrowsers := []string{}
	uniqueBrowsers := 0
	foundUsers := ""
	builder := strings.Builder{}
	i := -1

	// users := make([]*jsonpattern.User, 0)

	//Считывание построчно файла
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		//Начало обработки данных
		r := jlexer.Lexer{Data: sc.Bytes()}
		u := &jsonpattern.User{}
		u.UnmarshalEasyJSON(&r)
		i++
		isAndroid := false
		isMSIE := false
		for _, browser := range u.Browsers {
			if ok := strings.Contains(browser, "Android"); ok == true {
				isAndroid = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				//--------------!!!!!----------
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			} else if ok := strings.Contains(browser, "MSIE"); ok == true {
				isMSIE = true
				notSeenBefore := true
				for _, item := range seenBrowsers {
					if item == browser {
						notSeenBefore = false
					}
				}
				if notSeenBefore {
					// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
					seenBrowsers = append(seenBrowsers, browser)
					uniqueBrowsers++
				}
			}
		}
		if !(isAndroid && isMSIE) {
			continue
		}
		email := strings.Replace(u.Email, "@", " [at] ", -1)
		_, err = fmt.Fprintf(&builder, "[%d] %s <%s>\n", i, u.Name, email)
		if err != nil {
			panic(err)
		}
	}

	foundUsers = builder.String()
	fmt.Fprintln(out, "found users:\n"+foundUsers)
	fmt.Fprintln(out, "Total unique browsers", len(seenBrowsers))

	//--------------------LEGACY-----------------------------
	// for sc.Scan() {
	// 	//Начало обработки данных
	// 	r := jlexer.Lexer{Data: sc.Bytes()}
	// 	u := &jsonpattern.User{}
	// 	u.UnmarshalEasyJSON(&r)
	// 	// if err != nil {
	// 	// 	panic(err)
	// 	// }

	// 	users = append(users, u)
	// }

	// for i, user := range users {
	// 	isAndroid := false
	// 	isMSIE := false
	// 	for _, browser := range user.Browsers {
	// 		if ok := strings.Contains(browser, "Android"); ok == true {
	// 			isAndroid = true
	// 			notSeenBefore := true
	// 			for _, item := range seenBrowsers {
	// 				if item == browser {
	// 					notSeenBefore = false
	// 				}
	// 			}
	// 			if notSeenBefore {
	// 				// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
	// 				seenBrowsers = append(seenBrowsers, browser)
	// 				uniqueBrowsers++
	// 			}
	// 		} else if ok := strings.Contains(browser, "MSIE"); ok == true {
	// 			isMSIE = true
	// 			notSeenBefore := true
	// 			for _, item := range seenBrowsers {
	// 				if item == browser {
	// 					notSeenBefore = false
	// 				}
	// 			}
	// 			if notSeenBefore {
	// 				// log.Printf("SLOW New browser: %s, first seen: %s", browser, user["name"])
	// 				seenBrowsers = append(seenBrowsers, browser)
	// 				uniqueBrowsers++
	// 			}
	// 		}
	// 	}
	// 	if !(isAndroid && isMSIE) {
	// 		continue
	// 	}
	// 	email := strings.Replace(user.Email, "@", " [at] ", -1)
	// 	_, err = fmt.Fprintf(&builder, "[%d] %s <%s>\n", i, user.Name, email)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
}
