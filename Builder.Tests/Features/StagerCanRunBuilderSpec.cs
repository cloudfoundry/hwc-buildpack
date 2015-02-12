using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using Newtonsoft.Json.Linq;
using NSpec;

namespace Builder.Tests.Specs.Features
{
    internal class StagerCanRunBuilderSpec : nspec
    {
        private string arguments;

        private void describe_()
        {
            context["Given That I am a CC Bridge Stager"] = () =>
            {
                before = () =>
                {
                    var filename = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests", "tmp", "droplet");
                    File.Delete(filename);

                    arguments = new Dictionary<string, string>
                    {
                        {"-buildDir", "/app"},
                        {"-outputDroplet", "/tmp/droplet"},
                        {"-outputMetadata", "/tmp/result.json"},
                        {"-buildArtifactsCacheDir", "/tmp/cache"},
                        {"-buildpackOrder", "buildpackGuid1,buildpackGuid2"},
                        {"-outputBuildArtifactsCache", "/tmp/output-cache"},
                        {"-skipCertVerify", "false"}
                    }
                        .Select(x => x.Key + " " + x.Value)
                        .Aggregate((x, y) => x + " " + y);
                };

                context["When I invoke the tailor"] = () =>
                {
                    before = () =>
                    {
                        var workingDir = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests");
                        var process = new Process
                        {
                            StartInfo =
                            {
                                FileName = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder", "bin", "debug", "Builder.exe"),
                                Arguments = arguments,
                                WorkingDirectory = workingDir
                            }
                        };

                        process.Start();
                        process.WaitForExit();
                    };

                    it["Creates a droplet"] = () =>
                    {

                        var fileName = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests", "tmp", "droplet");
                        File.Exists(fileName).should_be_true();
                    };

                    it["Creates the result.json"] = () =>
                    {
                        var resultFile = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests", "tmp", "result.json");
                        File.Exists(resultFile).should_be_true();

                        JObject result = JObject.Parse(File.ReadAllText(resultFile));
                        var execution_metadata = JObject.Parse(result["execution_metadata"].Value<string>());
                        execution_metadata["start_command"].Value<string>().should_be("tmp/lifecycle/WebAppServer.exe");
                        execution_metadata["start_command_args"].Values<string>().should_be(new [] {"."});
                    };

                    it["includes magical json properties required for the diego lifecyle (in cf push) to work"] = () =>
                    {
                        var resultFile = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests", "tmp", "result.json");
                        File.Exists(resultFile).should_be_true();

                        JObject result = JObject.Parse(File.ReadAllText(resultFile));
                        result["detected_start_command"].should_not_be_null();
                        result["detected_start_command"]["web"].should_not_be_null();

                    };
                };

                after = () =>
                {
                    var filename = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Builder.Tests", "tmp", "droplet");
                    File.Delete(filename);
                };
            };
        }
    }
}