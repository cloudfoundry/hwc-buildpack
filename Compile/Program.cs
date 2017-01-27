using Base;
using System;
using System.IO;
using System.Linq;

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
                var webAppServerDestination = Path.Combine(buildPath, ".cloudfoundry", "WebAppServer");
                Directory.CreateDirectory(webAppServerDestination);
                Utils.CopyDirectory(Path.Combine(binDirectory.Parent.FullName, "WebAppServer", "bin"), webAppServerDestination);
            }

            Environment.Exit(0);
        }
    }
}
