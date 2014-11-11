using System;
using NSpec;
using System.IO;
using System.Diagnostics;

namespace Tailor.Tests
{
    class CommandOptionParsingSpec : nspec
    {
        string appDir, outputDroplet, outputMetadata;
        Process process;

        void before_each()
        {
            appDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(appDir);
            outputDroplet = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tgz");
            outputMetadata = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".json");

            process = Helpers.StartTailor(appDir, outputDroplet, outputMetadata);
            process.WaitForExit();
        }

        void after_each()
        {
            Directory.Delete(appDir, true);
            if(File.Exists(outputDroplet)) File.Delete(outputDroplet);
            if (File.Exists(outputMetadata)) File.Delete(outputMetadata);
        }

        void describe_command_option_parsing()
        {
            it["outputs outputMetadata directory"] = () =>
            {
                var stdout = process.StandardOutput.ReadToEnd();
                stdout.should_contain("OutputMetadata: " + outputMetadata);
            };
        }
    }
}



