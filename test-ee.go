package hello

import "fmt"

func test()  {
	fmt.Println("testfffaa")

        _, _ = os.Open("nonexistent_file.txt")
        fmt.Println("继续执行...")

        db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
        if err != nil {
                log.Fatal(err)
        }
        defer db.Close()

        username := "test' OR '1'='1"
	rows, err := db.Query("SELECT * FROM users WHERE username = ?", username)
        if err != nil {
                log.Fatal(err)
        }
        defer rows.Close()

        _, _ = os.Open("test.txt")


	_, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")


	_, _ = net.Dial("tcp", "127.0.0.1:8080")

}
