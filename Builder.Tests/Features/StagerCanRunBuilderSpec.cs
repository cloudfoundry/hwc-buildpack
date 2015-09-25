using System;
using Builder.Models;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;
using NSpec;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Xml;

namespace Builder.Tests.Specs.Features
{
    class StagerCanRunBuilderSpec : nspec
    {
        private string arguments;

        private void describe_()
        {
            Process process = null;
            string currentDirectory = null;
            string workingDirectory = null;
            string appDir = null;
            string tmpDir = null;

            act = () =>
            {
                process.StartInfo.WorkingDirectory = workingDirectory;
                process.Start();
                process.WaitForExit();
                process.ExitCode.should_be(0);
            };

            before = () =>
            {
                workingDirectory = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
                tmpDir = Path.Combine(workingDirectory, "tmp");
                appDir = Path.Combine(workingDirectory, "app");
                Directory.CreateDirectory(tmpDir);
                Directory.CreateDirectory(appDir);

                currentDirectory =
                    Path.GetFullPath(
                        Path.Combine(System.Reflection.Assembly.GetExecutingAssembly().CodeBase, "..", "..", "..",
                            "..").Replace("file:///", ""));
                var builderBinary = Path.Combine(currentDirectory, "Builder", "bin", "Builder.exe");

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
                process = new Process
                {
                    StartInfo =
                    {
                        FileName = builderBinary,
                        Arguments = arguments,
                        UseShellExecute = false,
                    }
                };
            };

            after = () =>
            {
                Directory.Delete(workingDirectory, true);
            };

            context["given no procfile, executable or Web.conf exists"] = () =>
            {
                string resultFile = null;

                before = () =>
                {
                    resultFile = Path.Combine(tmpDir, "result.json");
                };

                it["Creates the result.json"] = () =>
                {
                    File.Exists(resultFile).should_be_true();
                };
            };

            context["given i have an app similar to nora"] = () =>
            {
                string resultFile = null;

                before = () =>
                {
                    Microsoft.VisualBasic.FileIO.FileSystem.CopyDirectory(Path.Combine(currentDirectory, "Builder.Tests", "app"), appDir);
                    resultFile = Path.Combine(tmpDir, "result.json");
                };

                it["Creates a droplet"] = () =>
                {
                    var fileName = Path.Combine(tmpDir, "droplet");
                    File.Exists(fileName).should_be_true();
                };

                it["Creates the result.json"] = () =>
                {
                    File.Exists(resultFile).should_be_true();
                };

                context["the json file"] = () =>
                {
                    JObject result = null;

                    act = () =>
                    {
                        result = JObject.Parse(File.ReadAllText(resultFile));
                    };

                    it["includes the start command for 'web'"] = () =>
                    {
                        var processTypes = result["process_types"].Value<JObject>();
                        var webStartCommand = processTypes["web"].Value<string>();
                        webStartCommand.should_be("tmp/lifecycle/WebAppServer.exe");
                    };

                    it["includes execution metadata"] = () =>
                    {
                        var executionMetadataJson = result["execution_metadata"].Value<string>();
                        var executionMetadata = JsonConvert.DeserializeObject<ExecutionMetadata>(executionMetadataJson);
                        executionMetadata.StartCommand.should_be("tmp/lifecycle/WebAppServer.exe");
                        executionMetadata.StartCommandArgs.should_be(new string[] { });
                    };

                    it["doesn't have any other process types"] = () =>
                    {
                        var processTypes = result["process_types"].Value<JObject>();
                        processTypes.Count.should_be(1);
                    };

                    it["includes lifecycle metadata fields"] = () =>
                    {
                        result["lifecycle_type"].Value<string>().should_be("buildpack");
                        var metadata = result["lifecycle_metadata"].Value<JObject>();
                        metadata["detected_buildpack"].Value<string>().should_be("windows");
                        metadata["buildpack_key"].Value<string>().should_be("");
                    };
                };
            };

            context["not a nora"] = () =>
            {
                string configFile = null;
                before = () =>
                {
                    configFile = Path.Combine(appDir, "Web.config");
                    File.WriteAllText(configFile, "<configuration></configuration>");
                };

                context["when there is a web.config and a user-provided-service"] = () =>
                {
                    before = () =>
                    {
                        process.StartInfo.EnvironmentVariables["VCAP_SERVICES"] =
                        JsonConvert.SerializeObject(new Services
                        {
                            UserProvided = new List<Service>
                                    {
                                        new Service
                                        {
                                            Name = "aFoo",
                                            Credentials = new Dictionary<string, string>
                                            {
                                                {"name", "foo"},
                                                {"connectionString", "bar"},
                                                {"providerName","baz"}
                                            },
                                        }
                                    }
                        });
                    };

                    it["sets a connection string"] = () =>
                    {
                        var xml = File.ReadAllText(configFile);
                        xml.should_contain("name=\"foo\"");
                        xml.should_contain("connectionString=\"bar\"");
                        xml.should_contain("providerName=\"baz\"");
                    };
                };

                context["when there is a web.config and no user-provided-services"] = () =>
                {
                    it["doesn't alter the web.config"] = () =>
                    {
                        var doc = new XmlDocument();
                        doc.Load(configFile);
                        var rootNode = doc.SelectSingleNode("//configuration");
                        rootNode.ChildNodes.Count.should_be(0);
                    };
                };
            };
        }
    }
}