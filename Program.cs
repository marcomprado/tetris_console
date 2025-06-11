using TetrisConsole.Engine;

namespace TetrisConsole;

class Program
{
    static void Main(string[] args)
    {
        Console.Clear();
        var engine = new GameEngine();
        engine.Start();
    }
}
