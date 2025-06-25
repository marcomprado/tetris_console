# Tetris Console 🎮

Um Tetris clássico de console, feito em C# com visual retrô, jogabilidade tradicional e arquitetura moderna!

---

## 📸 Preview

```
+------------------------------++------------------+
|                              ||  PONTOS:      450 |
|                              ||  LINHAS:        3 |
|                              ||  NIVEL:         1 |
|                              ||                  |
|                              ||   PROXIMA PECA   |
|                              ||        [][]        |
|                              ||        [][]        |
|                              ||                  |
|                              ||                  |
|                              ||                  |
|                              ||                  |
|        [][][]                ||    CONTROLES:    |
|         []                   ||<- -> / A D: Mover|
|                              ||  ^ / W: Girar    |
|[][]    [][][]               ||  v / S: Descer   |
|[][][]    []      [][]       ||  Espaco: Queda   |
|[][][][][][][][][][][]       ||  P: Pausar       |
|[][][][][][][][][][][][]     ||  ESC/Q: Sair     |
|[][][][][][][][][][][][][]   ||                  |
|[][][][][][][][][][][][][][] ||                  |
+------------------------------++------------------+
```
*Exemplo de tela do jogo rodando no terminal com blocos ASCII.*

---

## 🚀 Como Executar

**Pré-requisitos:**  
- [.NET 9.0+](https://dotnet.microsoft.com/download/dotnet/9.0)
- Terminal compatível (Windows, macOS, Linux)

```bash
git clone https://github.com/marcomprado/tetris-console.git
cd tetris-console
dotnet run
```

---

## 🎮 Controles

| Tecla         | Ação                        |
|---------------|-----------------------------|
| ← → ou A D    | Mover peça                  |
| ↑ ou W        | Girar peça                  |
| ↓ ou S        | Descida suave (soft drop)   |
| Espaço        | Queda rápida (hard drop)    |
| P             | Pausar/Despausar            |
| ESC ou Q      | Sair do jogo                |

---

## 🧩 Funcionalidades

- 7 tipos de peças clássicas (I, O, T, S, Z, J, L)
- Sistema de pontuação e níveis progressivos
- Preview da próxima peça
- Rotação com wall-kick
- Pausa e interface em português
- Visual minimalista e retrô
- **Compatibilidade universal** (Windows, macOS, Linux)

---

## 🔧 Melhorias da Branch ASC2

### ✅ Compatibilidade Universal
- **Blocos ASCII**: Substituição de caracteres Unicode por `[]` simples
- **Sem piscamento**: Eliminação de problemas de renderização no Windows
- **Performance otimizada**: Renderização mais rápida e estável

### ✅ Refatoração de Código
- **Código mais limpo**: Eliminação de múltiplos `if/else if` verbosos
- **Melhor organização**: Separação da lógica de renderização do painel lateral
- **Manutenibilidade**: Estrutura mais modular e fácil de manter

### ✅ Layout Corrigido
- **Painel lateral**: Texto dos controles ajustado para caber na largura
- **Centralização**: Preview de peças centralizado corretamente
- **Interface consistente**: Visual uniforme em todas as plataformas

---

## 🧵 Concorrência e Arquitetura

A partir da versão atual, o jogo utiliza **duas threads principais**:
- **Thread de lógica:** processa entradas, atualiza o estado do jogo e controla as peças.
- **Thread de renderização:** desenha o tabuleiro e painel no console, de forma independente.

Isso garante uma interface mais fluida e separação clara entre lógica e visual!

---

## 📁 Estrutura do Projeto

```
tetris-console/
├── Engine/                 # Lógica principal do jogo
├── Graphics/               # Renderização no console
├── Models/                 # Modelos de dados e regras
└── Program.cs              # Ponto de entrada
```

---

## 🛠️ Configurações

Edite `Models/GameConfig.cs` para:
- Tamanho do tabuleiro
- Velocidade inicial
- Sistema de pontuação

---

## 🐛 Problemas Conhecidos

- ~~Alguns terminais podem não exibir Unicode corretamente~~ ✅ **Resolvido na branch ASC2**
- Requer terminal com pelo menos 24 linhas

---

## 📝 Licença

MIT — veja o arquivo [LICENSE](LICENSE).

---

## 👨‍💻 Autor

Desenvolvido por [marcomprado](https://github.com/marcomprado) como projeto acadêmico para Sistemas Operacionais.

---

**Divirta-se jogando!** 🎮 