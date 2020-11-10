package Unit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `jason:"name"`
	Age   int    `jason:"age"`
	Skill string `jason:"skill"`
}

func (m *Monster) Init() Monster {
	m.Name = "玛里苟斯"
	m.Age = 3000
	m.Skill = "寒冰吐息"
	return *m
}

func (m *Monster) Store() []byte {
	data, err := json.Marshal(m.Init())
	if err != nil {
		fmt.Printf("序列化失败，err = %v\n", err)
	}
	return data
}

func (m *Monster) ReStore(data []byte){
	monster := json.Unmarshal(data, m)
	fmt.Println(monster)
}

func (m *Monster) WriteFile(data []byte) {
	FilePath := "/Users/ace/Desktop/Monster.json"
	File, err := os.OpenFile(FilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件失败，err =%v\n", err)
		return
	}
	defer File.Close()
	writer := bufio.NewWriter(File)

	_, err = writer.Write(data)
	if err != nil {
		fmt.Printf("写入文件失败，err =%v\n", err)
		return
	}
	_ = writer.Flush()
}

func (m *Monster) ReadFile(){
	FilePath := "/Users/ace/Desktop/Monster.json"

	data, err := ioutil.ReadFile(FilePath)
	if err != nil{
		fmt.Printf("读取文件失败，err =%v\n", err)
	}
	err = json.Unmarshal(data, m)
	if err != nil{
		fmt.Printf("反序列化失败，err =%v\n", err)
	}
	fmt.Println(m)
}
