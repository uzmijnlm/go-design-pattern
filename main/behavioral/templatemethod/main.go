package main

func main() {
	var downloader Downloader
	downloader = &HTTPDownloader{&BaseDownloader{uri: "xxx"}}

	util := &DownloadUtil{}
	util.downloader = &downloader
	util.Download()

	downloader = &FTPDownloader{&BaseDownloader{uri: "yyy"}}
	util.downloader = &downloader
	util.Download()

}
