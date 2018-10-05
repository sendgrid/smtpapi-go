If you have a non-library SendGrid issue, please contact our [support team](https://support.sendgrid.com).

If you can't find an issue below, please open an [issue](https://github.com/sendgrid/smtpapi-go/issues).


## Table of Contents
 * [Viewing the Request Body](#request-body)

<a name="request-body"></a>
## Viewing the request body

If you are having issues with the SMTPAPI, or it is not working in the way you expect,
viewing the instructions that are being sent by you, is a good place to start your troubleshooting.

You can view the headers that we are passing to the SendGrid API in the following way.

```golang
  package  main

  import (
    "fmt"
    "github.com/sendgrid/smtpapi-go"
  )

  func main(){
    //instantiate NewSMTPAPIHeader
    header := smtapi.NewSMTPAPIHeader()

    //Set some value
    header.AddCategory("NewUser")

    fmt.Println(header.JSONString())  
  }
  ```

  ```json
    {"category":["NewUser"]} <nil>
  ```

  Now you can ensure that your headers will add up to your desired behaviour.
