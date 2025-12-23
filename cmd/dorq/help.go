package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func displayHelp() {
	cmdStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	argStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	flagStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
	requiredStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	successStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))

	fmt.Println()
	fmt.Println(successStyle.Render(" example:"))
	fmt.Printf("    %s -q %s\n", cmdStyle.Render("dorq"), argStyle.Render("site:hackerone.com inurl:login"))
	fmt.Printf("    %s -q %s -n %s -json\n\n", cmdStyle.Render("dorq"), argStyle.Render("inurl:admin"), argStyle.Render("20"))

	fmt.Println(successStyle.Render(" options:"))
	fmt.Printf("    %s      search query %s\n", flagStyle.Render("-q"), requiredStyle.Render("(required)"))
	fmt.Printf("    %s      max results %s\n", flagStyle.Render("-n"), argStyle.Render("(default: 10)"))
	fmt.Printf("    %s   stdout as JSON format\n", flagStyle.Render("-json"))
	fmt.Printf("    %s      show version\n", flagStyle.Render("-v"))
	fmt.Printf("    %s     update to latest version\n", flagStyle.Render("-up"))
	fmt.Printf("    %s      show this help message\n", flagStyle.Render("-h"))
	fmt.Printf("    %s   comma separated sources %s\n", flagStyle.Render("-sources"), argStyle.Render("(default: google,bing,duckduckgo)"))
	fmt.Printf("    %s           timeout in seconds %s\n", flagStyle.Render("-timeout"), argStyle.Render("(default: 20)"))
	fmt.Printf("    %s      proxy for browser requests\n", flagStyle.Render("-request-proxy"))
	fmt.Printf("    %s     proxy to forward found URLs\n", flagStyle.Render("-forward-proxy"))
	fmt.Printf("    %s    show browser windows for captcha\n", flagStyle.Render("-show-browser"))
	fmt.Printf("    %s timeout for antibot resolution %s\n", flagStyle.Render("-antibot-timeout"), argStyle.Render("(default: 60)"))
	fmt.Printf("    %s             debug mode\n", flagStyle.Render("-debug"))
	fmt.Printf("    %s          verbose output\n", flagStyle.Render("-verbose"))
	fmt.Printf("    %s       very verbose output\n", flagStyle.Render("-vverbose"))
	fmt.Printf("    %s resubmit search without subdomains\n", flagStyle.Render("-resubmit-without-subs"))
	fmt.Printf("    %s            output file\n\n", flagStyle.Render("-output"))

	fmt.Println(argStyle.Render("usage of google dorking for attacking targets without prior mutual consent is illegal!"))
	fmt.Println()
}
