using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using NSpec;

namespace Tailor.Tests.Specs.Features
{
    internal class StagerCanRunTailorSpec : nspec
    {
        private string arguments;

        private void describe_()
        {
            context["Given That I am a CC Bridge Stager"] = () =>
            {
                before = () =>
                {
                    var filename = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "UnitTestProject2", "app.zip");
                    File.Delete(filename);

                    arguments = new Dictionary<string, string>
                    {
                        {"-appDir", "/app"},
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
                        var process = new Process
                        {
                            StartInfo =
                            {
                                FileName = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "Tailor", "bin", "debug", "Tailor.exe"),
                                Arguments = arguments
                            }
                        };

                        process.Start();
                        process.WaitForExit();
                    };

                    it["Creates a droplet"] = () => { File.Exists("./tmp/droplet").should_be_true(); };
                };

                after = () =>
                {
                };
            };
        }
    }
}