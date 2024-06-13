package flyweight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	time.Sleep(1 * time.Second)
	viewer2 := NewImageViewer("image2.png")
	time.Sleep(1 * time.Second)
	viewer3 := NewImageViewer("image1.png")

	fmt.Println(viewer1.createAt.Unix())
	fmt.Println(viewer2.createAt.Unix())
	fmt.Println(viewer3.createAt.Unix())

	assert.Equal(t, viewer1, viewer3)

}
