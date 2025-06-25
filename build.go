package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	fmt.Println("🔨 Compilando executáveis cross-platform...")
	fmt.Println()

	// Configurações de compilação
	builds := []struct {
		os   string
		arch string
		name string
	}{
		{"windows", "amd64", "tetris-launcher.exe"},
		{"windows", "386", "tetris-launcher-32.exe"},
		{"darwin", "amd64", "tetris-launcher-mac"},
		{"darwin", "arm64", "tetris-launcher-mac-m1"},
		{"linux", "amd64", "tetris-launcher-linux"},
		{"linux", "386", "tetris-launcher-linux-32"},
		{"linux", "arm64", "tetris-launcher-linux-arm64"},
	}

	// Criar diretório de builds
	buildDir := "builds"
	if err := os.MkdirAll(buildDir, 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório builds: %v\n", err)
		os.Exit(1)
	}

	// Compilar para cada plataforma
	for _, build := range builds {
		fmt.Printf("📦 Compilando para %s/%s...\n", build.os, build.arch)
		
		outputPath := filepath.Join(buildDir, build.name)
		
		cmd := exec.Command("go", "build", 
			"-o", outputPath,
			"-ldflags", "-s -w", // Reduzir tamanho do executável
			"launcher.go")
		
		cmd.Env = append(os.Environ(),
			fmt.Sprintf("GOOS=%s", build.os),
			fmt.Sprintf("GOARCH=%s", build.arch),
		)
		
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Erro ao compilar %s: %v\n", build.name, err)
			continue
		}
		
		// Verificar tamanho do arquivo
		if info, err := os.Stat(outputPath); err == nil {
			size := info.Size()
			fmt.Printf("✅ %s (%d bytes)\n", build.name, size)
		}
	}

	fmt.Println()
	fmt.Println("🎉 Compilação concluída!")
	fmt.Printf("📁 Executáveis criados em: %s/\n", buildDir)
} 