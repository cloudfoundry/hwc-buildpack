using System;
using NSpec;
using System.IO;
using System.Diagnostics;

namespace Tailor.Tests
{
    class CommandOptionParsingSpec : nspec
    {
        string appDir, outputDroplet, outputMetadata;
        StreamReader stdout;

        void before_each()
        {
            appDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(appDir);
            outputDroplet = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tgz");
            outputMetadata = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".json");

            stdout = StartTailor(appDir, outputDroplet, outputMetadata);
        }

        void after_each()
        {
            Directory.Delete(appDir, true);
            if(File.Exists(outputDroplet)) File.Delete(outputDroplet);
            if (File.Exists(outputMetadata)) File.Delete(outputMetadata);
        }

        void describe_command_option_parsing()
        {
            it["uses outputDroplet"] = () =>
            {
                File.Exists(outputDroplet).should_be_true();
            };

            it["uses outputMetadata"] = () =>
            {
                File.Exists(outputMetadata).should_be_true();
            };
        }

        public static StreamReader StartTailor(string appDir, string outputDroplet, string outputMetadata)
        {
            var process = new Process();
            process.StartInfo.FileName = @"..\..\..\Tailor\bin\Debug\Tailor.exe";
            process.StartInfo.Arguments = "-appDir=\"" + appDir + "\" -outputDroplet=\"" + outputDroplet + "\" -outputMetadata=\"" + outputMetadata + "\"";
            process.StartInfo.UseShellExecute = false;
            process.StartInfo.RedirectStandardOutput = true;
            process.Start();
            process.WaitForExit();
            return process.StandardOutput;
        }
    }
}



