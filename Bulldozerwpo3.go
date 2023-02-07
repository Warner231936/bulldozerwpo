// go:build windows
//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v2"
)

type Point struct {
	X int `yaml:"x"`
	Y int `yaml:"y"`
}

type Delay struct {
	Base   int `yaml:"base"`
	Target int `yaml:"target"`
	Loop   int `yaml:"loop"`
}
type Tune struct {
	Search Point `yaml:"search"`
	Locate Point `yaml:"locate"`
	InputX Point `yaml:"input_x"`
	InputY Point `yaml:"input_y"`
	Look   Point `yaml:"look"`
	Center Point `yaml:"center"`
	Attack Point `yaml:"attack"`
	Preset Point `yaml:"preset"`
	Finish Point `yaml:"finish"`
}

type Config struct {
	TargetFile string `yaml:"target_file"`
	AltsCount  int    `yaml:"alts_count"`
	Delay      Delay  `yaml:"delay"`
	Tune       Tune   `yaml:"tune"`
}

var config Config
var targets []Point

func init() {
	configFile, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

	targetsFile, err := os.Open(config.TargetFile)
	if err != nil {
		panic(err)
	}
	defer targetsFile.Close()

	decodera := yaml.NewDecoder(targetsFile)
	err = decodera.Decode(&targets)
	if err != nil {
		panic(err)
	}
}

func main() {
	go func() {
		for {
			key := robotgo.AddEvent("q")
			if key {
				fmt.Println("Exiting.....")

				os.Exit(0)
			}
		}
	}()

	pre := "Bulldozer | "
	fmt.Println(pre + "War Planet Online")
	time.Sleep(2 * time.Second)
	fmt.Printf(pre+"Number of targets: %d\n", len(targets))
	time.Sleep(1 * time.Second)
	fmt.Println(pre + "Press Enter to start")
	fmt.Scanln()
	fmt.Println(pre + "Press Q at any time to exit")

	prea := "Bulldozer:Boot | "
	fmt.Print(prea + "Booting up")
	for i := 0; i < 20; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

	anyUnderLimit := false
	if config.Delay.Base < 200 {
		fmt.Printf(prea+"%s delay is very short (%dms), reccomended delay is at least %dms!\n", "Base", config.Delay.Base, 300)
		anyUnderLimit = true
	}
	if config.Delay.Target < 200 {
		fmt.Printf(prea+"%s delay is very short (%dms), reccomended delay is at least %dms!\n", "Target", config.Delay.Target, 300)
		anyUnderLimit = true
	}
	if config.Delay.Loop < 1800 {
		fmt.Printf(prea+"%s delay is very short (%ds), reccomended delay is at least %ds!\n", "Loop", config.Delay.Loop, 1800)
		anyUnderLimit = true
	}
	if anyUnderLimit {
		time.Sleep(5 * time.Second)
	}

	fmt.Println(prea + "Boot process complete")
	fmt.Println(pre + "Starting automation routine in 10 seconds")
	time.Sleep(2 * time.Second)
	fpid, err := robotgo.FindIds("WarPlanetOnline.exe")
	if err != nil {
		panic(err)
	}

	if len(fpid) < 1 {
		fmt.Println("No running instance detected! Launch the game before restarting this software.")
		fmt.Scanln()
		panic(1)
	}
	if len(fpid) > 1 {
		fmt.Println("Multiple running instances detected!")
		fmt.Scanln()
		panic(1)
	}
	robotgo.MaxWindow(fpid[0])
	robotgo.SetActive(robotgo.GetHandPid(fpid[0]))

	time.Sleep(8 * time.Second)
	fmt.Println(pre + "Automation routine started")

	for loopCount := 1; true; loopCount++ {
		pre := "Bulldozer:Loop " + strconv.Itoa(loopCount) + " | "
		fmt.Println(pre + "Starting run")

		for i, target := range targets {
			fmt.Printf(pre+"Target %d / %d\n", i+1, len(targets))
			click(config.Tune.Search)
			click(config.Tune.Locate)

			click(config.Tune.InputX)
			typea(target.X)
			enter()
			click(config.Tune.InputY)
			typea(target.Y)

			enter()
			enter()
			time.Sleep(1 * time.Second)

			doubleClick(config.Tune.Center)

			click(config.Tune.Attack)
			click(config.Tune.Preset)
			click(config.Tune.Finish)

			time.Sleep(time.Duration(config.Delay.Target) * time.Millisecond)
		}

		fmt.Println(pre + "Run complete")
		fmt.Printf(pre+"Next run starting in %d seconds", config.Delay.Loop)
		Sleeptime := time.Duration(config.Delay.Loop) * time.Second
		time.Sleep(Sleeptime)
	}
}

func click(pos Point) {
	robotgo.MoveMouse(pos.X, pos.Y)
	time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
	robotgo.MouseClick()
	time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
}
func doubleClick(pos Point) {
	robotgo.MoveMouse(pos.X, pos.Y)
	time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
	robotgo.MouseClick()
	time.Sleep(100 * time.Millisecond)
	robotgo.MouseClick()
	time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
}

func typea(num int) {
	text := strconv.Itoa(num)
	typeText(text)
}
func typeText(text string) {
	for _, char := range text {
		str := string(char)
		robotgo.KeyTap(str)
		time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
	}
}

func enter() {
	robotgo.KeyTap("enter")
	time.Sleep(time.Duration(config.Delay.Base) * time.Millisecond)
}
