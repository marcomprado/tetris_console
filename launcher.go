package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

const (
	launcherURL = "https://raw.githubusercontent.com/marcomprado/tetris_console/main/launcher.sh"
	launcherBatURL = "https://raw.githubusercontent.com/marcomprado/tetris_console/main/launcher.bat"
)

func main() {
	fmt.Println("🎮 Tetris Console Launcher")
	fmt.Println("==========================")
	fmt.Println()

	// Determinar o sistema operacional
	osType := runtime.GOOS
	fmt.Printf("📱 Sistema detectado: %s\n", osType)

	// Criar diretório temporário
	tempDir, err := os.MkdirTemp("", "tetris-launcher")
	if err != nil {
		fmt.Printf("❌ Erro ao criar diretório temporário: %v\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(tempDir)

	var launcherPath string
	var downloadURL string

	// Baixar o launcher apropriado
	if osType == "windows" {
		launcherPath = filepath.Join(tempDir, "launcher.bat")
		downloadURL = launcherBatURL
		fmt.Println("📥 Baixando launcher para Windows...")
	} else {
		launcherPath = filepath.Join(tempDir, "launcher.sh")
		downloadURL = launcherURL
		fmt.Println("📥 Baixando launcher...")
	}

	// Baixar o arquivo
	err = downloadFile(downloadURL, launcherPath)
	if err != nil {
		fmt.Printf("❌ Erro ao baixar launcher: %v\n", err)
		fmt.Println("💡 Verifique sua conexão com a internet")
		os.Exit(1)
	}

	// Tornar executável no Unix
	if osType != "windows" {
		err = os.Chmod(launcherPath, 0755)
		if err != nil {
			fmt.Printf("❌ Erro ao tornar executável: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("✅ Launcher baixado com sucesso!")
	fmt.Println("🚀 Iniciando Tetris Console...")
	fmt.Println()

	// Executar o launcher
	var cmd *exec.Cmd
	if osType == "windows" {
		cmd = exec.Command("cmd", "/c", launcherPath)
	} else {
		cmd = exec.Command("/bin/bash", launcherPath)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Run()
	if err != nil {
		fmt.Printf("❌ Erro ao executar launcher: %v\n", err)
		os.Exit(1)
	}
}

func downloadFile(url, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro HTTP: %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
} 