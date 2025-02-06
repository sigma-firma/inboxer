[//]: # (Copyright [c] 2025 sigma-firma)

[//]: # (Permission is hereby granted, free of charge, to any person obtaining a copy)
[//]: # (of this software and associated documentation files [the "Software"], to deal)
[//]: # (in the Software without restriction, including without limitation the rights)
[//]: # (to use, copy, modify, merge, publish, distribute, sublicense, and/or sell)
[//]: # (copies of the Software, and to permit persons to whom the Software is)
[//]: # (furnished to do so, subject to the following conditions:)

[//]: # (The above copyright notice and this permission notice shall be included in all)
[//]: # (copies or substantial portions of the Software.)

[//]: # (THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR)
[//]: # (IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,)
[//]: # (FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE)
[//]: # (AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER)
[//]: # (LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,)
[//]: # (OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE)
[//]: # (SOFTWARE.)

# Example Usage

### Validating credentials and connecting to the API:

It's important to note that for this module to work properly, you need to 
**enable the sheets and Gmail API(s) in Google Cloud Services**, and download the 
**credentials.json** file provided in the **APIs and Services** section of the
Google Cloud console.

If you're unsure how to do any of that or have never used a Google Service API 
such as the SheetsAPI or GmailAPI, please see the following link:

https://developers.google.com/sheets/api/quickstart/go

That link will walk you through enabling the sheets API through the Google 
Cloud console, and creating and downloading your `credentials.json` file.

Once you have enabled the API, download the `credentials.json` file and store 
somewhere safe. You can connect to the Gmail and Sheets APIs using the 
following:

```
func main() {
	var access *Access = NewAccess(
		os.Getenv("HOME")+"/credentials/credentials.json",
		os.Getenv("HOME")+"/credentials/quickstart.json",
		[]string{
			gmail.GmailComposeScope,
			sheets.SpreadsheetsScope,
		})
	access.ReadCredentials()
	access.Connect(&gmail.Service{})
	access.Connect(&sheets.Service{})
	fmt.Println(access)
}

```

### Reading values from a spreadsheet:

```
func main() {                                                                          
        package main

        import (
                "fmt"
                "log"

                "github.com/hartsfield/ohsheet"
        )

        // Connect to the API                                                          
        sheet := &ohsheet.Access{                                                      
                Token:       "token.json",                                             
                Credentials: "credentials.json",                                       
                // You may want a ReadOnly scope here instead
                Scopes:      []string{"https://www.googleapis.com/auth/spreadsheets"}, 
        }                                                                              
        srv := sheet.Connect()                                                         


        // Prints the names and majors of students in a sample spreadsheet:
        // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
        spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"
        readRange := "Class Data!A2:E"
        
        resp, err := sheet.Read(srv, spreadsheetId, readRange)
        if err != nil {
                fmt.Fatalf("Unable to retrieve data from sheet: %v", err)
        }

        if len(resp.Values) == 0 {
                fmt.Println("No data found.")
        } else {
                fmt.Println("Name, Major:")
                for _, row := range resp.Values {
                        // Print columns A and E, which correspond to indices 0 and 4.
                        fmt.Printf("%s, %s\n", row[0], row[4])
                }
        }
}
```

### Writing values to a spreadsheet:

```
package main

import (
        "fmt"
        "log"

        "github.com/hartsfield/ohsheet"
)

func main() {                                                                          
        // Connect to the API                                                          
        sheet := &ohsheet.Access{                                                      
                Token:       "token.json",                                             
                Credentials: "credentials.json",                                       
                Scopes:      []string{"https://www.googleapis.com/auth/spreadsheets"}, 
        }                                                                              
        srv := sheet.Connect()                                                         

        spreadsheetId := "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms"

        // Write to the sheet
        writeRange := "K2"
        data := []interface{}{"test data 3"}
        res, err := sheet.Write(srv, spreadsheetId, writeRange, data)
        if err != nil {
                log.Fatalf("Unable to retrieve data from sheet: %v", err)
        }
        fmt.Println(res)
}
`

A go package for managing gmail Tokens/credentials. 

# USE:


	func main() {
			// Connect to the Gmail API service. Here, we use a context and provide a
			// scope. The scope is used by the Gamil API to determine your privilege
			// level. gmailAPI.ConnectToService is a variadic function and thus can be
			// passed any number of scopes. For more information on scopes see:
			// https://developers.google.com/gmail/api/auth/scopes
			ctx := context.Background()
			srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

			// Get a list of your unread messages
			inbox, err := srv.Users.Messages.List("me").Q("in:UNREAD").Do()
			if err != nil {
					fmt.Println(err)
			}

			for _, message := range inbox.Messages {
					// To get the message content, you must make a second call
					// to the gmail API for each individual ID.
					msg, _ := srv.Users.Messages.Get("me", message.Id).Do()
					fmt.Println(msg.Snippet)
			}
	}


#### Turning on the gmail API

  - Use this wizard (https://console.developers.google.com/start/api?id=gmail) 
  to create or select a project in the Google Developers Console and 
  automatically turn on the API. Click Continue, 
  then Go to credentials.
 
  - On the Add credentials to your project page, click the Cancel button.
 
  - At the top of the page, select the OAuth consent screen tab. Select an 
  Email address, enter a Product name if not already set, and click the Save 
  button.
 
  - Select the Credentials tab, click the Create credentials button and select
  OAuth client ID.
 
  - Select the application type Other, enter the name "Gmail API Quickstart", 
  and click the Create button.
 
  - Click OK to dismiss the resulting dialog.
 
  - Click the file_download (Download JSON) button to the right of the client 
  ID.
 
  - Move this file to your working directory and rename it client_secret.json.


Also see our other package, [inboxer](https://gitlab.com/sigma-firma/inboxer), 
which makes performing basic actions on your inbox much more straight-forward. 


Package gsheet is a Go package for checking your gmail inbox, it has the following features:

  - Send emails
  - Mark emails (read/unread/important/etc)
  - Get labels used in inbox
  - Get emails by query (eg "in:sent after:2017/01/01 before:2017/01/30")
  - Get email metadata
  - Get email main body ("text/plain", "text/html")
  - Get the number of unread messages
  - Convert email dates to human readable format

#  USE


```go
package main

import (
        "context"
        "fmt"

        "github.com/sigma-firma/gmailAPI"
        "github.com/sigma-firma/inboxer"
        gmail "google.golang.org/api/gmail/v1"
)

func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.MailGoogleComScope)

        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                fmt.Println("========================================================")
                time, err := inboxer.ReceivedTime(msg.InternalDate)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println("Date: ", time)
                md := inboxer.GetPartialMetadata(msg)
                fmt.Println("From: ", md.From)
                fmt.Println("Sender: ", md.Sender)
                fmt.Println("Subject: ", md.Subject)
                fmt.Println("Delivered To: ", md.DeliveredTo)
                fmt.Println("To: ", md.To)
                fmt.Println("CC: ", md.CC)
                fmt.Println("Mailing List: ", md.MailingList)
                fmt.Println("Thread-Topic: ", md.ThreadTopic)
                fmt.Println("Snippet: ", msg.Snippet)
                body, err := inboxer.GetBody(msg, "text/plain")
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(body)
        }
}

```
## SENDING MAIL
```go
package main

import (
	"context"
	"log"

	"github.com/sigma-firma/gmailAPI"
	"github.com/sigma-firma/inboxer"
	gmail "google.golang.org/api/gmail/v1"
)

func main() {
    // Connect to the gmail API service.
	ctx := context.Background()
	srv := gmailAPI.ConnectToService(ctx, gmail.MailGoogleComScope)

    // Create a message
	var msg *inboxer.Msg = &inboxer.Msg{
		From:    "me",  // the authenticated user
		To:      "leadership@firma.com",
		Subject: "testing",
		Body:    "testing gmail api. lmk if you get this scott",
	}

    // send the email with the message
	err := msg.Send(srv)
	if err != nil {
		log.Println(err)
	}
}
```
## QUERIES

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)
        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                // do stuff
        }
}

```
## MARKING EMAILS

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        req := &gmail.ModifyMessageRequest{
                RemoveLabelIds: []string{"UNREAD"},
                AddLabelIds: []string{"OLD"}
        }

        // Range over the messages
        for _, msg := range msgs {
                msg, err := inboxer.MarkAs(srv, msg, req)
        }
}

```
## MARK ALL "UNREAD" EMAILS AS "READ"

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        inboxer.MarkAllAsRead(srv)
}
```
## GETTING LABELS

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        labels, err := inboxer.GetLabels(srv)
        if err != nil {
                fmt.Println(err)
        }

        for _, label := range labels {
                fmt.Println(label)
        }
}

```
## METADATA

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.MailGoogleComScope)

        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                fmt.Println("========================================================")
                md := inboxer.GetPartialMetadata(msg)
                fmt.Println("From: ", md.From)
                fmt.Println("Sender: ", md.Sender)
                fmt.Println("Subject: ", md.Subject)
                fmt.Println("Delivered To: ", md.DeliveredTo)
                fmt.Println("To: ", md.To)
                fmt.Println("CC: ", md.CC)
                fmt.Println("Mailing List: ", md.MailingList)
                fmt.Println("Thread-Topic: ", md.ThreadTopic)
        }
}

```
## GETTING THE EMAIL BODY

```go
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)
        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                body, err := inboxer.GetBody(msg, "text/plain")
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(body)
        }
}

```
## GETTING THE NUMBER OF UNREAD MESSAGES

```go
// NOTE: to actually view the email text use inboxer.Query and query for unread
// emails.
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        // num will be -1 on err
        num, err :=	inboxer.CheckForUnread(srv)
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("You have %s unread emails.", num)
}


```
## CONVERTING DATES

```go
// Convert UNIX time stamps to human readable format
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                // Convert the date
                time, err := inboxer.ReceivedTime(msg.InternalDate)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println("Date: ", time)
        }
}

```

## SNIPPET

```go
// Snippets are not really part of the package but I'm including them in the doc
// because they'll likely be useful to anyone working with this package.
func main() {
        // Connect to the gmail API service.
        ctx := context.Background()
        srv := gmailAPI.ConnectToService(ctx, gmail.GmailComposeScope)

        msgs, err := inboxer.Query(srv, "category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                // this one is part of the api
                fmt.Println(msg.Snippet)
        }
}
`

## MORE ON CREDENTIALS:

For gsheet to work you must have a gmail account and a file containing your 
authorization info in the root directory of your project. To obtain credentials 
please see step one of this guide: https://developers.google.com/gmail/api/quickstart/go

 Turning on the gmail API

 - Use this wizard (https://console.developers.google.com/start/api?id=gmail) to create or select a project in the Google Developers Console and automatically turn on the API. Click Continue, then Go to credentials.

 - On the Add credentials to your project page, click the Cancel button.

 - At the top of the page, select the OAuth consent screen tab. Select an Email address, enter a Product name if not already set, and click the Save button.

 - Select the Credentials tab, click the Create credentials button and select OAuth client ID.

 - Select the application type Other, enter the name "Gmail API Quickstart", and click the Create button.

 - Click OK to dismiss the resulting dialog.

 - Click the file_download (Download JSON) button to the right of the client ID.

 - Move this file to your working directory and rename it client_secret.json.


