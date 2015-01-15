using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using NSpec;

namespace WebAppServer.Tests.Specs
{
    class OptionsTest : nspec
    {
        private void describe_()
        {
            describe["Parse"] = () =>
            {
                Options options = null;
                string path = null;
                string port = null;

                before = () =>
                {
                    port = "9999";
                    path = "";
                    options = new Options();
                };

                act = () =>
                {
                    options.Parse(new[] { port, path });
                };

                it["parses port as an int"] = () =>
                {
                    options.Port.should_be(9999U);
                };

                context["when a full path is passed in"] = () =>
                {
                    before = () =>
                    {
                        path = @"C:\hi\guid";
                    };

                    it["uses a full path directly"] = () =>
                    {
                        options.WebRoot.should_be(@"C:\hi\guid");
                    };
                };

                context["when a relative path is passed in "] = () =>
                {
                    before = () =>
                    {
                        path = @"app/myapp";
                    };

                    it["appends a relative directory to the current directoy"] = () =>
                    {
                        options.WebRoot.should_be(Path.Combine(Directory.GetCurrentDirectory(), "app/myapp"));
                    };
                };
            };
        }
    }
}
