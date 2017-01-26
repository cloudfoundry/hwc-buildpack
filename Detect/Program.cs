using Base;
using System;
using System.IO;
using System.Linq;

namespace Detect
{
    class Program
    {
        static void Main(string[] args)
        {
            if (args.Length != 1)
            {
                Console.Error.WriteLine("Invalid usage. Expected: detect.exe <build_path>");
                Environment.Exit(1);
            }

            var buildPath = args[0];

            var files = Directory.EnumerateFiles(buildPath).ToList();
            if (Utils.HasWebConfig(files))
            {
                Console.Out.Write("WebAppServer");
                Environment.Exit(0);
            }
            if (Utils.ExeFiles(files).Count() == 1)
            {
                Console.Out.Write("Exe");
                Environment.Exit(0);
            }

            Environment.Exit(1);
        }
    }
}
