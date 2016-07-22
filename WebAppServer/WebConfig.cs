using System.IO;
using System.Xml;

namespace WebAppServer
{
    public class WebConfig
    {
        public static string Create(string rootWebConfigPath, string baseDir)
        {
            var webConfig = Path.Combine(baseDir, "config", "web.config");
            File.Copy(rootWebConfigPath, webConfig, true);

            var doc = new XmlDocument();
            doc.Load(webConfig);
            var tempDirAttr = doc.CreateAttribute("tempDirectory");
            tempDirAttr.Value = Path.Combine(baseDir, "tmp");
            doc.SelectSingleNode("//configuration/system.web/compilation").Attributes.Append(tempDirAttr);
            doc.Save(webConfig);

            return webConfig;
        }
    }
}
