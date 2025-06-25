#!/bin/bash

# Tetris Console Launcher
# Script simples para baixar e executar o Tetris Console

set -e

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configurações
REPO_URL="https://github.com/marcomprado/tetris_console.git"
PROJECT_DIR="tetris-console"
MIN_DOTNET_VERSION="9.0"

echo -e "${BLUE}🎮 Tetris Console Launcher${NC}"
echo -e "${BLUE}========================${NC}"
echo ""

# Função para verificar se o .NET está instalado
check_dotnet() {
    if ! command -v dotnet &> /dev/null; then
        echo -e "${RED}❌ .NET não encontrado!${NC}"
        echo -e "${YELLOW}📥 Instalando .NET 9.0...${NC}"
        
        if [[ "$OSTYPE" == "darwin"* ]]; then
            # macOS
            echo -e "${YELLOW}Para macOS, instale o .NET via:${NC}"
            echo "brew install dotnet"
            echo ""
            echo "Ou baixe de: https://dotnet.microsoft.com/download/dotnet/9.0"
        elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
            # Linux
            echo -e "${YELLOW}Para Linux, instale o .NET via:${NC}"
            echo "sudo apt-get update && sudo apt-get install -y dotnet-sdk-9.0"
            echo ""
            echo "Ou baixe de: https://dotnet.microsoft.com/download/dotnet/9.0"
        else
            echo -e "${YELLOW}Baixe o .NET 9.0 de: https://dotnet.microsoft.com/download/dotnet/9.0${NC}"
        fi
        exit 1
    fi
    
    # Verificar versão do .NET
    DOTNET_VERSION=$(dotnet --version)
    echo -e "${GREEN}✅ .NET $DOTNET_VERSION encontrado${NC}"
}

# Função para clonar ou atualizar o repositório
setup_project() {
    if [ -d "$PROJECT_DIR" ]; then
        echo -e "${YELLOW}📁 Projeto já existe. Atualizando...${NC}"
        cd "$PROJECT_DIR"
        git pull origin main
        cd ..
    else
        echo -e "${YELLOW}📥 Clonando repositório...${NC}"
        git clone "$REPO_URL" "$PROJECT_DIR"
    fi
    echo -e "${GREEN}✅ Projeto configurado${NC}"
}

# Função para restaurar dependências
restore_dependencies() {
    echo -e "${YELLOW}📦 Restaurando dependências...${NC}"
    cd "$PROJECT_DIR"
    dotnet restore
    cd ..
    echo -e "${GREEN}✅ Dependências restauradas${NC}"
}

# Função para executar o jogo
run_game() {
    echo -e "${GREEN}🚀 Iniciando Tetris Console...${NC}"
    echo -e "${BLUE}Pressione Ctrl+C para sair${NC}"
    echo ""
    cd "$PROJECT_DIR"
    dotnet run
}

# Função principal
main() {
    echo -e "${YELLOW}🔍 Verificando dependências...${NC}"
    check_dotnet
    
    echo -e "${YELLOW}📂 Configurando projeto...${NC}"
    setup_project
    
    echo -e "${YELLOW}🔧 Preparando ambiente...${NC}"
    restore_dependencies
    
    echo ""
    echo -e "${GREEN}🎯 Tudo pronto!${NC}"
    echo ""
    
    run_game
}

# Executar função principal
main "$@" 