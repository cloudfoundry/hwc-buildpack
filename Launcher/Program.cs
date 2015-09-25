using System;
using System.Diagnostics;
using System.Threading;
using Newtonsoft.Json;

namespace Launcher
{
    public class ExecutionMetadata
    {
        [JsonProperty("start_command")]
        public string StartCommand { get; set; }

        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs { get; set; }
    }

    internal class Program
    {
        private static int Main(string[] args)
        {
            if (Environment.GetEnvironmentVariable("ARGJSON") != null && Environment.GetEnvironmentVariable("ARGJSON").Length >= 2)
                args = JsonConvert.DeserializeObject<string[]>(Environment.GetEnvironmentVariable("ARGJSON"));

            if (args.Length < 3)
            {
                Console.Error.WriteLine("Launcher was run with insufficient arguments. The arguments were: {0}",
                    String.Join(" ", args));
                return 1;
            }


            var startInfo = new ProcessStartInfo
            {
                FileName = args[1],
                UseShellExecute = false,
            };

            if (String.IsNullOrWhiteSpace(startInfo.FileName))
            {
                Console.Error.WriteLine("Could not determine a start command. Use the -c flag to 'cf push' to specify a custom start command.");
                return 1;
            }
            Console.Out.WriteLine("Run {0} :with: {1}", startInfo.FileName, startInfo.Arguments);

            var process = Process.Start(startInfo);

            process.WaitForExit();

            return process.ExitCode;
        }
    }
}
