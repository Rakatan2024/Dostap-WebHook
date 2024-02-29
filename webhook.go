package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

//func verifySignature(message, signature []byte, secret string) bool {
//	mac := hmac.New(sha256.New, []byte(secret))
//	mac.Write(message)
//	expectedMAC := mac.Sum(nil)
//	return hmac.Equal(signature, expectedMAC)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		//fmt.Println("Zashel v handler")
		cmd2 := exec.Command("./Dostap-Backend/deploy.sh")
		err2 := cmd2.Run()
		cmd2.Stdout = os.Stdout
		cmd2.Stderr = os.Stderr
		//fmt.Println("after cmd2 run")
		if err2 != nil {
			//fmt.Println("KEK ERR2")
			//fmt.Println(err2)
			http.Error(w, "Failed to execute deploy script", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		//fmt.Println("VSE OKEY")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		//fmt.Println("Method not allowed")
	}
}

func main() {
	http.HandleFunc("/", handler)
	//fmt.Println("Listening on: http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("KEK:", err)
	}
}
