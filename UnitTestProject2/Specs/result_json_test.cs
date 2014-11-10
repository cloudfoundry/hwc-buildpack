using System;
using NSpec;
using System.IO;
using System.Diagnostics;
using System.Linq;
using Newtonsoft.Json.Linq;
using Tailor;

namespace Tailor.Tests
{
    class TheResultJsonSpec : nspec
    {
        Tailor.Options options; 
        Process process;

        void before_each()
        {
            options = new Tailor.Options
            {
                AppDir = GetTemporaryDirectory(),
                OutputDroplet = GetTemporaryDirectory(),
                OutputMetadata = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".json")
            };

            Tailor.Program.Run(options);
        }

        void after_each()
        {
            Directory.Delete(options.AppDir, true);
            Directory.Delete(options.OutputDroplet, true);
            File.Delete(options.OutputMetadata);
        }

        void describe_result_json()
        {
            it["exists, and contains the start command"] = () =>
            {
                string text = File.ReadAllText(options.OutputMetadata);
                JObject obj = JObject.Parse(text);

                ((string)obj.SelectToken("$.execution_metadata")).should_be("{\"start_command\":\"the start command\"}");
                ((string)obj.SelectToken("$.detected_start_command.web")).should_be("the start command");
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



