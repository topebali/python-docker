doc := Document{}
if doc.Error != nil {       // <- here
    fmt.Println(doc.Error)  // <- and here 
}

doc := Document{}
if Document.Error != nil {       // <- here
    fmt.Println(Document.Error)  // <- and here 
}
