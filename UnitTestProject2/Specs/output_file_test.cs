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
                AppDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName()),
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
            Directory.Delete(options.AppDir, true);
            File.Delete(options.OutputDroplet);
            File.Delete(options.OutputMetadata);
        }

        void specify_output_droplet_is_a_file()
        {
            File.Exists(options.OutputDroplet).should_be_true();
        }

        void given_files_in_the_input_app_dir()
        {
            string tgzExtractedDir = Path.Combine(Path.GetTempPath(), Path.GetRandomFileName());

            before = () =>
            {
                File.WriteAllText(Path.Combine(options.AppDir, "a_file.txt"), "Some exciting text");
                File.WriteAllText(Path.Combine(options.AppDir, "another_file.txt"), "Some different text");
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

            it["is a a tar gz with app.zip inside"] = () =>
            {
                File.Exists(Path.Combine(tgzExtractedDir, "app.zip")).should_be_true();
            };

            describe["app.zip inside OutputDroplet"] = () =>
            {
                var zipPath = "";
                before = () =>
                {
                    zipPath = Path.Combine(tgzExtractedDir, "app.zip");
                };

                it["is a zipfile"] = () =>
                {
                    using (var archive = ZipFile.OpenRead(zipPath))
                    {
                        archive.Entries.Count.should_be(2);
                    }
                };

                it["contains both files"] = () =>
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

        /****
         * 
        		w("the contents of the output tgz", func() { 
118 			var files []string 
119 
 
120 			JustBeforeEach(func() { 
121 				result, err := exec.Command("tar", "-tzf", outputDroplet).Output() 
122 				Ω(err).ShouldNot(HaveOccurred()) 
123 
 
124 				files = strings.Split(string(result), "\n") 
125 			}) 
126 
 
127 			It("should contain an /app dir with the contents of the compilation", func() { 
128 				Ω(files).Should(ContainElement("./app/")) 
129 				Ω(files).Should(ContainElement("./app/app.sh")) 
130 				Ω(files).Should(ContainElement("./app/compiled")) 
131 			}) 
132 
 
133 			It("should contain an empty /tmp directory", func() { 
134 				Ω(files).Should(ContainElement("./tmp/")) 
135 				Ω(files).ShouldNot(ContainElement(MatchRegexp("\\./tmp/.+"))) 
136 			}) 
137 
 
138 			It("should contain an empty /logs directory", func() { 
139 				Ω(files).Should(ContainElement("./logs/")) 
140 				Ω(files).ShouldNot(ContainElement(MatchRegexp("\\./logs/.+"))) 
141 			}) 
142 
 
143 			It("should contain a staging_info.yml with the detected buildpack", func() { 
144 				stagingInfo, err := exec.Command("tar", "-xzf", outputDroplet, "-O", "./staging_info.yml").Output() 
145 				Ω(err).ShouldNot(HaveOccurred()) 
146 
 
147 				expectedYAML := `detected_buildpack: Always Matching 
148 start_command: the start command 
149 ` 
150 				Ω(string(stagingInfo)).Should(Equal(expectedYAML)) 
151 			}) 
152 		}) 
         ****/

    }
}
