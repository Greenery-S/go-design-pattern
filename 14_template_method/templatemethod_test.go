package templatemethod

import "testing"

func TestTemplate_Download(t *testing.T) {
	var downloader Downloader

	downloader = NewHTTPDownloader()
	downloader.Download("example.com/abc.zip")

	downloader = NewFTPDownloader()
	downloader.Download("example.com/abc.zip")
}
