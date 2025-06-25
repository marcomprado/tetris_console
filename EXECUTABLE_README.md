# 🚀 Tetris Console - Executável

Executável único que baixa e executa o Tetris Console automaticamente!

## 🚀 Launcher Multiplataforma

Agora você pode baixar e executar o Tetris Console com apenas um clique!

### 📥 Downloads

- **Windows 64-bit**: tetris-launcher.exe
- **Windows 32-bit**: tetris-launcher-32.exe  
- **macOS Intel**: tetris-launcher-mac
- **macOS M1/M2**: tetris-launcher-mac-m1
- **Linux 64-bit**: tetris-launcher-linux
- **Linux 32-bit**: tetris-launcher-linux-32
- **Linux ARM64**: tetris-launcher-linux-arm64

### 🎯 Como usar

1. Baixe o arquivo para sua plataforma
2. Execute o arquivo
3. O launcher baixará e executará o Tetris automaticamente!

### 📋 Pré-requisitos

- .NET 9.0+ instalado
- Git instalado

Mais detalhes em [EXECUTABLE_README.md](./EXECUTABLE_README.md)

## 🚀 Baixe e Rode com 1 Clique!

Agora você pode baixar o executável para sua plataforma e rodar o Tetris Console automaticamente:

- [Windows 64-bit](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher.exe)
- [macOS Intel](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-mac)
- [macOS M1/M2](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-mac-m1)
- [Linux 64-bit](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-linux)
- [Linux ARM64](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-linux-arm64)

**Como usar:**  
Baixe, dê permissão de execução (no macOS/Linux: `chmod +x nome-do-arquivo`), e execute!

Mais detalhes em [EXECUTABLE_README.md](./EXECUTABLE_README.md)

## 📥 Download

Baixe o executável para sua plataforma:

### Windows
- **64-bit**: [tetris-launcher.exe](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher.exe)
- **32-bit**: [tetris-launcher-32.exe](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-32.exe)

### macOS
- **Intel**: [tetris-launcher-mac](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-mac)
- **Apple Silicon (M1/M2)**: [tetris-launcher-mac-m1](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-mac-m1)

### Linux
- **64-bit**: [tetris-launcher-linux](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-linux)
- **32-bit**: [tetris-launcher-linux-32](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-linux-32)
- **ARM64**: [tetris-launcher-linux-arm64](https://github.com/marcomprado/tetris-console/releases/latest/download/tetris-launcher-linux-arm64)

## 🎯 Como Usar

### Windows
1. Baixe o arquivo `.exe`
2. Dê duplo clique no executável
3. Ou execute no cmd: `tetris-launcher.exe`

### macOS
1. Baixe o arquivo para sua arquitetura
2. Abra o Terminal
3. Navegue até a pasta onde baixou
4. Execute:
   ```bash
   chmod +x tetris-launcher-mac
   ./tetris-launcher-mac
   ```

### Linux
1. Baixe o arquivo para sua arquitetura
2. Abra o terminal
3. Navegue até a pasta onde baixou
4. Execute:
   ```bash
   chmod +x tetris-launcher-linux
   ./tetris-launcher-linux
   ```

## 🔧 O que o Executável Faz

1. **🔍 Detecta o sistema operacional** automaticamente
2. **📥 Baixa o launcher apropriado** do GitHub
3. **🚀 Executa o launcher** que:
   - Verifica dependências (.NET e Git)
   - Clona o repositório do Tetris
   - Restaura dependências
   - Inicia o jogo

## 📋 Pré-requisitos

O executável ainda precisa das dependências instaladas:

### .NET 9.0
- **Windows**: https://dotnet.microsoft.com/download/dotnet/9.0
- **macOS**: `brew install dotnet`
- **Linux**: `sudo apt-get install dotnet-sdk-9.0`

### Git
- **Windows**: https://git-scm.com/download/win
- **macOS**: `brew install git`
- **Linux**: `sudo apt-get install git`

## 🛠️ Compilando Localmente

Se quiser compilar o executável:

```bash
# Instalar Go (se não tiver)
# https://golang.org/dl/

# Compilar para sua plataforma
go build -o tetris-launcher launcher.go

# Compilar para todas as plataformas
go run build.go
```

## 🔒 Segurança

- O executável baixa apenas scripts do repositório oficial
- Não requer privilégios de administrador
- Código fonte aberto e verificável
- Usa HTTPS para downloads

## 🐛 Solução de Problemas

### "Executável não encontrado"
- Verifique se baixou o arquivo correto para sua plataforma
- Confirme se o arquivo tem permissão de execução

### "Erro de download"
- Verifique sua conexão com a internet
- Confirme se o GitHub está acessível

### "Erro de execução"
- Verifique se .NET e Git estão instalados
- Execute manualmente os comandos do launcher

## 📊 Tamanho dos Executáveis

- **Windows**: ~2-3 MB
- **macOS**: ~2-3 MB  
- **Linux**: ~2-3 MB

## 🔄 Atualizações

O executável sempre baixa a versão mais recente do launcher, garantindo que você sempre tenha a versão atualizada do Tetris Console.

---

**Apenas execute e jogue!** 🎮 