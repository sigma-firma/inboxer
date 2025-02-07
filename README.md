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

# `gsheet` Features 

  - Access
    - Connect/Authenticate with sheets and/or gmail api using your credentials
      and recieve a refresh token
    - Auto refresh the token every 12 hours by default (can be adjusted)

  - Sheets
    - Append row to spread sheet in google Sheets
    - Read data from spread sheet

  - Gmail
    - Send emails
    - Mark emails (read/unread/important/etc)
    - Get labels used in inbox
    - Get emails by query (eg "in:sent after:2017/01/01 before:2017/01/30")
    - Get email metadata
    - Get email main body ("text/plain", "text/html")
    - Get the number of unread messages
    - Convert email dates to human readable format


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

```go
func main() {
	var access *gsheet.Access = gsheet.NewAccess(
        //location of credentials.json 
        // NOTE: Get this from Google
		os.Getenv("HOME")+"/credentials/credentials.json",
        // Location of token.json, or where/what it should be saved /as. 
        // NOTE: This will automatically download if you don't have it
		os.Getenv("HOME")+"/credentials/token.json",
        // Scopes
		[]string{
			gmail.GmailComposeScope,
			sheets.SpreadsheetsScope,
		})

    // connect to gmail
    gm := access.Gmail()

    // connect to sheets
    sh := access.Sheets()
}

```


# Example Usage

## Getting *Access

### Setting up credentials/tokens and refreshing them

```go
package main

import (
        "fmt"
        "log"

        "github.com/sigma-firma/gsheet"
)

// Instantiate a new *Access struct with essential values
var access *gsheet.Access = gsheet.NewAccess(
	os.Getenv("HOME")+"/credentials/credentials.json",
	os.Getenv("HOME")+"/credentials/quickstart.json",
	[]string{gmail.GmailComposeScope, sheets.SpreadsheetsScope},
)

// Connect to the Gmail API
var gm *gsheet.Gmailer = access.Gmail()

// Connect to the Sheets API
var sh *gsheet.Sheeter = access.Gmail()
```

## Sheets

### Reading values from a spreadsheet:

```go
func main() {                                                                          
        package main

        import (
                "fmt"
                "log"

                "github.com/sigma-firma/gsheet"
        )

        // Connect to the API                                                          
        sh := access.Sheets()


        // Prints the names and majors of students in a sample spreadsheet:
        // https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
       	var req *gsheet.Spread = &gsheet.Spread{
                ID:               "1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms",
                ReadRange:        "Class Data!A2:E",
       }
 
        resp, err := sh.Read(req)
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

```go
package main

import (
        "fmt"
        "log"

        "github.com/sigma-firma/gsheet"
)

func main() {                                                                          

	var row []interface{} = []interface{}{"hello A1","world B1"}))

	var req *gsheet.Spread = &gsheet.Spread{
		ID:               "1cZVwQaY8LqsIUwzbCm_yG8tcR5RDog9jD1sHJtF9mSA",
		WriteRange:       "A1",
		Vals:             row,
		ValueInputOption: "RAW",
	}

	sh := access.Sheets()
    _, err := sh.AppendRow(req)
	if err != nil {
		return err
	}
}

```

## GMAIL

### Check for new unread messages

```go
	func main() {
			// Connect to the Gmail API service. Here, we use a context and provide a
			// scope. The scope is used by the Gamil API to determine your privilege
			// level. gmailAPI.ConnectToService is a variadic function and thus can be
			// passed any number of scopes. For more information on scopes see:
			// https://developers.google.com/gmail/api/auth/scopes
            gm := access.Gmail()

			// Check if you have any unread messages
			count, err := gm.CheckForUnread()
			if err != nil {
					fmt.Println(err)
			}
            if count >0 {
                    fmt.Println("You've got mail.")
            }

	}

```


### Query

```go
package main

import (
        "context"
        "fmt"

        "github.com/sigma-firma/gmailAPI"
        "github.com/sigma-firma/gm"
        gmail "google.golang.org/api/gmail/v1"
)

func main() {
        gm := access.Gmail()
        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                fmt.Println("========================================================")
                time, err := gm.ReceivedTime(msg.InternalDate)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println("Date: ", time)
                md := gm.GetPartialMetadata(msg)
                fmt.Println("From: ", md.From)
                fmt.Println("Sender: ", md.Sender)
                fmt.Println("Subject: ", md.Subject)
                fmt.Println("Delivered To: ", md.DeliveredTo)
                fmt.Println("To: ", md.To)
                fmt.Println("CC: ", md.CC)
                fmt.Println("Mailing List: ", md.MailingList)
                fmt.Println("Thread-Topic: ", md.ThreadTopic)
                fmt.Println("Snippet: ", msg.Snippet)
                body, err := gm.GetBody(msg, "text/plain")
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(body)
        }
}

```

### Sending mail

```go
package main

import (
	"context"
	"log"

	"github.com/sigma-firma/gmailAPI"
	"github.com/sigma-firma/gm"
	gmail "google.golang.org/api/gmail/v1"
)

func main() {
    // Create a message
	var msg *gsheet.Msg = &gsheet.Msg{
		From:    "me",  // the authenticated user
		To:      "leadership@firma.com",
		Subject: "testing",
		Body:    "testing gmail api. lmk if you get this scott",
	}

    // send the email with the message
	err := msg.Send()
	if err != nil {
		log.Println(err)
	}
}
```

### Marking emails

```go
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        req := &gmail.ModifyMessageRequest{
                RemoveLabelIds: []string{"UNREAD"},
                AddLabelIds: []string{"OLD"}
        }

        // Range over the messages
        for _, msg := range msgs {
                msg, err := gm.MarkAs(msg, req)
        }
}

```
### Mark all "unread" emails as "read"

```go
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()
        gm.MarkAllAsRead()
}
```
### Getting labels

```go
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        labels, err := gm.GetLabels()
        if err != nil {
                fmt.Println(err)
        }

        for _, label := range labels {
                fmt.Println(label)
        }
}

```
### Metadata

```go
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                fmt.Println("========================================================")
                md := gm.GetPartialMetadata(msg)
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
### Getting the email body

```go
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()
        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                body, err := gm.GetBody(msg, "text/plain")
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println(body)
        }
}

```
### Getting the number of unread messages

```go
// NOTE: to actually view the email text use gm.Query and query for unread
// emails.
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        // num will be -1 on err
        num, err :=	gm.CheckForUnread()
        if err != nil {
                fmt.Println(err)
        }
        fmt.Printf("You have %s unread emails.", num)
}


```
### Converting dates

```go
// Convert UNIX time stamps to human readable format
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                // Convert the date
                time, err := gm.ReceivedTime(msg.InternalDate)
                if err != nil {
                        fmt.Println(err)
                }
                fmt.Println("Date: ", time)
        }
}

```

### Snippet

```go
// Snippets are not really part of the package but I'm including them in the doc
// because they'll likely be useful to anyone working with this package.
func main() {
        // Connect to the gmail API service.
        gm := access.Gmail()

        msgs, err := gm.Query("category:forums after:2017/01/01 before:2017/01/30")
        if err != nil {
                fmt.Println(err)
        }

        // Range over the messages
        for _, msg := range msgs {
                // this one is part of the api
                fmt.Println(msg.Snippet)
        }
}
```

# More on credentials:

For gsheet to work you must have a gmail account and a file containing your 
authorization info in the directory you will specify when setting up gsheet. 
To obtain credentials please see step one of this guide: 

https://developers.google.com/gmail/api/quickstart/go

Turning on the gmail API (should be similar for the Sheets API)

 - Use this wizard (https://console.developers.google.com/start/api?id=gmail) to create or select a project in the Google Developers Console and automatically turn on the API. Click Continue, then Go to credentials.

 - On the Add credentials to your project page, click the Cancel button.

 - At the top of the page, select the OAuth consent screen tab. Select an Email address, enter a Product name if not already set, and click the Save button.

 - Select the Credentials tab, click the Create credentials button and select OAuth client ID.

 - Select the application type Other, enter the name "Gmail API Quickstart", and click the Create button.

 - Click OK to dismiss the resulting dialog.

 - Click the file_download (Download JSON) button to the right of the client ID.

 - Move this file to your working directory and rename it client_secret.json.
