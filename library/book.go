package library

type Book struct {
	Id, Total_page, Category_id                                  int
	Release_year                                                 int `validate:"min=1980,max=2021"`
	Title, Description, Price, Thickness, Created_at, Updated_at string
	Image_url                                                    string `validate:"url"`
	Category                                                     Category
}
