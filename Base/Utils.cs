using Base.Models;
using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Xml;

namespace Base
{
    public class Utils
    {
        public static bool HasWebConfig(IList<string> files)
        {
            return files.Any(x => Path.GetFileName(x).ToLower() == "web.config");
        }

        public static IEnumerable<string> ExeFiles(IList<string> files)
        {
            return  files.Where(x => x.EndsWith(".exe", true, CultureInfo.InvariantCulture));
        }

        public static string GetStartCommand(IList<string> files)
        {
            string startCommand = null;
            var executables = ExeFiles(files).ToList();
            if (HasWebConfig(files))
            {
                startCommand = @"WebAppServer\WebAppServer.exe";
            }
            else if (executables.Any())
            {
                if (executables.Count() > 1)
                {
                    return null;
                    // throw new Exception("Directory contains more than 1 executable file.");
                }

                startCommand = Path.GetFileName(executables.First());
            }

            return startCommand;
        }

        public  static string GenerateReleaseInfo(string buildPath)
        {
            var files = Directory.EnumerateFiles(buildPath).ToList();

            using (var releaseInfoYaml = new StringWriter())
            {
                releaseInfoYaml.WriteLine("---");
                var startCmd = GetStartCommand(files);
                if (startCmd == null)
                {
                    releaseInfoYaml.WriteLine("default_process_types: {}");
                }
                else
                {
                    releaseInfoYaml.WriteLine("default_process_types:");
                    releaseInfoYaml.WriteLine("  web: \"{0}\"", startCmd.Replace("\\", "\\\\"));
                }

                return releaseInfoYaml.ToString();
            }
        }

        public static void SetConnectionStrings(IList<string> files)
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

        static public void CopyDirectory(string sourcePath, string destiationPath)
        {
            foreach (string dirPath in Directory.GetDirectories(sourcePath, "*", SearchOption.AllDirectories))
            {
                Directory.CreateDirectory(dirPath.Replace(sourcePath, destiationPath));
            }

            foreach (string newPath in Directory.GetFiles(sourcePath, "*.*", SearchOption.AllDirectories))
            {
                File.Copy(newPath, newPath.Replace(sourcePath, destiationPath), true);
            }
        }
    }
}
