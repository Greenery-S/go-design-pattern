package flyweight

import (
	"fmt"
	"time"
)

//// * 享元模式

// ImageFlyweightFactory 图片享元工厂 用于保存图片数据, 如果图片已经存在,则直接获取,否则创建一个新的图片
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

// GetImageFlyweightFactory 获取一个 图片享元工厂 实例,全局唯一
// 如果不存在,则创建一个新的实例
var imageFactory *ImageFlyweightFactory

// GetImageFlyweightFactory 获取一个ImageFlyweightFactory实例,全局唯一
func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

// Get 用于获取图片数据, 如果图片已经存在,则直接获取,否则创建一个新的图片
func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}

//// 定义单个图片的数据(单个图片的数据是不变的,所以可以共享)

// ImageFlyweight 用于保存图片数据
type ImageFlyweight struct {
	data     string
	createAt time.Time // 保存创建时间,用于判断是否需要重新加载
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data:     data,
		createAt: time.Now(),
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

// ImageViewer 用于显示图片, 通过ImageFlyweightFactory获取图片数据
// 如果查看的图片已经存在,则直接获取(享元模式),否则创建一个新的图片
type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
