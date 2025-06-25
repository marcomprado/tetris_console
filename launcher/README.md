# Tetris Console Launcher

Este é um launcher automatizado para o jogo Tetris Console. Ele baixa automaticamente o jogo, verifica/instala o .NET 9 e executa o jogo.

## Funcionalidades

- ✅ Verifica se o .NET 9 está instalado
- ✅ Instala o .NET 9 automaticamente se necessário
- ✅ Baixa o jogo do GitHub na primeira execução
- ✅ Atualiza o jogo automaticamente nas execuções subsequentes
- ✅ Executa o jogo com `dotnet run`
- ✅ Interface amigável no console

## Como usar

1. Baixe o executável para sua plataforma
2. Execute o arquivo
3. O launcher fará todo o resto automaticamente!

## Plataformas suportadas

- Windows (x64, x86)
- macOS (Intel, Apple Silicon)
- Linux (x64, x86, ARM64)

## Requisitos

- Git instalado no sistema
- Conexão com a internet (para download inicial e atualizações)

## Como compilar

Se você quiser compilar o launcher por conta própria:

```bash
cd launcher
go run build.go
```

Os executáveis serão criados na pasta `builds/`.

## Estrutura do projeto

```
launcher/
├── main.go      # Código principal do launcher
├── build.go     # Script de compilação
├── go.mod       # Dependências do Go
├── README.md    # Este arquivo
└── builds/      # Executáveis compilados
```

## Como funciona

1. **Verificação do .NET 9**: O launcher verifica se o .NET 9 está instalado executando `dotnet --version`
2. **Instalação automática**: Se não estiver instalado, tenta instalar automaticamente:
   - Windows: Baixa e executa o instalador oficial
   - macOS: Usa Homebrew
   - Linux: Usa o script oficial da Microsoft
3. **Download do jogo**: Clona o repositório para `~/.tetris_console/`
4. **Atualização**: Nas execuções subsequentes, faz `git pull` para atualizar
5. **Execução**: Executa `dotnet run` no projeto

## Solução de problemas

### Erro ao instalar .NET 9
Se a instalação automática falhar, instale manualmente:
- Windows: https://dotnet.microsoft.com/download/dotnet/9.0
- macOS: `brew install dotnet@9`
- Linux: Siga as instruções em https://docs.microsoft.com/en-us/dotnet/core/install/linux

### Erro ao baixar o jogo
Verifique se:
- Git está instalado
- Há conexão com a internet
- O repositório está acessível

### Erro ao executar o jogo
Verifique se:
- O .NET 9 foi instalado corretamente
- O projeto foi baixado completamente
- Há permissões suficientes no diretório 