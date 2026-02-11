package main

import (
	"time"

	"github.com/chromedp/chromedp"
)

type HotspotConfig struct {
	CurrentAdminPassword string `json:"current_admin_password"`
	NewAdminPassword     string `json:"new_admin_password"`
	Wifi24Pw             string `json:"wifi_24_pw"`
	Wifi5gPw             string `json:"wifi_5g_pw"`
	SSID                 string `json:"ssid"`
	DNSMode              string `json:"dns_mode"`
	DNSPri               string `json:"dns_pri"`
	DNSSec               string `json:"dns_sec"`
}

func ConfigureDevice(cfg HotspotConfig) chromedp.Tasks {
	return chromedp.Tasks{
		// --- 1. LOGIN ---
		chromedp.Navigate(`http://192.168.1.1/common/login.html`),
		chromedp.WaitVisible(`body`), // Wait for the page itself to exist
		chromedp.Click(`body`),       // Click anywhere to pull focus from the address bar
		chromedp.Sleep(2 * time.Second),
		chromedp.Focus(`#username`),
		chromedp.Click(`#username`),
		chromedp.SendKeys(`#username`, `admin`),
		chromedp.Focus(`#password`),
		chromedp.SendKeys(`#password`, cfg.CurrentAdminPassword),
		chromedp.SendKeys(`#password`, "\n"),
		//chromedp.Click(`#loginBtn`),

		// --- 2. WIFI & SSID ---
		chromedp.Navigate(`http://192.168.1.1/html/wlansettings.html`),
		chromedp.WaitVisible(`#ssidname24`),
		chromedp.SetValue(`#ssidname24`, cfg.SSID),
		chromedp.SetValue(`#ssidname5`, cfg.SSID),
		// ... set wifi passwords ...
		chromedp.SetValue(`securitykey24`, cfg.Wifi24Pw),
		chromedp.SetValue(`securitykey5`, cfg.Wifi5gPw),
		chromedp.Click(`#apply`),
		chromedp.Sleep(2 * time.Second), // Brief pause for hardware to commit

		// --- 3. DNS ---
		chromedp.Navigate(`http://192.168.1.1/html/dns.html`),
		chromedp.WaitVisible(`#dnsswitch`),
		// 1. Wait for the DNS dropdown to be ready
		chromedp.WaitVisible(`#dnsswitch`),

		// 2. Force a click on it first (helps wake up the listener)
		chromedp.Click(`#dnsswitch`),

		// 3. Select the "Manual" option (value "1")
		chromedp.SetValue(`#dnsswitch`, "1"),

		// This tells the Orbic's JavaScript "Hey, the user changed this!"
		chromedp.Evaluate(`document.querySelector("#dnsswitch").dispatchEvent(new Event('change', { bubbles: true }));`, nil),

		chromedp.Click(`body`),
		chromedp.Sleep(500 * time.Millisecond),

		// 5. Short pause to let any hidden DNS input boxes appear
		chromedp.Sleep(1 * time.Second),

		chromedp.SetValue(`#primarydns`, cfg.DNSPri),
		chromedp.SetValue(`#secondarydns`, cfg.DNSSec),
		chromedp.Click(`#apply-dns`),

		// --- 4. UPDATE ADMIN PASSWORD ---
		// Doing this last ensures you have authorization for all previous steps
		chromedp.Navigate(`http://192.168.1.1/html/systemadmin.html`),
		chromedp.WaitVisible(`#curpsd`),
		chromedp.SetValue(`#curpsd`, cfg.CurrentAdminPassword),
		chromedp.SetValue(`#newpsd`, cfg.NewAdminPassword),
		chromedp.SetValue(`#confirmpsd`, cfg.NewAdminPassword),
		chromedp.Sleep(1 * time.Second),
		chromedp.Click(`#apply-psw`),
		chromedp.Sleep(1 * time.Second),
	}
}
