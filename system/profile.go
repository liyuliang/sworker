package system

import (
	"github.com/mackerelio/go-osstat/memory"
	"github.com/mackerelio/go-osstat/cpu"
	"fmt"
	"os"
	"time"
	"runtime"
	"os/exec"
	"bytes"
	"strings"
)

//获取系统发行版本
func GetLinuxVersion() (v string) {

	switch runtime.GOOS {
	case "darwin":
		v, _ = shell("sw_vers")
	case "linux":
		v, _ = shell("cat /etc/issue")
		if strings.Contains(strings.ToLower(v), "ubuntu") == false {
			v, _ = shell("cat /etc/redhat-release")

		}
	}

	v = strings.Split(v, "\n")[0]
	v = strings.Replace(v, "ProductName:", "", -1)
	v = strings.Replace(v, "\t", "", -1)
	return v
}

//获取系统名称
func GetHostName() string {
	name, err := os.Hostname()
	if err == nil {
		return name
	}
	return ""
}

//获取系统核心数量
func GetCoreNum() (num int) {
	num = runtime.GOMAXPROCS(0)
	return num
}

//获取负载
func GetLoadAverage() (s string) {
	switch runtime.GOOS {
	case "darwin":
		s, _ = shell(`w | head -n1`)
	case "linux":
		s, _ = shell(`uptime`)
	}

	s = strings.Replace(s, "\n", "", -1)
	return s
}

//获取磁盘占用量
func GetDiskUsage() (u string) {
	str, _ := shell("df -lh")
	s := strings.Split(str, "\n")
	real := ""
	for value, _ := range s {
		if strings.Index(s[value]+" ", " / ") != -1 {
			real = s[value]
		}
	}
	real = strings.Replace(real, "  ", " ", -1)
	realData := strings.Split(real, " ")

	data := make([]string, 0)

	for index := range realData {
		if strings.Replace(realData[index], " ", "", -1) != "" {
			data = append(data, realData[index])
		}
	}

	u = fmt.Sprintf("%s/%s, %s", data[2], data[1], data[4])
	return u
}

//执行系统命令
func shell(s string) (string, error) {

	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetCpuUsage() (u string) {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	total := float64(after.Total - before.Total)

	u = fmt.Sprintf("%f/%f", float64(after.User-before.User)/total*100, float64(after.System-before.System)/total*100)
	return u
}

func GetMemUsage() (u string) {

	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	used := bToMb(memory.Used)
	total := bToMb(memory.Total)
	if total > 1024 {
		u = fmt.Sprintf("%d GB/%d GB", bToG(memory.Used), bToG(memory.Total))
	} else {
		u = fmt.Sprintf("%d MB/%d MB", used, total)
	}
	return u
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func bToG(b uint64) uint64 {
	m := bToMb(b)
	if m > 1024 {

	}
	return bToMb(b) / 1024
}
