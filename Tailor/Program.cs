using System.ComponentModel;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using SharpCompress.Common;
using SharpCompress.Writer;
using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tailor
{
    public class Program
    {
        public static void Run(Options options)
        {
            var appPath = System.IO.Directory.GetCurrentDirectory() + options.BuildDir;
            var outputDropletPath = System.IO.Directory.GetCurrentDirectory() + options.OutputDroplet;
            TarGZFile.CreateFromDirectory(appPath, outputDropletPath);

            // Result.JSON
            GenerateOutputMetadata(options.OutputMetadata);
        }

        private static void GenerateOutputMetadata(string fileName)
        {
            JObject execution_metadata = new JObject();
            execution_metadata["start_command"] = "the start command";
            JObject detected_start_command = new JObject();
            detected_start_command["web"] = "the start command";

            JObject obj = new JObject();
            obj["execution_metadata"] = execution_metadata.ToString(Formatting.None);
            obj["detected_start_command"] = detected_start_command;
            System.IO.File.WriteAllText(System.IO.Directory.GetCurrentDirectory() + fileName, obj.ToString());
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
