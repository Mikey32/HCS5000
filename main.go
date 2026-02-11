package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 1. Load the "Permanent" settings from JSON
	cfg, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("❌ Error: Missing config.json: %v", err)
	}

	// 2. Get the "Unique" setting from the user
	fmt.Println("Welcome to the Hotspot Configuration Script.... 5000!")
	fmt.Print("🔑 Enter admin password for THIS device: ")
	var stickerPass string
	fmt.Scanln(&stickerPass)

	// Inject the manual password into our config struct
	cfg.CurrentAdminPassword = stickerPass

	fmt.Println("🚀 Initializing...")

	// 3. Setup Headless Browser
	opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", true), chromedp.Flag("disable-session-crashed-bubble", true), chromedp.Flag("incognito", true), chromedp.Flag("disable-infobars", true))
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// 4. Execute with 60s timeout
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	fmt.Println("⚙️  Programming Orbic Hotspot... ")

	err = chromedp.Run(ctx, ConfigureDevice(cfg))

	if err != nil {
		fmt.Printf("\n❌ Failed: %v\n", err)
	} else {
		fmt.Println("\n✅ Success! Device configured.")
	}
}

func LoadConfig(filename string) (HotspotConfig, error) {
	var config HotspotConfig
	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decode := json.NewDecoder(file)
	err = decode.Decode(&config)
	return config, err
}
