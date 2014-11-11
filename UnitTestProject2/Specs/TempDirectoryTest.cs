using System;
using NSpec;
using System.IO;
using System.Diagnostics;
using System.Linq;
using Newtonsoft.Json.Linq;
using Tailor;

namespace Tailor.Tests
{
    class TempDirectoryTest : nspec
    {
        void it_creates_and_deletes_a_temporary_directory_accessible_as_path_string()
        {
            string path;
            using (var tmp = new TempDirectory())
            {
                path = tmp.PathString();
                Directory.Exists(path).should_be_true();
            }
            Directory.Exists(path).should_be_false();
        }

        void describe_combine()
        {
            it["adds the string to the path"] = () =>
            {
                using (var tmp = new TempDirectory())
                {
                    var path = tmp.Combine("fred.txt");
                    path.should_be(tmp.PathString() + @"\fred.txt");
                }
            };
        }
    }
}



