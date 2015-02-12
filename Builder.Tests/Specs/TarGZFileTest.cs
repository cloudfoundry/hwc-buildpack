using System;
using NSpec;
using System.IO;
using System.Diagnostics;
using System.Linq;
using Newtonsoft.Json.Linq;
using Builder;
using SharpCompress.Archive;

namespace Builder.Tests
{
    class TarGZFileTest : nspec
    {
        string tgzPath, tmpDir;
        void before_each() {
            tgzPath = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            tmpDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
            Directory.CreateDirectory(tmpDir);
        }

        void after_each()
        {
            if (File.Exists(tgzPath)) File.Delete(tgzPath);
            Directory.Delete(tmpDir, true);
        }

        void describe_CreateFromDirectory()
        {
            before = () => {
                File.WriteAllText(Path.Combine(tmpDir, "a_file.txt"), "Some exciting text");
                File.WriteAllText(Path.Combine(tmpDir,"another_file.txt"), "Some different text");
                TarGZFile.CreateFromDirectory(tmpDir, tgzPath);
            };

            it["creates the tgz file"] = () =>
            {
                File.Exists(tgzPath).should_be_true();
            };

            it["puts the files inside the file"] = () =>
            {
                using (var archive = ArchiveFactory.Open(tgzPath))
                {
                    var tarFile = archive.Entries.First(entry => entry.Key == "Tar.tar");
                    using (var ms = new MemoryStream())
                    {
                        tarFile.OpenEntryStream().CopyTo(ms);
                        ms.Position = 0;
                        // now work with ms

                        var files = ArchiveFactory.Open(ms).Entries.ToArray();

                        files[0].Key.should_be("another_file.txt");
                        GetString(files[0]).should_be("Some different text");

                        files[1].Key.should_be("a_file.txt");
                        GetString(files[1]).should_be("Some exciting text");
                    }
                }
            };
        }

        static string GetString(IArchiveEntry entry)
        {
            byte[] bytes = new byte[entry.Size];
            entry.OpenEntryStream().Read(bytes, 0, bytes.Length);
            return System.Text.Encoding.Default.GetString(bytes);
        }
    }
}



