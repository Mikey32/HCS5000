# HCS5000: Hotspot Configuration Script 5000
## This for the Orbic Hotspots

HCS5000 is an administrative utility for lifecycle management and batch provisioning of Orbic devices in a workplace environment.

These things have emergency resets if you mess it up... see your owners manual.

** Must have Chrome installed for this to work.

If you're forking or working on this or what have you, you will want to .gitignore your config.json. I am leaving it on here as an example. I don't plan on doing a lot of work with this app whilst monkeying with actual production data.

This works with the physical USB connection to the device.

This program will load the values in config.json into your hotspot device:

Hopefully, the fields are self-explanatory:

    "current_admin_password": "UniqueStickerPass123",
    "new_admin_password": "YourAdminPass_",
    "wifi_24_pw": "Secure24ghzPass", #wifi 2.4ghz password
    "wifi_5g_pw": "HighSpeed5ghzPass", #wifi 2.5 password
    "ssid": "Orbic_Hotspot_Alpha", #gets plugged into both 2.4 and 5g
    "dns_mode": "Manual",
    "dns_pri": "8.8.8.8",
    "dns_sec": "8.8.4.4"
