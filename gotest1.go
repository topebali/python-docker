doc := Document{}
if doc.Error != nil {       // <- here
    fmt.Println(doc.Error)  // <- and here 
}
