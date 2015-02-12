using Microsoft.Web.Administration;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Data;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices;
using System.Security;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Newtonsoft.Json.Linq;

namespace Launcher
{
    public class ExecutionMetadata
    {
        [JsonProperty("start_command")]
        public string DetectedStartCommand { get; set; }

        [JsonProperty("start_command_args")]
        public string[] StartCommandArgs { get; set; }
    }

    class Program
    {
        static int Main(string[] args)
        {
            if (args.Length < 3)
            {
                Console.Error.WriteLine("Launcher was run with insufficient arguments. The arguments were: {0}",
                    String.Join(" ", args));
                return 1;
            }

            ExecutionMetadata executionMetadata = null;

            try
            {
                executionMetadata = JsonConvert.DeserializeObject<ExecutionMetadata>(args[2]);
            }
            catch (Exception ex)
            {
                Console.Error.WriteLine(
                    "Launcher was run with invalid JSON for the metadata argument. The JSON was: {0}", args[2]);
                return 1;
            }

            var startInfo = new ProcessStartInfo
            {
                UseShellExecute = false,
                FileName = executionMetadata.DetectedStartCommand,
                Arguments = ArgumentEscaper.Escape(executionMetadata.StartCommandArgs),
            };

            var process = Process.Start(startInfo);

            process.WaitForExit();

            return process.ExitCode;
        }
    }
}
