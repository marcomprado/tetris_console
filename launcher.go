package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
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
	fmt.Printf("🔗 Tentando baixar de: %s\n", downloadURL)
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
		cmd = exec.Command("sh", launcherPath)
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
	// Criar contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Criar requisição com contexto
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("erro ao criar requisição: %v", err)
	}

	// Configurar User-Agent para evitar bloqueios
	req.Header.Set("User-Agent", "Tetris-Launcher/1.0")

	fmt.Printf("📡 Fazendo requisição HTTP para: %s\n", url)
	
	// Fazer a requisição
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("erro na requisição HTTP: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("📊 Status code: %d\n", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("erro HTTP: %d - %s", resp.StatusCode, resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %v", err)
	}
	defer out.Close()

	bytesWritten, err := io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("erro ao copiar dados: %v", err)
	}
	
	fmt.Printf("✅ Download concluído: %d bytes\n", bytesWritten)
	return nil
} 