using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using NSpec;
using System.IO;
using System.Diagnostics;
using System.Linq;
using Newtonsoft.Json.Linq;
using Builder;
using NSpec.Domain;

namespace Builder.Tests
{
    class StartCommandTest : nspec
    {
        void describe_()
        {
            OutputMetadata obj = null;
            List<string> files = null;

            context["a web.config file exists"] = () =>
            {
                before = () => files = new List<string> {@"app\foo", @"app\bar", @"app\Web.Config"};
                act = () => obj = new OutputMetadata(files);

                it["sets WebAppServer as the start command (CF)"] = () =>
                {
                    obj.DetectedStartCommand.Web.should_be("tmp/lifecycle/WebAppServer.exe .");
                };

                it["sets WebAppServer as the start command (Diego)"] = () =>
                {
                    obj.ExecutionMetadata.StartCommand.should_be("tmp/lifecycle/WebAppServer.exe");
                    obj.ExecutionMetadata.StartCommandArgs.should_be(new string[] { "." });
                };
            };

            context["an exe file exists"] = () =>
            {
                before = () => files = new List<string> { @"app\foo", @"app\bar", @"app\jane.exe" };
                act = () => obj = new OutputMetadata(files);

                it["sets the exe as the start command (CF)"] = () =>
                {
                    obj.DetectedStartCommand.Web.should_be(@"jane.exe");
                };

                it["sets the exe as the start command (Diego)"] = () =>
                {
                    obj.ExecutionMetadata.StartCommand.should_be(@"jane.exe");
                    obj.ExecutionMetadata.StartCommandArgs.should_be_empty();
                };
            };

            context["two exe files exist"] = () =>
            {
                before = () => files = new List<string> { @"app\foo", @"app\jill.exe", @"app\bar", @"app\jane.exe" };

                it["throws an exception"] = () =>
                {
                    Exception exception = null;
                    try
                    {
                        new OutputMetadata(files);
                    }
                    catch (Exception e)
                    {
                        exception = e;
                    }
                    exception.should_not_be_null();
                    exception.Message.should_be("Directory contained more than 1 executable file.");
                };
            };

            context["neither web.config nor exe file exist"] = () =>
            {
                before = () => files = new List<string> { "foo" };

                it["throws an exception"] = () =>
                {
                    Exception exception = null;
                    try
                    {
                        new OutputMetadata(files);
                    }
                    catch (Exception e)
                    {
                        exception = e;
                    }
                    exception.should_not_be_null();
                    exception.Message.should_be("No runnable application found.");
                };
            };

            context["both web.config and an exe file exist"] = () =>
            {
                before = () => files = new List<string> { "foo.exe", "Web.config" };
                act = () => obj = new OutputMetadata(files);

                it["goes with the Web.config"] = () =>
                {
                    obj.DetectedStartCommand.Web.should_be("tmp/lifecycle/WebAppServer.exe .");
                    obj.ExecutionMetadata.StartCommand.should_be("tmp/lifecycle/WebAppServer.exe");
                    obj.ExecutionMetadata.StartCommandArgs.should_be(new string[] { "." });
                };
            };
        }
    }
}



