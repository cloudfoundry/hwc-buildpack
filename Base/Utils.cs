using System.Collections.Generic;
using System.IO;
using System.Linq;

namespace Base
{
    public class Utils
    {
        public static bool HasWebConfig(IList<string> files)
        {
            return files.Any(x => Path.GetFileName(x).ToLower() == "web.config");
        }

        public static string GetStartCommand(IList<string> files)
        {
            string startCommand = null;
            if (HasWebConfig(files))
            {
                startCommand = Path.Combine(".cloudfoundry", "WebAppServer", "WebAppServer.exe");
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
