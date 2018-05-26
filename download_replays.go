package main

import (
    "net/http"
    "fmt"
	"path"
	"os"
	"io"
	"time"
)

func downloadFile(url string, destinationDir string) (err error) {
	filename := path.Base(url)
	destPath := path.Join(destinationDir, filename)

	output, err := os.Create(destPath)
	if err != nil {
		return err
	}

	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("Received HTTP error: %s", response.Status)
	}

	_, err = io.Copy(output, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	linkTemplate := "https://www.wcreplays.com/replay.php?m=DownloadReplay&rid=%v"
	waitDuration := 1000 * time.Millisecond
	replaysDir := "/tmp/w3replays"
	
	for i := 2; i < 10; i++ {
		url := fmt.Sprintf(linkTemplate, i)
		err := downloadFile(url, replaysDir)
		
		if err != nil {
			fmt.Printf("Error: %s", err)
		}
		
		time.Sleep(waitDuration)
	}
}
