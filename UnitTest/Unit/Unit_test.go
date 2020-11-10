package Unit

import (
	"fmt"
	"testing"
)

func TestMonster_Store(t *testing.T) {
	monster := &Monster{"辛达苟萨", 5000, "绝对零度"}
	res :=monster.Store()
	fmt.Println(string(res))
}