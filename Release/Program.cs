using Base;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Release
{
    class Program
    {
        static void Main(string[] args)
        {
            if (args.Length != 1)
            {
                Console.Error.WriteLine("Invalid usage. Expected: release.exe <build_path>");
                Environment.Exit(1);
            }

            var buildPath = args[0];

            Console.Out.WriteLine(Utils.GenerateReleaseInfo(buildPath));
        }
    }
}
