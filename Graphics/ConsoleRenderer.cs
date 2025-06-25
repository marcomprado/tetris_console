using TetrisConsole.Models;
using System.Drawing;

namespace TetrisConsole.Graphics;

public class ConsoleRenderer
{
    private readonly Board _board;
    private readonly GameState _gameState;
    private const string EMPTY_STRING = "  ";
    private const string BLOCK_STRING = "[]";

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
        DrawPerformanceInfo();
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
        Console.Write("+" + new string('-', _board.Columns * 2) + "+");
        Console.WriteLine("+" + new string('-', GameConfig.SidePanelWidth - 2) + "+");

        // Desenha o tabuleiro e painel lateral
        for (int row = 0; row < _board.Rows; row++)
        {
            // Desenha linha do tabuleiro
            Console.Write("|");
            for (int col = 0; col < _board.Columns; col++)
            {
                string blockString = GetCellString(col, row);
                Console.Write(blockString);
            }
            Console.Write("|");
            
            // Desenha linha do painel lateral
            Console.Write(GetSidePanelText(row));
            Console.WriteLine();
        }

        // Desenha borda inferior
        Console.Write("+" + new string('-', _board.Columns * 2) + "+");
        Console.Write("+" + new string('-', GameConfig.SidePanelWidth - 2) + "+");
    }

    private void DrawPerformanceInfo()
    {
        // Posiciona o cursor abaixo do jogo para mostrar informações de performance
        Console.SetCursorPosition(0, _board.Rows + 2);
        
        var threadInfo = ThreadMonitor.GetThreadInfo();
        
        // Desenha uma linha separadora
        Console.WriteLine("-" + new string('-', (_board.Columns * 2) + GameConfig.SidePanelWidth + 1));
        
        // Linha 1: Título e total de threads em destaque
        Console.ForegroundColor = ConsoleColor.Cyan;
        Console.Write("MONITOR DE PERFORMANCE");
        Console.ResetColor();
        Console.Write(" | ");
        Console.ForegroundColor = ConsoleColor.White;
        Console.BackgroundColor = ConsoleColor.DarkBlue;
        Console.Write($" TOTAL THREADS: {threadInfo.TotalThreads} ");
        Console.ResetColor();
        Console.Write($" | Memoria: ");
        Console.ForegroundColor = ConsoleColor.Green;
        Console.Write($"{threadInfo.MemoryUsageMB} MB");
        Console.ResetColor();
        Console.Write($" | CPU: ");
        Console.ForegroundColor = ConsoleColor.Magenta;
        Console.Write($"{threadInfo.CpuTime.TotalSeconds:F1}s");
        Console.ResetColor();
        Console.WriteLine();
        
        // Linha 2: Informações detalhadas de threads
        Console.Write("DISTRIBUICAO: ");
        
        Console.ForegroundColor = ConsoleColor.Green;
        Console.Write($"Ativas: {threadInfo.ActiveThreads}");
        Console.ResetColor();
        
        Console.Write(" | ");
        Console.ForegroundColor = ConsoleColor.Yellow;
        Console.Write($"Executando: {threadInfo.RunningThreads}");
        Console.ResetColor();
        
        Console.Write(" | ");
        Console.ForegroundColor = ConsoleColor.Red;
        Console.Write($"Aguardando: {threadInfo.WaitingThreads}");
        Console.ResetColor();
        
        Console.Write(" | ");
        Console.ForegroundColor = ConsoleColor.Gray;
        Console.Write($"Inativas: {threadInfo.InactiveThreads}");
        Console.ResetColor();
        
        Console.WriteLine();
        
        // Linha 3: Barra visual de proporção de threads com total
        Console.Write($"VISUALIZACAO ({threadInfo.TotalThreads} threads): ");
        DrawThreadBar(threadInfo);
        
        Console.WriteLine();
    }

    private void DrawThreadBar(ThreadMonitor.ThreadInfo threadInfo)
    {
        const int BAR_WIDTH = 30; // Reduzido para dar espaço ao texto adicional
        
        Console.Write("[");
        
        if (threadInfo.TotalThreads > 0)
        {
            int activeBlocks = Math.Max(1, (threadInfo.ActiveThreads * BAR_WIDTH) / threadInfo.TotalThreads);
            int runningBlocks = Math.Max(0, (threadInfo.RunningThreads * BAR_WIDTH) / threadInfo.TotalThreads);
            int waitingBlocks = Math.Max(0, (threadInfo.WaitingThreads * BAR_WIDTH) / threadInfo.TotalThreads);
            int inactiveBlocks = BAR_WIDTH - activeBlocks - waitingBlocks;
            
            // Ajusta para não exceder o tamanho da barra
            if (activeBlocks + waitingBlocks + inactiveBlocks > BAR_WIDTH)
            {
                inactiveBlocks = BAR_WIDTH - activeBlocks - waitingBlocks;
            }
            
            // Desenha blocos para threads executando (verde brilhante)
            Console.ForegroundColor = ConsoleColor.Green;
            Console.Write(new string('#', runningBlocks));
            
            // Desenha blocos para threads ativas mas não executando (verde escuro)
            Console.ForegroundColor = ConsoleColor.DarkGreen;
            Console.Write(new string('#', Math.Max(0, activeBlocks - runningBlocks)));
            
            // Desenha blocos para threads aguardando (vermelho)
            Console.ForegroundColor = ConsoleColor.Red;
            Console.Write(new string('#', waitingBlocks));
            
            // Desenha blocos para threads inativas (cinza)
            Console.ForegroundColor = ConsoleColor.DarkGray;
            Console.Write(new string('#', Math.Max(0, inactiveBlocks)));
        }
        else
        {
            Console.ForegroundColor = ConsoleColor.DarkGray;
            Console.Write(new string('#', BAR_WIDTH));
        }
        
        Console.ResetColor();
        Console.Write("]");
        
        // Adiciona percentuais e contagem total
        if (threadInfo.TotalThreads > 0)
        {
            double activePercent = (double)threadInfo.ActiveThreads / threadInfo.TotalThreads * 100;
            double inactivePercent = (double)threadInfo.InactiveThreads / threadInfo.TotalThreads * 100;
            
            Console.Write($" {activePercent:F0}%^ {inactivePercent:F0}%v");
        }
    }

    private string GetCellString(int col, int row)
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
                return BLOCK_STRING;
            }
        }

        // Caso contrário, obtém a célula do tabuleiro
        int cell = _board.GetCell(new Point(col, row));
        return cell != 0 ? BLOCK_STRING : EMPTY_STRING;
    }

    private string GetSidePanelText(int row)
    {
        if (row == 0)
            return "|  PONTOS:" + _gameState.Score.ToString().PadLeft(8) + " |";
        else if (row == 1)
            return "|  LINHAS:" + _gameState.Lines.ToString().PadLeft(8) + " |";
        else if (row == 2)
            return "|  NIVEL: " + _gameState.Level.ToString().PadLeft(8) + " |";
        else if (row == 4)
            return "|   PROXIMA PECA   |";
        else if (row >= 5 && row < 5 + GameConfig.NextPiecePreviewSize && _gameState.NextPiece != null)
        {
            return "|     " + DrawNextPiecePreview(row - 5) + "     |";
        }
        else if (row == 12)
            return "|    CONTROLES:    |";
        else if (row == 13)
            return "|<- -> / A D: Mover|";
        else if (row == 14)
            return "|  ^ / W: Girar    |";
        else if (row == 15)
            return "|  v / S: Descer   |";
        else if (row == 16)
            return "|  Espaco: Queda   |";
        else if (row == 17)
            return "|  P: Pausar       |";
        else if (row == 18)
            return "|  ESC/Q: Sair     |";
        else
            return "|" + new string(' ', GameConfig.SidePanelWidth - 2) + "|";
    }

    private string DrawNextPiecePreview(int row)
    {
        const int PREVIEW_WIDTH = 8; // 4 blocos * 2 chars
        
        if (_gameState.NextPiece == null || row >= _gameState.NextPiece.Shape.GetLength(0))
        {
            return new string(' ', PREVIEW_WIDTH);
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
            return new string(' ', PREVIEW_WIDTH);
        }
        
        pieceWidth = (lastBlock - firstBlock + 1) * 2; // 2 caracteres por bloco
        int leftPadding = (PREVIEW_WIDTH - pieceWidth) / 2;
        
        // Constrói a string da peça
        string pieceString = "";
        for (int col = firstBlock; col <= lastBlock; col++)
        {
            pieceString += _gameState.NextPiece.Shape[row, col] > 0 ? BLOCK_STRING : EMPTY_STRING;
        }
        
        int rightPadding = PREVIEW_WIDTH - leftPadding - pieceWidth;
        return new string(' ', leftPadding) + pieceString + new string(' ', rightPadding);
    }
} 