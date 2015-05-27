using System;
using System.Diagnostics;
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
                UseShellExecute = false,
            };

            if (args[1] != "")
            {
                startInfo.FileName = args[1];
            }
            else
            {
                try
                {
                    var executionMetadata = JsonConvert.DeserializeObject<ExecutionMetadata>(args[2]);
                    startInfo.FileName = executionMetadata.StartCommand;
                    startInfo.Arguments = ArgumentEscaper.Escape(executionMetadata.StartCommandArgs);
                }
                catch (Exception)
                {
                    Console.Error.WriteLine(
                        "Launcher was run with invalid JSON for the metadata argument. The JSON was: {0}", args[2]);
                    return 1;
                }
            }

            Console.Out.WriteLine("Run {0} :with: {1}", startInfo.FileName, startInfo.Arguments);

            var process = Process.Start(startInfo);

            process.WaitForExit();

            return process.ExitCode;
        }
    }
}
