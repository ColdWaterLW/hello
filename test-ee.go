package hello

import "fmt"

func test()  {
	fmt.Println("testfffaa")

        file, err = os.Open("test2222.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

        db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

	var b *string
	fmt.Println(*b)
	
        username := "test"
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
        if err != nil {
                log.Fatal(err)
        }
        defer rows.Close()

        file, err = os.Open("test.txt")
        if err != nil {
                log.Fatal(err)
        }
	defer file.Close()

	db, err := sql.Open("mysql", dsn)
        if err != nil {
                log.Fatal(err)
        }
	defer db.Close()


	conn, err = net.Dial("tcp", "127.0.0.1:8080")
        if err != nil {
                log.Fatal(err)
        }
	defer conn.Close()

}
