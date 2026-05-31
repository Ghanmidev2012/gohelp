  <img src ="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQfivIz6Gvc5puPeT9lUSroVMXzciuiPyqFBA&s "> 

# 🚀 gohelp
gohelp is a lightweight, open-source Go package designed to simplify core terminal operations, print custom greetings, generate welcome messages, and show live workspace system times.

---

## 🛠️ How to Work

Open your terminal inside **VS Code** and run the following initialization setup commands step-by-step:

1. **Initialize your local Go module:**
   ```bash
   go mod init test-my-package
   ```
   2 **Import gohelp:**
   * In gohelp is have 5 version
   ```bash
   go get github.com/Ghanmidev2012/gohelp/helper@v1.0.0
   ```
   * Version 2:
   ```bash
    go get github.com/Ghanmidev2012/gohelp/helper@v1.1.1
     ```
     * Version 3:
     ```bash
       go get github.com/Ghanmidev2012/gohelp/helper@v1.2.0
     ```
     * Version 4:
       ```bash
       go get github.com/Ghanmidev2012/gohelp/helper@v1.2.1
       ```
       # Exemples
       1. **Say hello:**
          ```go
          package main

          import "[github.com/Ghanmidev2012/gohelp/helper](https://github.com/Ghanmidev2012/gohelp/helper)"

          func main() {
          helper.SayHello("Adam")
          }
          ```
          2. **Say message
             ```go
             package main

              import (
             "fmt"
             "[github.com/Ghanmidev2012/gohelp/message](https://github.com/Ghanmidev2012/gohelp/message)"
             )

              func main() {
              msg := message.WelcomeMessage("adam")
               fmt.Println(msg)
              }    

               ```
             3. **Time
                ```go
                package main

                import "[github.com/Ghanmidev2012/gohelp/time](https://github.com/Ghanmidev2012/gohelp/time)"

                func main() {
                time.Time()
                }
                ```
                4. **Input
                   ```go
                     package main

                    import (
	
                	"github.com/Ghanmidev2012/gohelp/helper"
	
	
                    )

                    func main(){
                      	helper.Input("type")
                    }
                   ```
     
         
