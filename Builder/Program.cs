using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace Builder
{

    public class ExecutionMetadata
    {
        public ExecutionMetadata()
        {
            StartCommand = "";
            StartCommandArgs = new string[]{};
        }

        [JsonProperty("start_command")]
        public string StartCommand
        {
            get;
            set;
        }

        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs
        {
            get;
            set;
        }
    }

    public class DetectedStartCommand
    {
        [JsonProperty("web")]
        public string Web { get; set; }
    }

    public class OutputMetadata
    {
        public ExecutionMetadata ExecutionMetadata { get; set; }

        [JsonProperty("execution_metadata")]
        public string execution_metadata
        {
            get { return JsonConvert.SerializeObject(ExecutionMetadata); }
        }

        [JsonProperty("detected_start_command")]
        public DetectedStartCommand DetectedStartCommand { get; set; }

        public OutputMetadata()
        {
            ExecutionMetadata = new ExecutionMetadata();
            DetectedStartCommand = new DetectedStartCommand();
        }

        public OutputMetadata(IList<string> files) : this()
        {
            if (files.Any((x) => Path.GetFileName(x).ToLower() == "web.config"))
            {
                ExecutionMetadata.StartCommand = "tmp/lifecycle/WebAppServer.exe";
                ExecutionMetadata.StartCommandArgs = new string[] {"."};
            }
            else {
                var executables = files.Where((x) => x.EndsWith(".exe")).ToList();
                if (executables.Any())
                {
                    if (executables.Count() > 1) throw new Exception("Directory contained more than 1 executable file.");
                    ExecutionMetadata.StartCommand = Path.GetFileName(executables.First());
                }
                else
                {
                    throw new Exception("No runnable application found.");
                }
            }
            DetectedStartCommand.Web = ExecutionMetadata.StartCommand;
            if (ExecutionMetadata.StartCommandArgs.Any())
            {
                DetectedStartCommand.Web += " " + String.Join(" ", ExecutionMetadata.StartCommandArgs);
            }
        }
    }

    public class Program
    {
        public static void Run(Options options)
        {
            var appPath = Directory.GetCurrentDirectory() + options.BuildDir;
            var outputDropletPath = Directory.GetCurrentDirectory() + options.OutputDroplet;
            TarGZFile.CreateFromDirectory(appPath, outputDropletPath);

            // Result.JSON
            GenerateOutputMetadata(appPath, options.OutputMetadata);
        }

        private static void GenerateOutputMetadata(string appPath, string fileName)
        {
            var obj = new OutputMetadata(Directory.EnumerateFiles(appPath).ToList());
            File.WriteAllText(Directory.GetCurrentDirectory() + fileName, JsonConvert.SerializeObject(obj));
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
