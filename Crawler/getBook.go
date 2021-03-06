package Crawler

import (
	"github.com/tebeka/selenium"
	"strconv"
)

const (
	xpathTitleInit = "//*[@id=\"caracteristicas\"]/div[2]/div["
	xpathTitleEnd  = "]/div[1]"
	ISBN           = "Código ISBN"
	Date           = "Data de disponibilidade"
	Pages          = "Número de páginas"
	Year           = "Ano de publicação"
)

func getBook(driver selenium.WebDriver, bookName string) (Book, error) {
	var isbn string
	var dtDis string
	var pgs string
	var yr string

	bookStruct, err := getBookLink(driver, bookOpenLink, productSpecificationTR, bookName)
	if err != nil {
		return Book{}, err
	}

	for i := 0; i < bookStruct.Amount; i++ {
		specification, err := getTitle(driver, i)
		if err != nil {
			return Book{}, err
		}

		if specification != "" {
			position := i + 1
			switch specification {
			case ISBN:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				isbn = text
			case Date:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				dtDis = text
			case Pages:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				pgs = text
			case Year:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(position-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				yr = text
			}
		}
	}

	return Book{
		ISBN:          isbn,
		AvailableDate: dtDis,
		Pages:         pgs,
		PubYear:       yr,
		Link:          bookStruct.Link,
	}, nil

}

func getTitle(driver selenium.WebDriver, i int) (string, error) {
	xpath := xpathTitleInit + strconv.Itoa(i) + xpathTitleEnd

	titles, err := driver.FindElements(selenium.ByXPATH, xpath)
	if err != nil {
		return "", err
	}

	if len(titles) != 0 {
		textTitle, err := titles[0].Text()
		if err != nil {
			return "", err
		}

		return textTitle, nil
	}

	return "", nil
}
