package logic

//Downloader 下载接口
type Downloader interface {
	Download(uri string) (string, error)
	GetLocalPath() string
}
