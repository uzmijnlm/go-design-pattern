package main

import "fmt"

type Downloader interface {
	Before()
	DoDownload()
	After()
}

type BaseDownloader struct {
	uri string
}

func (downloader *BaseDownloader) Before() {
	fmt.Println("before...")
}

func (downloader *BaseDownloader) After() {
	fmt.Println("after...")
}

type HTTPDownloader struct {
	*BaseDownloader
}

func (downloader *HTTPDownloader) DoDownload() {
	fmt.Println("HTTP Download " + downloader.uri)
}

type FTPDownloader struct {
	*BaseDownloader
}

func (downloader *FTPDownloader) DoDownload() {
	fmt.Println("FTP Download " + downloader.uri)
}

type DownloadUtil struct {
	downloader *Downloader
}

func (util *DownloadUtil) Download() {
	(*util.downloader).Before()
	(*util.downloader).DoDownload()
	(*util.downloader).After()
}
