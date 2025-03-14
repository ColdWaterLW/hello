package hello

import "fmt"

func test()  {
	fmt.Println("testfffaa")

        _, _ = os.Open("test2222.txt")

        db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        username := "test"
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
        if err != nil {
                log.Fatal(err)
        }
        defer rows.Close()

        _, err = os.Open("test.txt")
        if err != nil {
                log.Fatal(err)
        }


	_, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
        if err != nil {
                log.Fatal(err)
        }

        var ptr *int
        *ptr = 10

	_, err = net.Dial("tcp", "127.0.0.1:8080")
        if err != nil {
                log.Fatal(err)
        }

	var i *int
	*i=0
}
