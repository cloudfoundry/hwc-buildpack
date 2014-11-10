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
            appDir = GetTemporaryDirectory();
            outputDroplet = GetTemporaryDirectory();
            outputMetadata = GetTemporaryDirectory();

            process = Helpers.StartTailor(appDir, outputDroplet, outputMetadata); 
        }

        void after_each()
        {
            process.WaitForExit();
            Directory.Delete(appDir, true);
            Directory.Delete(outputDroplet, true);
            Directory.Delete(outputMetadata, true);
        }

        void describe_command_option_parsing()
        {
            it["outputs outputMetadata directory"] = () =>
            {
                var stdout = process.StandardOutput.ReadToEnd();
                stdout.should_contain("OutputMetadata: " + outputMetadata);
            };
        }


        public string GetTemporaryDirectory()
        {
            string tempDirectory = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(tempDirectory);
            return tempDirectory;
        }
    }
}



