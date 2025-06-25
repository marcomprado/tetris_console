package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type BuildTarget struct {
	OS   string
	Arch string
	Ext  string
}

func main() {
	targets := []BuildTarget{
		{OS: "windows", Arch: "amd64", Ext: ".exe"},
		{OS: "windows", Arch: "386", Ext: ".exe"},
		{OS: "linux", Arch: "amd64", Ext: ""},
		{OS: "linux", Arch: "386", Ext: ""},
		{OS: "linux", Arch: "arm64", Ext: ""},
		{OS: "darwin", Arch: "amd64", Ext: ""},
		{OS: "darwin", Arch: "arm64", Ext: ""},
	}

	// Criar diretório builds se não existir
	if err := os.MkdirAll("builds", 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório builds: %v\n", err)
		return
	}

	fmt.Println("🔨 Compilando Tetris Launcher para múltiplas plataformas...")
	fmt.Println()

	for _, target := range targets {
		filename := fmt.Sprintf("tetris-launcher-%s-%s%s", target.OS, target.Arch, target.Ext)
		outputPath := filepath.Join("builds", filename)

		fmt.Printf("📦 Compilando para %s/%s... ", target.OS, target.Arch)

		cmd := exec.Command("go", "build", "-o", outputPath, "main.go")
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("GOOS=%s", target.OS),
			fmt.Sprintf("GOARCH=%s", target.Arch),
		)

		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Erro: %v\n", err)
		} else {
			fmt.Printf("✅ %s\n", filename)
		}
	}

	fmt.Println()
	fmt.Println("🎉 Compilação concluída!")
	fmt.Println("📁 Executáveis disponíveis na pasta 'builds/'")
} 