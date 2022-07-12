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

type Book struct {
	SearchName    string
	ISBN          string
	AvailableDate string
	Pages         string
	PubYear       string
}

func bookFounded(driver selenium.WebDriver, bookName string) (Book, error) {
	var isbn string
	var dtDis string
	var pgs string
	var yr string

	loop, err := amountSpecification(driver, bookOpenLink, productSpecificationTR)
	if err != nil {
		return Book{}, err
	}

	for i := 0; i < loop; i++ {
		specification, err := title(driver, i)
		if err != nil {
			return Book{}, err
		}

		if specification == "" {
			switch specification {
			case ISBN:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(i-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				isbn = text
			case Date:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(i-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				dtDis = text
			case Pages:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(i-1)+"]/div[2]")
				if err != nil {
					return Book{}, err
				}

				text, err := elem.Text()
				if err != nil {
					return Book{}, err
				}

				pgs = text
			case Year:
				elem, err := driver.FindElement(selenium.ByXPATH, "//*[@id=\"caracteristicas\"]/div[2]/div["+strconv.Itoa(i-1)+"]/div[2]")
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
		SearchName:    bookName,
		ISBN:          isbn,
		AvailableDate: dtDis,
		Pages:         pgs,
		PubYear:       yr,
	}, nil

}

func title(driver selenium.WebDriver, j int) (string, error) {
	xpath := xpathTitleInit + strconv.Itoa(j) + xpathTitleEnd

	titles, err := driver.FindElements(selenium.ByXPATH, xpath)
	if err != nil {
		return "", err
	}

	if len(titles) != 0 {
		title, err := titles[0].Text()
		if err != nil {
			return "", err
		}

		return title, nil
	}

	return "", nil
}