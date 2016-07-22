using System.IO;
using System.Xml;
using NSpec;

namespace WebAppServer.Tests.Specs
{
    class WebConfigTest : nspec
    {
        public string GetTemporaryDirectory()
        {
            string tempDirectory = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(tempDirectory);
            return tempDirectory;
        }

        public const string DefaultRootWebConfig = @"<?xml version=""1.0"" encoding=""utf-8""?>
<configuration>
     <system.web>
        <compilation>
            <assemblies>
                <add assembly=""mscorlib"" />
            </assemblies>
        </compilation>
    </system.web>
</configuration>";

        public string RootWebConfigPath()
        {
            var file = Path.Combine(GetTemporaryDirectory(), "web.config");
            File.WriteAllText(file, DefaultRootWebConfig);
            return file;
        }

        void describe_()
        {
            describe["WebConfig"] = () =>
            {
                string baseDir = "";
                string rootWebConfigPath = "";

                before = () =>
                {
                    baseDir = GetTemporaryDirectory();
                    Directory.CreateDirectory(Path.Combine(baseDir, "config"));
                    rootWebConfigPath = RootWebConfigPath();
                };

                after = () =>
                {
                    Directory.Delete(baseDir, true);
                    File.Delete(rootWebConfigPath);
                };

                it["Creates a new configuration file"] = () =>
                        {
                            var webConfigPath = WebConfig.Create(rootWebConfigPath, baseDir);
                            webConfigPath.should_contain(baseDir);
                            var doc = new XmlDocument();
                            doc.Load(webConfigPath);
                            XmlAttribute[] attributes = new XmlAttribute[1];
                            doc.SelectSingleNode("//configuration/system.web/compilation").Attributes.CopyTo(attributes, 0);
                            attributes.Length.should_be(1);
                            attributes[0].Value.should_be(Path.Combine(baseDir, "tmp"));
                            attributes[0].Name.should_be("tempDirectory");
                        };
            };
        }
    }
}
