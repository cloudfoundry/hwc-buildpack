using System;
using System.IO;
using System.Security.Cryptography;
using System.Text;
using System.Xml.Linq;
using System.Xml.XPath;
using WebAppServer.Properties;

namespace WebAppServer
{
    internal class ConfigSettings
    {
        public string RootWebConfigPath { get; set; }
        public string AppConfigPath { get; set; }
    }

    internal class ConfigGenerator : IDisposable
    {
        private readonly string configPath;
        private readonly string webRootPath;
        private readonly string logsRootPath;
        private readonly string tempPath;

        public ConfigGenerator(string webRoot)
        {
            if (string.IsNullOrWhiteSpace(webRoot))
            {
                throw new ArgumentNullException("webRoot");
            }
            this.webRootPath = webRoot;

            var baseDirectory = AppDomain.CurrentDomain.BaseDirectory;
            configPath = Path.Combine(baseDirectory, "config");
            logsRootPath = Path.Combine(baseDirectory, "log");
            tempPath = Path.Combine(baseDirectory, "tmp");

            EnsureDirectory(configPath);
            EnsureDirectory(logsRootPath);
            EnsureDirectory(tempPath);
        }

        /// <summary>
        /// 
        /// </summary>
        /// <param name="port">The application port</param>
        /// <param name="rootWebConfigPath">The to the framework's root web.config based on runtime version</param>
        /// <param name="runtimeVersion">One of the <see><cref>Constants.RuntimeVersion</cref></see> values.</param>
        /// <param name="pipelineMode">One of the <see><cref>Constants.PipelineMode</cref></see> values.</param>
        /// <returns></returns>
        public ConfigSettings Create(uint port, string rootWebConfigPath, string runtimeVersion, string pipelineMode, string userName, string password)
        {
            var settings = new ConfigSettings
            {
                RootWebConfigPath = Environment.ExpandEnvironmentVariables(rootWebConfigPath),
                AppConfigPath = Path.Combine(configPath, "applicationHost.config")
            };

            var clrConfigPath = Path.Combine(configPath, "aspnet.config"); // TODO: might need to just -> runtime version aspnet.config file

            File.WriteAllText(clrConfigPath, Resources.aspnet);
            File.WriteAllText(settings.AppConfigPath, runtimeVersion == Constants.RuntimeVersion.VersionFourDotZero ? Resources.applicationhost : Resources.v2_0AppHost);

            // TODO: Randomize AES Session Keys??
            // TODO: Generage new machine key - use BuildMachineKeyElement()
            // TODO: /configuration/system.webServer/security/authentication/anonymousAuthentication[username] = sandbox username or just IUSR?

            var appHostConfig = XDocument.Load(settings.AppConfigPath);

            // set the application pools config files
            var appPoolName = "AppPool" + port;
            appHostConfig.XPathSelectElement(Constants.ConfigXPath.AppPools).RemoveNodes();
            appHostConfig.AddToElement(Constants.ConfigXPath.AppPools,
                BuildApplicationPool(appPoolName, runtimeVersion, pipelineMode, clrConfigPath, userName, password));

            // logging paths
            var iisLogDir = Path.Combine(logsRootPath, "IIS");
            appHostConfig.SetValue(Constants.ConfigXPath.SiteDefaults + "/logFile", "directory", iisLogDir);
            EnsureDirectory(iisLogDir);

            var traceLogDir = Path.Combine(logsRootPath, "TraceLogFiles");
            appHostConfig.SetValue(Constants.ConfigXPath.SiteDefaults + "/traceFailedRequestsLogging", "directory", traceLogDir);
            EnsureDirectory(traceLogDir);

            // add the site and settings to the app host config
            appHostConfig.AddToElement(Constants.ConfigXPath.Sites, BuildSiteElement("IronFoundrySite", webRootPath, appPoolName, port));

            appHostConfig.SetValue(Constants.ConfigXPath.Sites + "/applicationDefaults", "applicationPool", appPoolName);
            appHostConfig.SetValue(Constants.ConfigXPath.Sites + "/virtualDirectoryDefaults", "allowSubDirConfig", true);

            appHostConfig.Save(settings.AppConfigPath);
            return settings;
        }

        private void EnsureDirectory(string directory)
        {
            if (!Directory.Exists(directory))
            {
                Directory.CreateDirectory(directory);
            }
        }

        protected virtual XElement BuildApplicationPool(
            string name, string runtimeVersion,
            string pipelineMode, string configFile,
            string userName, string password)
        {
            var pool = new XElement("add");
            pool.Add(new XAttribute("name", name));
            pool.Add(new XAttribute("managedRuntimeVersion", runtimeVersion));
            pool.Add(new XAttribute("managedPipelineMode", pipelineMode));
            pool.Add(new XAttribute("CLRConfigFile", configFile));
            pool.Add(new XAttribute("autoStart", true));
            pool.Add(new XAttribute("startMode", "AlwaysRunning"));

            if (!(string.IsNullOrWhiteSpace(userName) || string.IsNullOrWhiteSpace(password)))
            {
                var processModel = new XElement("processModel");
                processModel.Add(new XAttribute("identityType", "SpecificUser"));
                processModel.Add(new XAttribute("userName", userName));
                processModel.Add(new XAttribute("password", password));
                processModel.Add(new XAttribute("manualGroupMembership", "true"));
                pool.Add(processModel);
            }

            return pool;
        }

        protected virtual XElement BuildMachineKeyElement()
        {
            Func<Int32, String> keyGen = length =>
            {
                var buff = new Byte[length];
                new RNGCryptoServiceProvider().GetBytes(buff);
                var sb = new StringBuilder(length * 2);
                foreach (byte t in buff)
                {
                    sb.AppendFormat("{0:X2}", t);
                }
                return sb.ToString();
            };

            var node = new XElement("machineKey");
            node.Add(new XAttribute("validationKey", keyGen(64)));
            node.Add(new XAttribute("decryptionKey", keyGen(32)));
            node.Add(new XAttribute("validation", "SHA1"));
            node.Add(new XAttribute("decryption", "AES"));
            return node;
        }
        
        protected virtual XElement BuildSiteElement(string name, string physicalPath, string appPoolName, uint port)
        {
            var site = new XElement("site");
            site.Add(new XAttribute("name", name));
            site.Add(new XAttribute("id", 1));
            site.Add(new XAttribute("serverAutoStart", true));
            site.Add(BuildApplicationElement("/", physicalPath, appPoolName));
            site.Add(new XElement("bindings", BuildBindingElement("http", port, String.Empty))); // "localhost" vs String.Empty for local ONLY binding
            return site;
        }

        protected virtual XElement BuildApplicationElement(string path, string physicalPath, string applicationPool)
        {
            var app = new XElement("application");
            app.Add(new XAttribute("path", path));
            app.Add(new XAttribute("applicationPool", applicationPool));
            app.Add(BuildVirtualDirectoryElement("/", physicalPath));
            return app;
        }

        protected virtual XElement BuildVirtualDirectoryElement(string path, string physicalPath)
        {
            var vDir = new XElement("virtualDirectory");
            vDir.Add(new XAttribute("path", path));
            vDir.Add(new XAttribute("physicalPath", physicalPath));
            return vDir;
        }

        protected virtual XElement BuildBindingElement(string protocol, uint port, string host)
        {
            var binding = new XElement("binding");
            binding.Add(new XAttribute("protocol", protocol));
            binding.Add(new XAttribute("bindingInformation", String.Format("*:{0}:{1}", port, host)));
            return binding;
        }

        public void Dispose()
        {
            if (Directory.Exists(configPath))
            {
                Directory.Delete(configPath, true);
            }
        }
    }
}
