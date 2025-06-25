@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

REM Tetris Console Launcher para Windows
REM Script simples para baixar e executar o Tetris Console

echo 🎮 Tetris Console Launcher
echo ========================
echo.

REM Configurações
set REPO_URL=https://github.com/marcomprado/tetris_console.git
set PROJECT_DIR=tetris_console

REM Verificar se o .NET está instalado
echo 🔍 Verificando dependências...
dotnet --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ .NET não encontrado!
    echo 📥 Instalando .NET 9.0...
    echo.
    echo Para Windows, baixe o .NET 9.0 de:
    echo https://dotnet.microsoft.com/download/dotnet/9.0
    echo.
    echo Ou use o winget:
    echo winget install Microsoft.DotNet.SDK.9
    echo.
    pause
    exit /b 1
)

for /f "tokens=*" %%i in ('dotnet --version') do set DOTNET_VERSION=%%i
echo ✅ .NET !DOTNET_VERSION! encontrado

REM Verificar se o Git está instalado
echo 🔍 Verificando Git...
git --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Git não encontrado!
    echo 📥 Instale o Git de: https://git-scm.com/download/win
    echo.
    echo Ou use o winget:
    echo winget install Git.Git
    echo.
    pause
    exit /b 1
)

echo ✅ Git encontrado

REM Configurar projeto
echo 📂 Configurando projeto...
if exist "%PROJECT_DIR%" (
    echo 📁 Projeto já existe. Atualizando...
    cd "%PROJECT_DIR%"
    git pull origin main
    if %errorlevel% neq 0 (
        git pull origin master
    )
    cd ..
) else (
    echo 📥 Clonando repositório...
    git clone "%REPO_URL%" "%PROJECT_DIR%"
)

if %errorlevel% neq 0 (
    echo ❌ Erro ao clonar/atualizar o repositório
    pause
    exit /b 1
)

echo ✅ Projeto configurado

REM Restaurar dependências
echo 📦 Restaurando dependências...
cd "%PROJECT_DIR%"
dotnet restore
if %errorlevel% neq 0 (
    echo ❌ Erro ao restaurar dependências
    cd ..
    pause
    exit /b 1
)
cd ..

echo ✅ Dependências restauradas

REM Executar o jogo
echo.
echo 🎯 Tudo pronto!
echo.
echo 🚀 Iniciando Tetris Console...
echo 💡 Pressione Ctrl+C para sair
echo.

cd "%PROJECT_DIR%"
dotnet run

if %errorlevel% neq 0 (
    echo ❌ Erro ao executar o jogo
    cd ..
    pause
    exit /b 1
)

cd ..
pause 