package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrototypeManager(t *testing.T) {
	mng := NewPrototypeManager()

	t1 := &Type1{text: "type1"} // 指针类型
	t2 := Type2{val: 1}         // 值类型
	
	mng.Set("t1", t1)
	mng.Set("t2", t2)

	t1Clone := mng.Get("t1")
	t2Clone := mng.Get("t2")

	// 通过原型管理器获取的克隆对象, 与原对象的值相等
	assert.Equal(t, t1.text, t1Clone.(*Type1).text)
	assert.Equal(t, t2.val, t2Clone.(Type2).val)

	// 虽然是不同的对象, 但是值相等, golang就会认为2 struct是相等的
	assert.Equal(t, t1, t1Clone)
	assert.Equal(t, t2, t2Clone)

	// 通过原型管理器获取的克隆对象, 与原对象不是同一个对象
	assert.NotEqual(t, &t1, &t1Clone)
	assert.NotEqual(t, &t2, &t2Clone)
}
