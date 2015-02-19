using System.Data.Odbc;
using System.Diagnostics;
using System.IO;
using Newtonsoft.Json;
using NSpec;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Launcher.Tests.Features
{
    class LauncherCanRunStartCommandSpec : nspec
    {
        private void describe_()
        {
            string[] normalStartArguments = null;
            ProcessStartInfo normalStartInfo = null;
            ProcessStartInfo explodingStartInfo = null;
            Process process = null;

            before = () =>
            {
                var workingDirectory = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Launcher.Tests", "bin", "Debug");
                Directory.SetCurrentDirectory(workingDirectory);
            };

            after = () =>
            {
                File.Delete("Bean.txt");
            };

            describe["Given that a start command is provided to the Launcher"] = () =>
            {
                before = () =>
                {
                    normalStartArguments = new[]
                    {
                        "/app",
                        "",
                        "{\"start_command\":\"Fixtures\\CivetCat.bat\", \"start_command_args\":[\"bean1\", \"bean\\2\"]}"
                    };
                    normalStartInfo = new ProcessStartInfo
                    {
                        UseShellExecute = false,
                        FileName = "Launcher.exe",
                        Arguments = ArgumentEscaper.Escape(normalStartArguments),
                        RedirectStandardOutput = true,
                        RedirectStandardError = true,
                    };

                    explodingStartInfo = new ProcessStartInfo
                    {
                        UseShellExecute = false,
                        FileName = "Launcher.exe",
                        Arguments = ArgumentEscaper.Escape(
                            new[]
                            {
                               "/app",
                               "",
                               "{\"start_command\":\"Fixtures\\Explosions.bat\", \"start_command_args\":[\"boom\"]}"
                            }
                        ),
                        RedirectStandardOutput = true,
                        RedirectStandardError = true,
                    };
                };

                act = () =>
                {
                    process = Process.Start(normalStartInfo);
                    process.WaitForExit();
                };

                describe["as command line arguments"] = () =>
                {
                    it["runs the start command with the correct arguments"] = () =>
                    {
                        var beans = File.ReadAllText("Bean.txt");
                        beans.should_contain("\"bean1\" \"bean\\\\2\"");
                    };

                    it["returns the exit code from the start command"] = () =>
                    {
                        process.ExitCode.should_be(0);

                        process = Process.Start(explodingStartInfo);
                        process.WaitForExit();
                        process.ExitCode.should_be(1);
                    };

                    it["propagates stdout from the start command"] = () =>
                    {
                        var stdout = process.StandardOutput.ReadToEnd();
                        stdout.should_contain("This is STDOUT");
                    };

                    it["propagates stderr from the start command"] = () =>
                    {
                        var stderr = process.StandardError.ReadToEnd();
                        stderr.should_contain("This is STDERR");
                    };

                    describe["When the soldier is provided with insufficient arguments"] = () =>
                    {
                        before = () =>
                        {
                            normalStartInfo.Arguments = "IamNotEnough";
                        };

                        it["outputs a message onto STDERR"] = () =>
                        {
                            var stderr = process.StandardError.ReadToEnd();
                            stderr.should_contain("Launcher was run with insufficient arguments");
                        };

                        it["returns an exit code of 1"] = () =>
                        {
                            process.ExitCode.should_be(1);
                        };
                    };

                    describe["When the soldier is provided with invalid json for the metadata argument"] = () =>
                    {
                        before = () =>
                        {
                            normalStartInfo.Arguments = "\"/app\" \"\" \"{I am bad JSON\"";
                        };

                        it["outputs a message onto STDERR"] = () =>
                        {
                            var stderr = process.StandardError.ReadToEnd();
                            stderr.should_contain("Launcher was run with invalid JSON for the metadata argument");
                        };

                        it["returns an exit code of 1"] = () =>
                        {
                            process.ExitCode.should_be(1);
                        };
                    };
                };

                describe["as ENV[ARGJSON]"] = () =>
                {
                    before = () =>
                    {
                        normalStartArguments[2] = "{\"start_command\":\"Fixtures/CivetCat.bat\", \"start_command_args\":[\"bean1\", \"bean\\\\2\"]}";
                        normalStartInfo.EnvironmentVariables["ARGJSON"] = JsonConvert.SerializeObject(normalStartArguments);
                        normalStartInfo.Arguments = "args are overrriden";
                    };

                    it["runs the start command with the correct arguments"] = () =>
                    {
                        var beans = File.ReadAllText("Bean.txt");
                        beans.should_contain("\"bean1\" \"bean\\\\2\"");
                    };

                    it["returns the exit code from the start command"] = () =>
                    {
                        process.ExitCode.should_be(0);
                    };
                };
            };          
        }
    }
}
