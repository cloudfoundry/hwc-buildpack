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
        private string containerDir;

        void before_each()
        {
            options = new Tailor.Options
            {
                AppDir = "/app",
                OutputDroplet = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tgz"),
                OutputMetadata = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".json")
            };

            containerDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(containerDir);

            Tailor.Program.Run(options, containerDir);
        }

        void after_each()
        {
            Directory.Delete(options.AppDir, true);
            File.Delete(options.OutputDroplet);
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
    }
}



