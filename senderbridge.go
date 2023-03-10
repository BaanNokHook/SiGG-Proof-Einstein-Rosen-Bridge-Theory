package main

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import (
	"fmt"
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type SenderBridge struct {
	core.QObject

	_ func() `constructor:"init"`
	_ func() `signal:"clickAddFile,auto"`
	_ string `property:"code"`

	_ *SendFileTableModel `property:"TableModel"`
}

func (l *SenderBridge) init() {
	senderBridge = l
	l.SetTableModel(senderTableModel)
}

func (b *SenderBridge) clickAddFile() {
	jobtotal := new(int64)
	jobdone := new(int64)
	feedbackstr := new(string)
	mycode := new(string) // The real code
	code := new(string)   // Transport the code from thread to copy into mycode
	var tableIndex int = 0

	filenames := widgets.QFileDialog_GetOpenFileNames(nil, "some caption", "", "", "", 0)
	if len(filenames) < 1 {
		return
	}
	filename := filenames[0]
	fmt.Printf("Sending: %s\n", filename)

	stat, err := os.Stat(filename)
	if err != nil {
		showError(fmt.Sprintf("Failed to read %s: %s", filename, err))
	}
	*jobtotal = stat.Size()
	*jobdone = 0

	tableIndex = senderTableModel.addNative(
		filename,
		*mycode,
		strconv.FormatInt(*jobtotal, 10),
		"0",
		"Added")

	// Handling results of changes in
	// jobdone, feedbackstr, code
	t := core.NewQTimer(nil)
	t.ConnectEvent(func(e *core.QEvent) bool {
		// Handle errors and finish
		if len(*feedbackstr) > 0 {
			t.DisconnectEvent()

			if *feedbackstr == "Done" {
				senderTableModel.editIdx(
					tableIndex,
					filename,
					*mycode,
					strconv.FormatInt(*jobtotal, 10),
					strconv.FormatInt(*jobdone, 10),
					"Done")
			} else {
				senderTableModel.editIdx(
					tableIndex,
					filename,
					*mycode,
					strconv.FormatInt(*jobtotal, 10),
					strconv.FormatInt(*jobdone, 10),
					"Error")

				showError(*feedbackstr)
			}

			return true
		}

		// Handle start and upload updates
		if *jobdone > 0 && *jobdone < *jobtotal {
			senderTableModel.editIdx(
				tableIndex,
				filename,
				*mycode,
				strconv.FormatInt(*jobtotal, 10),
				strconv.FormatInt(*jobdone, 10),
				"Uploading") //
		}

		// Handle successful registering/adding of file
		// Once!
		if len(*code) > 0 {
			*mycode = *code
			senderTableModel.editIdx(
				tableIndex,
				filename,
				*mycode,
				strconv.FormatInt(*jobtotal, 10),
				strconv.FormatInt(*jobdone, 10),
				"Added")
			*code = ""
		}
		return true
	})
	t.Start(200) // Every x ms

	// Thread doing the actual file sending
	// It communicates with parent via variables
	// jobdone, feedbackstr, code
	go func() {
		/*
			defer func() {
				if err := recover(); err != nil {
					*errstr = fmt.Sprintf("%v", err)
				}
			}()
		*/

		// Start downloading it in the background
		if stat.IsDir() {
			code2, status, err := sendDir(filename, jobdone, feedbackstr)
			if err != nil {
				*feedbackstr = fmt.Sprintf("Error sending message: %s", err)
				return
			}

			// Cant update TableModel here in this go func()
			// Return it to the QTimer
			// before waiting for upload result
			*code = code2

			// Wait till its finished
			s := <-status
			if s.OK {
				fmt.Println("file sent")
				*feedbackstr = "Done"
			} else {
				*feedbackstr = fmt.Sprintf("Send error: %s", s.Error)
			}
		} else {
			code2, status, err := sendFile(filename, jobdone, feedbackstr)
			if err != nil {
				*feedbackstr = fmt.Sprintf("Error sending message: %s", err)
				return
			}

			// Cant update TableModel here in this go func()
			// Return it to the QTimer
			// before waiting for upload result
			*code = code2

			// Wait till its finished
			s := <-status
			if s.OK {
				fmt.Println("file sent")
				*feedbackstr = "Done"
			} else {
				*feedbackstr = fmt.Sprintf("Send error: %s", s.Error)
			}
		}

		//*feedbackstr = "ok"
	}()

}
