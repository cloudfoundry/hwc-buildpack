using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace Builder
{
    public class Program
    {
        public static ExecutionMetadata GenerateExecutionMetadata(IList<string> files)
        {
            var executionMetadata = new ExecutionMetadata();
            var procfiles = files.Where(x => Path.GetFileName(x).ToLower() == "procfile").ToList();
            var executables = files.Where(x => x.EndsWith(".exe")).ToList();
            if (procfiles.Any())
            {
                var file = File.ReadAllLines(procfiles.First());
                var webline = file.Where(x => x.StartsWith("web:"));
                if (webline.Any())
                {
                    var contents = webline.First().Substring(4).Trim().Split(new[] { ' ' });
                    executionMetadata.StartCommand = contents[0];
                    executionMetadata.StartCommandArgs = contents.Skip(1).ToArray();
                }
                else
                {
                    throw new Exception("Procfile didn't contain a web line");
                }
            }
            else if (files.Any(x => Path.GetFileName(x).ToLower() == "web.config"))
            {
                executionMetadata.StartCommand = "tmp/lifecycle/WebAppServer.exe";
                executionMetadata.StartCommandArgs = new[] { "." };
            }
            else if (executables.Any())
            {
                if (executables.Count() > 1)
                    throw new Exception("Directory contained more than 1 executable file.");
                executionMetadata.StartCommand = Path.GetFileName(executables.First());
            }
            else
            {
                throw new Exception("No runnable application found.");
            }

            return executionMetadata;
        }

        static void Main(string[] args)
        {
            SanitizeArgs(args);
            var options = new Options();
            if (!CommandLine.Parser.Default.ParseArguments(args, options))
            {
                Environment.Exit(1);
            }

            Run(options);
        }

        private static void Run(Options options)
        {
            var appPath = Directory.GetCurrentDirectory() + options.BuildDir;
            var outputDropletPath = Directory.GetCurrentDirectory() + options.OutputDroplet;
            TarGZFile.CreateFromDirectory(appPath, outputDropletPath);

            // Result.JSON
            var obj = GenerateOutputMetadata(appPath);
            File.WriteAllText(Path.Combine(Directory.GetCurrentDirectory(), options.OutputMetadata), JsonConvert.SerializeObject(obj));
        }

        private static OutputMetadata GenerateOutputMetadata(string appPath)
        {
            var files = Directory.EnumerateFiles(appPath).ToList();
            var executionMetadata = GenerateExecutionMetadata(files);
            return new OutputMetadata()
            {
                ExecutionMetadata = executionMetadata,
            };
        }

        private static void SanitizeArgs(string[] args)
        {
            for (int i = 0; i < args.Length; i++)
            {
                if (args[i].StartsWith("-") && !args[i].StartsWith("--"))
                {
                    args[i] = "-" + args[i];
                }
            }
        }
    }
}
