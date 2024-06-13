package templatemethod

import "fmt"

//// * 模版方法模式
// 模版方法模式是一种行为设计模式，它在超类中定义了一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。
// 如果超类中的方法是抽象的，那么子类必须实现它们。

// Downloader 是一个下载器的接口
type Downloader interface {
	Download(uri string)
}

// template 是一个下载器的模版, 实现了 Downloader 接口, 是所有下载器的基类
// 它包含一个 implement 接口, 该接口包含了 download 和 save 两个方法:
//   - download 方法用于下载资源, 必须由子类实现
//   - save 方法用于保存资源, 可以被子类覆盖
//
// # Download 方法定义了下载资源的流程, 该流程中调用了 download 和 save 两个方法
//
// uri 是下载资源的地址, 独立于 implement 接口, 实现了数据和行为的分离
type template struct {
	implement        // 下载行为模版
	uri       string // 下载资源的地址
}

// implement 是一个下载器的实现接口
type implement interface {
	download()
	save()
}

// save 实现了部分模版方法, 是一个默认的保存方法
func (t *template) save() {
	fmt.Print("default save\n")
}

// Download 实现下载器接口, 定义了下载资源的流程
func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")

}

func newTemplate(impl implement) *template {
	return &template{
		implement: impl,
	}
}

//// * http 下载器

// HTTPDownloader 是 implement 接口的一个实现, 定义了 http 下载器的行为
type HTTPDownloader struct {
	*template
}

// NewHTTPDownloader 返回一个 http 下载器
// 实例化一个http版的 implement 实现, 并将其传递给Downloader模版
func NewHTTPDownloader() Downloader {
	downloader := &HTTPDownloader{}
	tmpl := newTemplate(downloader)
	downloader.template = tmpl
	return downloader
}

// download 实现了下载行为
func (d *HTTPDownloader) download() {
	fmt.Printf("download %s via http\n", d.uri)
}

// save 重写了保存行为
func (*HTTPDownloader) save() {
	fmt.Printf("http save\n")
}

//// * ftp 下载器

type FTPDownloader struct {
	*template
}

func NewFTPDownloader() Downloader {
	downloader := &FTPDownloader{}
	template := newTemplate(downloader)
	downloader.template = template
	return downloader
}

func (d *FTPDownloader) download() {
	fmt.Printf("download %s via ftp\n", d.uri)
}
