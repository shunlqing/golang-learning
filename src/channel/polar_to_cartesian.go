//将极坐标转换为笛卡尔坐标

/**
要求:
1. 有交互控制台程序
2. 转换处理在协程中进行
3. 使用channel1接收极坐标,使用channel2接收笛卡尔坐标
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"runtime"
	"strconv"
	"strings"
	"math"
)

type polar struct {
	radius	float64
	theta	float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "Polar: radius=%0.2f angle=%0.2f degress -- Cartesian: x=%0.2f y=%0.2f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

// 创建处理器
// 处理器的创建接收指定的channel用于输入,返回用于输出结果的channel
func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			theta := polarCoord.theta * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(theta)
			y := polarCoord.radius * math.Sin(theta)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

//交互程序
func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)

	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n') //返回字符串line,包含'\n'
		if err != nil {
			break
		}
		line = line[:len(line)-1] // 调整字符串的大小
		if numbers := strings.Fields(line); len(numbers) == 2 { //以空格符分割获取字符串切片
			polars, err := floatsForStrings(numbers)
			if err != nil {
				fmt.Println(os.Stderr, "invalid number")
				continue
			}
			questions <- polar{polars[0], polars[1]} //将参数输入到通道中
			coord := <-answers  //从通道中获取结果
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			fmt.Fprintln(os.Stderr, "invalid input")
		}
	}
	fmt.Println()
}

// 将字符串切片转换为浮点数切片
func floatsForStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}

	return floats, nil
}