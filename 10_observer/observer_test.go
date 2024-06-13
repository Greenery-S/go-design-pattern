package observer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSubject(t *testing.T) {
	var (
		subject         *Subject
		newspaperReader *NewPaperReader
		mobileReader    *MobileReader
	)

	// 实例化读者
	newspaperReader = &NewPaperReader{name: "Old William"}
	mobileReader = &MobileReader{name: "Young Sherry"}

	// 实例化主题
	subject = NewSubject()
	subject.Attach(newspaperReader)
	subject.Attach(mobileReader)

	// 发布内容
	subject.content = "Welcome to subscribe our news!"
	subject.notify()

	assert.Equal(t, "Welcome to subscribe our news!", newspaperReader.latestContent)
	assert.Equal(t, "Welcome to subscribe our news!", mobileReader.latestContent)

	// 发布内容
	subject.content = "Today is a good day!"
	subject.notify()

	assert.Equal(t, "Today is a good day!", newspaperReader.latestContent)
	assert.Equal(t, "Today is a good day!", mobileReader.latestContent)
}
