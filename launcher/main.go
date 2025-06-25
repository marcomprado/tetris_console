package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	repoURL     = "https://github.com/marcomprado/tetris_console.git"
	repoName    = "tetris_console"
	projectPath = "tetris_console"
)

func main() {
	fmt.Println("=== Tetris Console Launcher ===")
	fmt.Println()

	// Verificar se o Git está instalado
	if !checkGit() {
		fmt.Println("❌ Git não encontrado!")
		fmt.Println("📥 Por favor, instale o Git primeiro:")
		fmt.Println("   Windows: https://git-scm.com/download/win")
		fmt.Println("   macOS: brew install git")
		fmt.Println("   Linux: sudo apt-get install git")
		pause()
		return
	} else {
		fmt.Println("✅ Git encontrado!")
	}

	// Verificar se o .NET 9 está instalado
	if !checkDotNet9() {
		fmt.Println("❌ .NET 9 não encontrado!")
		fmt.Println("📥 Instalando .NET 9...")
		if !installDotNet9() {
			fmt.Println("❌ Falha ao instalar .NET 9. Por favor, instale manualmente.")
			fmt.Println("   Visite: https://dotnet.microsoft.com/download/dotnet/9.0")
			pause()
			return
		}
	} else {
		fmt.Println("✅ .NET 9 encontrado!")
	}

	// Verificar se o repositório já existe
	repoPath := getRepoPath()
	fmt.Printf("📁 Caminho do repositório: %s\n", repoPath)
	
	if !repoExists(repoPath) {
		fmt.Println("📥 Baixando o jogo Tetris...")
		if !cloneRepo(repoPath) {
			fmt.Println("❌ Falha ao baixar o repositório.")
			fmt.Println("   Verifique sua conexão com a internet.")
			fmt.Println("   URL do repositório:", repoURL)
			pause()
			return
		}
		fmt.Println("✅ Jogo baixado com sucesso!")
	} else {
		fmt.Println("✅ Jogo já está baixado!")
		fmt.Println("🔄 Atualizando para a versão mais recente...")
		if !updateRepo(repoPath) {
			fmt.Println("⚠️  Falha ao atualizar, mas continuando com a versão local...")
		} else {
			fmt.Println("✅ Jogo atualizado!")
		}
	}

	// Executar o jogo
	fmt.Println()
	fmt.Println("🎮 Iniciando o jogo Tetris...")
	fmt.Println("Pressione Ctrl+C para sair do jogo.")
	fmt.Println()

	if !runGame(repoPath) {
		fmt.Println("❌ Falha ao executar o jogo.")
		pause()
		return
	}
}

func checkGit() bool {
	cmd := exec.Command("git", "--version")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	
	version := strings.TrimSpace(string(output))
	fmt.Printf("🔍 Git encontrado: %s\n", version)
	return true
}

func checkDotNet9() bool {
	cmd := exec.Command("dotnet", "--version")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	version := strings.TrimSpace(string(output))
	fmt.Printf("🔍 .NET encontrado: %s\n", version)
	return strings.HasPrefix(version, "9.")
}

func installDotNet9() bool {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// Para Windows, baixar e executar o instalador
		fmt.Println("📥 Baixando instalador do .NET 9 para Windows...")
		cmd = exec.Command("powershell", "-Command", 
			"Invoke-WebRequest -Uri 'https://download.microsoft.com/download/8/4/c/84c6c430-e0f5-476d-bf57-0c24e0c5e5c8/dotnet-sdk-9.0.100-win-x64.exe' -OutFile 'dotnet-installer.exe'")
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Erro ao baixar instalador: %v\n", err)
			return false
		}

		fmt.Println("🔧 Instalando .NET 9...")
		cmd = exec.Command("dotnet-installer.exe", "/quiet", "/norestart")
		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Erro ao instalar .NET: %v\n", err)
			return false
		}

		// Limpar o instalador
		os.Remove("dotnet-installer.exe")

	case "darwin":
		// Para macOS, usar Homebrew
		fmt.Println("📥 Instalando .NET 9 via Homebrew...")
		cmd = exec.Command("brew", "install", "dotnet@9")

	case "linux":
		// Para Linux, usar o script oficial
		fmt.Println("📥 Instalando .NET 9 via script oficial...")
		cmd = exec.Command("bash", "-c", 
			"curl -sSL https://dot.net/v1/dotnet-install.sh | bash /dev/stdin --channel 9.0")
	}

	if cmd != nil {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run() == nil
	}

	return false
}

func getRepoPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return repoName
	}
	return filepath.Join(homeDir, ".tetris_console")
}

func repoExists(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".git"))
	return err == nil
}

func cloneRepo(path string) bool {
	// Criar diretório pai se não existir
	parentDir := filepath.Dir(path)
	if err := os.MkdirAll(parentDir, 0755); err != nil {
		fmt.Printf("❌ Erro ao criar diretório: %v\n", err)
		return false
	}

	fmt.Printf("🔍 Clonando repositório para: %s\n", path)
	cmd := exec.Command("git", "clone", repoURL, path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Erro ao clonar repositório: %v\n", err)
		return false
	}
	
	return true
}

func updateRepo(path string) bool {
	cmd := exec.Command("git", "pull", "origin", "main")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run() == nil
}

func runGame(path string) bool {
	projectFile := filepath.Join(path, "tetris_so.csproj")
	if _, err := os.Stat(projectFile); os.IsNotExist(err) {
		fmt.Printf("❌ Arquivo do projeto não encontrado em: %s\n", projectFile)
		fmt.Printf("📁 Conteúdo do diretório %s:\n", path)
		listDirectory(path)
		return false
	}

	cmd := exec.Command("dotnet", "run", "--project", projectFile)
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run() == nil
}

func listDirectory(path string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("❌ Erro ao listar diretório: %v\n", err)
		return
	}
	
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("📁 %s/\n", entry.Name())
		} else {
			fmt.Printf("📄 %s\n", entry.Name())
		}
	}
}

func pause() {
	fmt.Println()
	fmt.Println("Pressione Enter para sair...")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
} 