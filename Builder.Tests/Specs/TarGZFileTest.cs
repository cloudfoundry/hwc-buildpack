using System;
using System.Diagnostics;
using System.Linq;
using System.Reflection;
using Builder.Tests.Properties;
using NSpec;
using System.IO;

namespace Builder.Tests
{
    class TarGZFileTest : nspec
    {
        string tgzPath, tmpDir, extractDir;
        private static string TarArchiverPath(string filename)
        {
            var uri = new Uri(Assembly.GetExecutingAssembly().CodeBase);
            return Path.Combine(Path.GetDirectoryName(uri.LocalPath), filename);
        }

        void before_all()
        {
            File.WriteAllBytes(TarArchiverPath("tar.exe"), Resources.bsdtar);
            File.WriteAllBytes(TarArchiverPath("zlib1.dll"), Resources.zlib1);
        }

        void before_each() {
            tgzPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            tmpDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(tmpDir);
            File.WriteAllText(Path.Combine(tmpDir, "a_file.txt"), "Some exciting text");
            Directory.CreateDirectory(Path.Combine(tmpDir, "a dir"));
            File.WriteAllText(Path.Combine(tmpDir, "a dir", "another_file.txt"), "Some spacey text");
        }

        void after_each()
        {
            Directory.Delete(tmpDir, true);
            if (Directory.Exists(extractDir))
            {
                Directory.Delete(extractDir, true);
            }
            if (File.Exists(tgzPath))
            {
                File.Delete(tgzPath);
            }
        }

        private void describe_CreateFromDirectory()
        {
            it["creates the tgz file"] = () =>
            {
                TarGZFile.CreateFromDirectory(tmpDir, tgzPath);
                File.Exists(tgzPath).should_be_true();
            };

            it["puts the files inside the file"] = () =>
            {
                TarGZFile.CreateFromDirectory(tmpDir, tgzPath);

                var process = new Process();
                var processStartInfo = process.StartInfo;
                processStartInfo.FileName = TarArchiverPath("tar.exe");
                processStartInfo.Arguments = "tf " + tgzPath;
                processStartInfo.RedirectStandardOutput = true;
                processStartInfo.UseShellExecute = false;
                process.Start();
                process.WaitForExit();
                var fileNames = process.StandardOutput.ReadToEnd();
                fileNames.should_contain(Path.GetFileName(tmpDir) + "/a_file.txt");
                fileNames.should_contain(Path.GetFileName(tmpDir) + "/a dir/another_file.txt");
            };

            it["can deal with Unicode in filenames"] = () =>
            {
                File.WriteAllText(Path.Combine(tmpDir, "新闻.txt"), "Chinese news");
                extractDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
                Directory.CreateDirectory(extractDir);

                TarGZFile.CreateFromDirectory(tmpDir, tgzPath);

                var process = new Process();
                var processStartInfo = process.StartInfo;
                processStartInfo.FileName = TarArchiverPath("tar.exe");
                processStartInfo.Arguments = "xf " + tgzPath + " -C " + extractDir;
                processStartInfo.UseShellExecute = false;
                process.Start();
                process.WaitForExit();
                var fileNames = Directory.EnumerateFiles(extractDir, "*", SearchOption.AllDirectories).ToList();
                fileNames.should_contain(Path.Combine(extractDir, Path.GetFileName(tmpDir)) +  @"\新闻.txt");
            };
        }
    }
}



