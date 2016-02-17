using Base;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Compile
{
    class Program
    {
        static void Main(string[] args)
        {
            if (args.Length != 2)
            {
                Console.Error.WriteLine("Invalid usage. Expected: compile.exe <build_path> <cache_path>");
                Environment.Exit(1);
            }

            var buildPath = args[0];
            var cachePath = args[1];

            var binDirectory = new FileInfo(System.Reflection.Assembly.GetEntryAssembly().Location).Directory;

            var files = Directory.EnumerateFiles(buildPath).ToList();
            if (Utils.HasWebConfig(files))
            {
                Directory.CreateDirectory(Path.Combine(buildPath, "WebAppServer", "bin"));
                Utils.CopyDirectory(Path.Combine(binDirectory.Parent.FullName, "WebAppServer", "bin"), Path.Combine(buildPath, "WebAppServer"));
                Utils.SetConnectionStrings(files);
            }
            if (Utils.ExeFiles(files).Count() == 1)
            {
                // Nothing to do.
            }

            Environment.Exit(0);
        }
    }
}
