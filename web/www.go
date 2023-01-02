package main
import (
	"fmt"
	"os"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil{
		return nil, err
	} 
	return &Page(Title: title, Body: body), err
}

func main(){
	p1 := &Page{Title: "TestPage", Body: []byte("ini halaman contoh")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body)) 
}

// masih error :(
/*
Err Msg:

# command-line-arguments
./www.go:23:20: syntax error: unexpected : in argument list; possibly missing comma or )
*/