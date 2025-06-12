# Tetris Console 🎮

Um jogo de Tetris clássico implementado em C# para rodar no console/terminal, com visual retrô e jogabilidade tradicional.

## 📸 Preview

```
╔══════════════════════╗╔══════════════════╗
║                      ║║  PONTOS:      450║
║                      ║║  LINHAS:        3║
║                      ║║  NÍVEL:         1║
║                      ║║                  ║
║                      ║║   PRÓXIMA PEÇA   ║
║                      ║║       ■■         ║
║                      ║║       ■■         ║
║                      ║║                  ║
║                      ║║                  ║
║                      ║║                  ║
║        ■■■           ║║                  ║
║         ■            ║║    CONTROLES:    ║
║                      ║║  ← → / A D: Mover║
║                      ║║  ↑ / W: Girar    ║
║■■      ■■■           ║║  ↓ / S: Descer   ║
║■■■      ■      ■■    ║║  Espaço: Queda   ║
║■■■■■■■■■■■■■■■■■■    ║║  P: Pausar       ║
║■■■■■■■■■■■■■■■■■■■■  ║║  ESC/Q: Sair     ║
║■■■■■■■■■■■■■■■■■■■■■ ║║                  ║
║■■■■■■■■■■■■■■■■■■■■■■║║                  ║
╚══════════════════════╝╚══════════════════╝
```

## 🎯 Funcionalidades

- ✅ **7 tipos de peças clássicas** (I, O, T, S, Z, J, L)
- ✅ **Sistema de pontuação** com bonificações por descida rápida
- ✅ **Níveis progressivos** - a velocidade aumenta conforme você avança
- ✅ **Preview da próxima peça**
- ✅ **Rotação de peças** com sistema de wall-kick
- ✅ **Pausa do jogo**
- ✅ **Interface totalmente em português**
- ✅ **Visual minimalista e retrô**

## 🚀 Como Executar

### Pré-requisitos

- [.NET 9.0](https://dotnet.microsoft.com/download/dotnet/9.0) ou superior
- Terminal com suporte a caracteres Unicode

### Instalação e Execução

1. Clone o repositório:
```bash
git clone https://github.com/seu-usuario/tetris-console.git
cd tetris-console
```

2. Compile e execute:
```bash
dotnet run
```

## 🎮 Como Jogar

### Controles

| Tecla | Ação |
|-------|------|
| ← → ou A D | Mover peça para esquerda/direita |
| ↑ ou W | Girar peça |
| ↓ ou S | Descida suave (soft drop) |
| Espaço | Queda rápida (hard drop) |
| P | Pausar/Despausar |
| ESC ou Q | Sair do jogo |

### Sistema de Pontuação

- **Linha completa**: 100 pontos × nível atual
- **Descida suave**: 1 ponto por linha
- **Queda rápida**: 2 pontos por linha
- **Múltiplas linhas**: Pontos são multiplicados pelo número de linhas

### Níveis

- O nível aumenta a cada 10 linhas completas
- Cada nível aumenta a velocidade de queda das peças

## 📁 Estrutura do Projeto

```
tetris-console/
├── Engine/                 # Lógica principal do jogo
│   ├── GameEngine.cs      # Motor do jogo
│   ├── InputHandler.cs    # Gerenciamento de entrada
│   └── InputAction.cs     # Enum de ações
├── Graphics/              # Renderização
│   └── ConsoleRenderer.cs # Renderização no console
├── Models/                # Modelos de dados
│   ├── Board.cs          # Tabuleiro do jogo
│   ├── GameConfig.cs     # Configurações
│   ├── GameState.cs      # Estado do jogo
│   ├── Piece.cs          # Modelo de peça
│   └── PieceFactory.cs   # Criação de peças
└── Program.cs            # Ponto de entrada
```

## 🛠️ Detalhes Técnicos

### Arquitetura

O jogo segue uma arquitetura limpa com separação clara de responsabilidades:

- **Engine**: Contém toda a lógica do jogo, processamento de entrada e game loop
- **Graphics**: Responsável por toda a renderização no console
- **Models**: Classes de dados e lógica de domínio

### Características do Código

- Código totalmente em português (comentários e interface)
- Uso de caracteres Unicode para bordas e blocos
- Sistema de configuração centralizado
- Renderização otimizada com limpeza de tela

## 🔧 Configurações

As configurações do jogo podem ser ajustadas em `Models/GameConfig.cs`:

- Dimensões do tabuleiro (padrão: 20x10)
- Velocidade inicial
- Sistema de pontuação
- Tamanhos da interface

## 🐛 Problemas Conhecidos

- Em alguns terminais, os caracteres Unicode podem não ser exibidos corretamente
- O jogo requer um terminal com pelo menos 24 linhas de altura

## 🚧 Melhorias Futuras

- [ ] Sistema de high scores persistente
- [ ] Efeitos sonoros
- [ ] Modo de jogo fantasma (ghost piece)
- [ ] Estatísticas detalhadas
- [ ] Diferentes modos de jogo
- [ ] Customização de teclas

## 📝 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

## 👨‍💻 Autor

Desenvolvido como projeto acadêmico para a disciplina de Sistemas Operacionais.

---

**Divirta-se jogando! 🎮** 