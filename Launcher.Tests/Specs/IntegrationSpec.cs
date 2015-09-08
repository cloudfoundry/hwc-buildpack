using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Reflection;
using Newtonsoft.Json;
using NSpec;

namespace Launcher.Tests.Specs
{
    class IntegrationSpec : nspec
    {
        public static string AssemblyDirectory
        {
            get
            {
                string codeBase = Assembly.GetExecutingAssembly().CodeBase;
                UriBuilder uri = new UriBuilder(codeBase);
                string path = Uri.UnescapeDataString(uri.Path);
                return Path.GetDirectoryName(path);
            }
        }

        public void describe_()
        {
            describe["an app started with no start command"] = () =>
            {
                it["prints an error message"] = () =>
                {
                    var startInfo = new ProcessStartInfo(AssemblyDirectory + @"\..\..\..\Launcher\bin\Launcher.exe")
                    {
                        UseShellExecute = false,
                        RedirectStandardError = true,
                        RedirectStandardOutput = true,

                    };
                    var args = new string[]
                    {
                        "", // ignored
                        "", // optional filename
                        JsonConvert.SerializeObject(new ExecutionMetadata()
                        {
                            StartCommand = "",
                            StartCommandArgs = new string[] {"foo"},
                        })
                    };
                    startInfo.EnvironmentVariables["ARGJSON"] = JsonConvert.SerializeObject(args);
                    var process = Process.Start(startInfo);
                    process.WaitForExit(1000);
                    var stderr = process.StandardError.ReadToEnd();
                    stderr.should_contain("Could not determine a start command");
                    process.ExitCode.should_be(1);
                };
            };
        }
    }
}
