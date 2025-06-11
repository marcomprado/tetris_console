using TetrisConsole.Models;
using System.Drawing;

namespace TetrisConsole.Graphics;

public class ConsoleRenderer
{
    private readonly Board _board;
    private readonly GameState _gameState;
    private const char EMPTY_CHAR = ' ';
    private const char BLOCK_CHAR = '■';

    public ConsoleRenderer(Board board, GameState gameState)
    {
        _board = board;
        _gameState = gameState;
        Console.CursorVisible = false;
        Console.Clear();
    }

    public void Render()
    {
        // Limpa o console e reposiciona o cursor no topo
        Console.Clear();
        Console.SetCursorPosition(0, 0);
        DrawFrame();
    }

    public void RenderWithMessage(string message)
    {
        Render();
        DisplayCenteredMessage(message);
    }

    private void DisplayCenteredMessage(string message)
    {
        int boardWidth = _board.Columns * 2;
        int messageX = (boardWidth - message.Length) / 2 + 1;
        int messageY = _board.Rows / 2;

        Console.SetCursorPosition(messageX, messageY);
        Console.BackgroundColor = ConsoleColor.DarkBlue;
        Console.ForegroundColor = ConsoleColor.White;
        Console.Write(" " + message + " ");
        Console.ResetColor();
    }

    private void DrawFrame()
    {
        // Desenha borda superior
        Console.Write("╔" + new string('═', _board.Columns * 2) + "╗");
        Console.WriteLine("╔" + new string('═', GameConfig.SidePanelWidth - 2) + "╗");

        // Desenha o tabuleiro e painel lateral
        for (int row = 0; row < _board.Rows; row++)
        {
            // Desenha linha do tabuleiro
            Console.Write("║");
            for (int col = 0; col < _board.Columns; col++)
            {
                char blockChar = GetCellChar(col, row);
                Console.Write(blockChar + " ");
            }
            Console.Write("║");
            
            // Desenha linha do painel lateral
            if (row == 0)
                Console.Write("║  PONTOS:" + _gameState.Score.ToString().PadLeft(8) + " ║");
            else if (row == 1)
                Console.Write("║  LINHAS:" + _gameState.Lines.ToString().PadLeft(8) + " ║");
            else if (row == 2)
                Console.Write("║  NÍVEL: " + _gameState.Level.ToString().PadLeft(8) + " ║");
            else if (row == 4)
                Console.Write("║   PRÓXIMA PEÇA   ║");
            else if (row >= 5 && row < 5 + GameConfig.NextPiecePreviewSize && _gameState.NextPiece != null)
            {
                Console.Write("║     ");
                DrawNextPiecePreview(row - 5);
                Console.Write("     ║");
            }
            else if (row == 12)
                Console.Write("║    CONTROLES:    ║");
            else if (row == 13)
                Console.Write("║  ← → / A D: Mover║");
            else if (row == 14)
                Console.Write("║  ↑ / W: Girar    ║");
            else if (row == 15)
                Console.Write("║  ↓ / S: Descer   ║");
            else if (row == 16)
                Console.Write("║  Espaço: Queda   ║");
            else if (row == 17)
                Console.Write("║  P: Pausar       ║");
            else if (row == 18)
                Console.Write("║  ESC/Q: Sair     ║");
            else
                Console.Write("║" + new string(' ', GameConfig.SidePanelWidth - 2) + "║");
            
            Console.WriteLine();
        }

        // Desenha borda inferior
        Console.Write("╚" + new string('═', _board.Columns * 2) + "╝");
        Console.Write("╚" + new string('═', GameConfig.SidePanelWidth - 2) + "╝");
    }

    private char GetCellChar(int col, int row)
    {
        // Primeiro verifica se a peça atual ocupa esta célula
        if (_gameState.CurrentPiece != null)
        {
            var piece = _gameState.CurrentPiece;
            int pieceRow = row - piece.Position.Y;
            int pieceCol = col - piece.Position.X;

            if (pieceRow >= 0 && pieceRow < piece.Shape.GetLength(0) &&
                pieceCol >= 0 && pieceCol < piece.Shape.GetLength(1) &&
                piece.Shape[pieceRow, pieceCol] != 0)
            {
                return BLOCK_CHAR;
            }
        }

        // Caso contrário, obtém a célula do tabuleiro
        int cell = _board.GetCell(new Point(col, row));
        return cell != 0 ? BLOCK_CHAR : EMPTY_CHAR;
    }

    private void DrawNextPiecePreview(int row)
    {
        const int PREVIEW_WIDTH = 8; // Largura total da área de preview
        
        if (_gameState.NextPiece == null || row >= _gameState.NextPiece.Shape.GetLength(0))
        {
            Console.Write(new string(' ', PREVIEW_WIDTH));
            return;
        }

        // Calcula a largura real da peça nesta linha
        int pieceWidth = 0;
        int firstBlock = -1;
        int lastBlock = -1;
        
        for (int col = 0; col < _gameState.NextPiece.Shape.GetLength(1); col++)
        {
            if (_gameState.NextPiece.Shape[row, col] > 0)
            {
                if (firstBlock == -1) firstBlock = col;
                lastBlock = col;
            }
        }
        
        if (firstBlock == -1)
        {
            // Linha vazia
            Console.Write(new string(' ', PREVIEW_WIDTH));
            return;
        }
        
        pieceWidth = lastBlock - firstBlock + 1;
        int leftPadding = (PREVIEW_WIDTH - pieceWidth) / 2;
        
        // Desenha com centralização adequada
        Console.Write(new string(' ', leftPadding));
        
        for (int col = firstBlock; col <= lastBlock; col++)
        {
            char blockChar = _gameState.NextPiece.Shape[row, col] > 0 
                ? BLOCK_CHAR 
                : EMPTY_CHAR;
            Console.Write(blockChar);
        }
        
        int rightPadding = PREVIEW_WIDTH - leftPadding - pieceWidth;
        Console.Write(new string(' ', rightPadding));
    }
} 