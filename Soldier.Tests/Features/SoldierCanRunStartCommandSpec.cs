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

namespace Soldier.Tests.Features
{
    class SoldierCanRunStartCommandSpec : nspec
    {
        private void describe_()
        {
            ProcessStartInfo normalStartInfo = null;
            ProcessStartInfo explodingStartInfo = null;
            Process process = null;

            describe["Given that a start command is provided to the Soldier"] = () =>
            {
                before = () =>
                {
                    normalStartInfo = new ProcessStartInfo
                    {
                        UseShellExecute = false,
                        FileName = "Soldier.exe",
                        Arguments = ArgumentEscaper.Escape(
                            new[]
                            {
                                "/app",
                                "",
                                "{\"start_command\":\"Fixtures\\CivetCat.bat\", \"start_command_args\":[\"bean1\", \"bean\\2\"]}"
                            }
                        ),
                        RedirectStandardOutput = true,
                        RedirectStandardError = true,
                    };

                    explodingStartInfo = new ProcessStartInfo
                    {
                        UseShellExecute = false,
                        FileName = "Soldier.exe",
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
                        stderr.should_contain("Soldier was run with insufficient arguments");
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
                        stderr.should_contain("Soldier was run with invalid JSON for the metadata argument");
                    };

                    it["returns an exit code of 1"] = () =>
                    {
                        process.ExitCode.should_be(1);
                    };
                };
            };          
        }
    }
}
