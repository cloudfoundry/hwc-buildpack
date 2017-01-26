using System;
using System.IO;
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
                    Environment.SetEnvironmentVariable("PORT", port);
                    options = new Options();
                };

                act = () =>
                {
                    options.Parse(new string[] { });
                };

                it["parses port as an int"] = () =>
                {
                    options.Port.should_be(9999U);
                };


                it["returns the full path to the current directory"] = () =>
                {
                    options.WebRoot.should_be(Directory.GetCurrentDirectory());
                };
            };
        }
    }
}
