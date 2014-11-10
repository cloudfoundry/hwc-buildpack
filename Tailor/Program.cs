using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Tailor
{
    public class Program
    {
        public static void Run(Options options)
        {
            // Values are available here
            Console.WriteLine("OutputMetadata: {0}", options.OutputMetadata);

            JObject execution_metadata = new JObject();
            execution_metadata["start_command"] = "the start command";
            JObject detected_start_command = new JObject();
            detected_start_command["web"] = "the start command";

            JObject obj = new JObject();
            obj["execution_metadata"] = execution_metadata.ToString(Formatting.None);
            obj["detected_start_command"] = detected_start_command;
            System.IO.File.WriteAllText(options.OutputMetadata, obj.ToString());
        }

        static void Main(string[] args)
        {
            var options = new Options();
            if (!CommandLine.Parser.Default.ParseArguments(args, options))
            {
                Environment.Exit(1);
            }

            Run(options);
        }
    }
}
