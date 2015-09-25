using System.Threading;
using Builder.Models;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Xml;

namespace Builder
{
    public class Program
    {
        public static ExecutionMetadata GenerateExecutionMetadata(IList<string> files)
        {
            var executionMetadata = new ExecutionMetadata();
            var procfiles = files.Where(x => Path.GetFileName(x).ToLower() == "procfile").ToList();
            var executables = files.Where(x => x.EndsWith(".exe")).ToList();
            if (procfiles.Any())
            {
                var file = File.ReadAllLines(procfiles.First());
                var webline = file.Where(x => x.StartsWith("web:"));
                if (webline.Any())
                {
                    var contents = webline.First().Substring(4).Trim().Split(new[] { ' ' });
                    executionMetadata.StartCommand = contents[0];
                    executionMetadata.StartCommandArgs = contents.Skip(1).ToArray();
                }
                else
                {
                    throw new Exception("Procfile didn't contain a web line");
                }
            }
            else if (files.Any(x => Path.GetFileName(x).ToLower() == "web.config"))
            {
                executionMetadata.StartCommand = "tmp/lifecycle/WebAppServer.exe";
                executionMetadata.StartCommandArgs = new[] { "." };
            }
            else if (executables.Any())
            {
                if (executables.Count() > 1)
                    throw new Exception("Directory contained more than 1 executable file.");
                executionMetadata.StartCommand = Path.GetFileName(executables.First());
                executionMetadata.StartCommandArgs = new string[] { };
            }
            else
            {
                Console.Error.WriteLine("No start command detected");
            }

            return executionMetadata;
        }

        static void Main(string[] args)
        {
            SanitizeArgs(args);
            var options = new Options();
            if (!CommandLine.Parser.Default.ParseArguments(args, options))
            {
                Environment.Exit(1);
            }

            Run(options);
        }

        private static void Run(Options options)
        {
            var appPath = Directory.GetCurrentDirectory() + options.BuildDir;
            var files = Directory.EnumerateFiles(appPath).ToList();

            // Set connection string if there's a web config
            SetConnectionStrings(files);

            // Result.JSON
            var obj = GenerateOutputMetadata(files);
            File.WriteAllText(Directory.GetCurrentDirectory() + options.OutputMetadata, JsonConvert.SerializeObject(obj));

            // create droplet
            var outputDropletPath = Directory.GetCurrentDirectory() + options.OutputDroplet;
            TarGZFile.CreateFromDirectory(appPath, outputDropletPath);
        }

        private static void SetConnectionStrings(IList<string> files)
        {
            var webConfig = files.FirstOrDefault(x => Path.GetFileName(x).ToLower() == "web.config");
            var vcapServices = Environment.GetEnvironmentVariable("VCAP_SERVICES");
            if (webConfig == null || vcapServices == null)
            {
                return;
            }
            var services = JsonConvert.DeserializeObject<Services>(vcapServices);
            var doc = new XmlDocument();
            doc.Load(webConfig);
            SetConnectionStrings(doc, services);
            doc.Save(webConfig);
        }

        public static void SetConnectionStrings(XmlDocument doc, Services services)
        {
            if (services.UserProvided.Count == 0)
            {
                return;
            }

            var xmlNode = doc.SelectSingleNode("//configuration/connectionStrings");
            if (xmlNode == null)
            {
                xmlNode = doc.SelectSingleNode("//configuration");
                if (xmlNode == null)
                {
                    throw new Exception("invalid webconfig");
                }
                var connectionStrings = doc.CreateElement("connectionStrings", null);
                xmlNode.AppendChild(connectionStrings);
                xmlNode = connectionStrings;
            }
            xmlNode.RemoveAll();

            foreach (var service in services.UserProvided)
            {
                var addNode = doc.CreateElement("add", null);
                string name;
                service.Credentials.TryGetValue("name", out name);
                string connectionString;
                service.Credentials.TryGetValue("connectionString", out connectionString);
                string providerName;
                service.Credentials.TryGetValue("providerName", out providerName);
                if (name == null || connectionString == null || providerName == null)
                {
                    continue;
                }
                AddAttribute(addNode, "name", name);
                AddAttribute(addNode, "connectionString", connectionString);
                AddAttribute(addNode, "providerName", providerName);
                xmlNode.AppendChild(addNode);
            }
        }

        private static void AddAttribute(XmlElement elem, string name, string value)
        {
            var doc = elem.OwnerDocument;
            var attr = doc.CreateAttribute(name);
            attr.Value = value;
            elem.Attributes.Append(attr);
        }

        private static OutputMetadata GenerateOutputMetadata(IList<string> files)
        {
            var executionMetadata = GenerateExecutionMetadata(files);
            return new OutputMetadata()
            {
                ExecutionMetadata=  executionMetadata,
            };
        }

        private static void SanitizeArgs(string[] args)
        {
            for (int i = 0; i < args.Length; i++)
            {
                if (args[i].StartsWith("-") && !args[i].StartsWith("--"))
                {
                    args[i] = "-" + args[i];
                }
            }
        }
    }
}
