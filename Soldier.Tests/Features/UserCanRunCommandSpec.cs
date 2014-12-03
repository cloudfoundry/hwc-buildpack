using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using NSpec;

namespace Solider.Tests
{
    public class UserCanRunCommandSpec : nspec
    {

        private void describe_()
        {
            Process proc = null;
            String containerName = null;
            String rootDir = null;

            after = () =>
            {
                Directory.Delete(Path.Combine(rootDir, "containerizer", containerName), true);

            };

            context["As a user of soldier"] = () =>
            {
                before = () =>
                {

                    var path = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "soldier", "bin", "debug", "soldier.exe");
                    containerName = Guid.NewGuid().ToString();
                    rootDir =
                        Directory.GetDirectoryRoot(
                            Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location));
                    Directory.CreateDirectory(Path.Combine(rootDir, "containerizer", containerName));
                    proc = new Process { StartInfo = { FileName = path } };
                };

                context["When  I pass a start command in"] = () =>
                {
                    String dirName = null;

                    context["without arguments"] = () =>
                    {
                        before = () =>
                        {
                            dirName = Guid.NewGuid().ToString();
                            var touchPath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "generateFile", "bin", "debug", "generateFile.exe");
                            proc.StartInfo.Arguments += touchPath;
                            proc.StartInfo.Arguments += " " + Path.Combine(rootDir, "containerizer", containerName);
                        };

                        it["runs my command in the given directory"] = () =>
                        {
                            proc.Start();
                            proc.WaitForExit();
                            var path = Path.Combine(rootDir, "containerizer", containerName, "dummyFile.txt");
                            File.Exists(path).should_be_true();
                        };
                    };

                    context["with arguments"] = () =>
                    {
                        before = () =>
                        {
                            dirName = Guid.NewGuid().ToString();
                            var mkDirPath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "mkdir", "bin", "debug", "mkdir.exe");
                            proc.StartInfo.Arguments += "\"" + mkDirPath + " " + dirName + "\"";
                            proc.StartInfo.Arguments += " " + Path.Combine(rootDir, "containerizer", containerName);
                        };

                        it["runs my command in the given directory"] = () =>
                        {
                            proc.Start();
                            proc.WaitForExit();
                            var path = Path.Combine(rootDir, "containerizer", containerName, dirName);
                            Directory.Exists(path).should_be_true();
                        };
                    };
                };
            };
        }
    }
}
/*
As a User of Solider,
When I pass a start command in
And I pass a directory in
Then my command should be ran in the given directory
Labels
*/