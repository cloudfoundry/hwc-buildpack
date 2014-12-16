using System;
using System.Diagnostics;
using System.IO;
using NSpec;

namespace Tailor.Tests
{
    internal class CommandOptionParsingSpec : nspec
    {
        private string appDir, outputDroplet, outputMetadata;
        private StreamReader stdout;

        private void before_each()
        {
            var containerPath = Path.GetTempPath();
            appDir = @"/app";
            Directory.CreateDirectory(appDir);
            outputDroplet = @"/tmp/droplet";
            outputMetadata = @"/tmp/result.json";
            string buildArtifactsCacheDir = @"/tmp/cache";
            string buildpackOrder = @"buildpackGuid1,buildpackGuid2";
            string buildpacksDir = @"/tmp/buildpacks";
            string outputBuildArtifactsCache = @"/tmp/output-cache";
            string skipCertVerify = "false";
            stdout = StartTailor(
                appDir: appDir,
                buildArtifactsCacheDir: buildArtifactsCacheDir,
                buildpackOrder: buildpackOrder,
                outputBuildArtifactsCache: outputBuildArtifactsCache,
                buildpacksDir: buildpacksDir,
                outputDroplet: outputDroplet,
                outputMetadata: outputMetadata,
                skipCertVerify: skipCertVerify);
        }

        private void after_each()
        {
            Directory.Delete(appDir, true);
            if (File.Exists(outputDroplet)) File.Delete(outputDroplet);
            if (File.Exists(outputMetadata)) File.Delete(outputMetadata);
        }

        private void describe_command_option_parsing()
        {
            it["uses outputDroplet"] = () => { File.Exists("." + outputDroplet).should_be_true(); };

            it["uses outputMetadata"] = () => { File.Exists("." + outputMetadata).should_be_true(); };
        }

        public static StreamReader StartTailor(
            string appDir,
            string buildArtifactsCacheDir,
            string buildpackOrder,
            string outputBuildArtifactsCache,
            string buildpacksDir,
            string outputDroplet,
            string outputMetadata,
            string skipCertVerify)
        {
            var process = new Process();
            process.StartInfo.FileName = @"..\..\..\Tailor\bin\Debug\Tailor.exe";
            process.StartInfo.Arguments =
                String.Format(
                    @"-appDir={0} -buildArtifactsCacheDir={1} -buildpackOrder={2} -buildpacksDir={3} -outputBuildArtifactsCache={4} -outputDroplet={5} -outputMetadata={6} -skipCertVerify={7}",
                    appDir,
                    buildArtifactsCacheDir,
                    buildpackOrder,
                    buildpacksDir,
                    outputBuildArtifactsCache,
                    outputDroplet,
                    outputMetadata,
                    skipCertVerify
                    );

            process.StartInfo.UseShellExecute = false;
            process.StartInfo.RedirectStandardOutput = true;
            process.Start();
            process.WaitForExit();
            return process.StandardOutput;
        }
    }
}