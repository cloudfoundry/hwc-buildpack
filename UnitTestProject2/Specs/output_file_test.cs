using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using NSpec;
using System.IO;
using System.IO.Compression;
using SharpCompress.Reader;
using SharpCompress.Common;

namespace TailorTest
{
    class the_contents_of_the_output_droplet : nspec
    {
        Tailor.Options options;

        void before_each()
        {
            options = new Tailor.Options
            {      
                AppDir = "/app",
                OutputDroplet = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".tgz"),
                OutputMetadata = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName() + ".json")
            };
            Directory.CreateDirectory(options.AppDir);
        }

        void act_each()
        {
            Tailor.Program.Run(options);
        }

        void after_each()
        {
            File.Delete(options.OutputDroplet);
            File.Delete(options.OutputMetadata);
        }

        void given_files_in_the_input_app_dir()
        {
            string tgzExtractedDir = null;

            before = () =>
            {
                File.WriteAllText(Path.Combine(options.AppDir, "a_file.txt"), "Some exciting text");
                File.WriteAllText(Path.Combine(options.AppDir, "another_file.txt"), "Some different text");
                tgzExtractedDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());
                Directory.CreateDirectory(tgzExtractedDir);
            };

            act = () =>
            {
                using (Stream stream = File.OpenRead(options.OutputDroplet))
                {
                    var reader = ReaderFactory.Open(stream);
                    while (reader.MoveToNextEntry())
                    {
                        if (!reader.Entry.IsDirectory)
                        {
                            Console.WriteLine(reader.Entry.Key);
                            reader.WriteEntryToDirectory(tgzExtractedDir, ExtractOptions.ExtractFullPath | ExtractOptions.Overwrite);
                        }
                    }
                }
            };

            after = () =>
            {
                Directory.Delete(tgzExtractedDir, true);
            };

            xit["is a a tar gz with app.zip inside"] = () =>
            {
                File.Exists(Path.Combine(tgzExtractedDir, "app.zip")).should_be_true();
            };

            describe["app.zip inside OutputDroplet"] = () => 
            {
                string zipPath = null;
                before = () =>
                {
                    zipPath = Path.Combine(tgzExtractedDir, "app.zip");
                };

                xit["is a zipfile"] = () =>
                {
                    using (var archive = ZipFile.OpenRead(zipPath))
                    {
                        archive.Entries.Count.should_be(2);
                    }
                };

                xit["contains both files"] = () =>
                {
                    using (var archive = ZipFile.OpenRead(zipPath))
                    {
                        var expected = new string[2] { "another_file.txt", "a_file.txt" };
                        var actual = archive.Entries.Select(entry => entry.Name).ToArray();
                        actual.should_be(expected);
                    }
                };
            };
        }
    }
}
