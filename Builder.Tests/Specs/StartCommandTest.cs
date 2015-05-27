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
            Exception exception = null;
            act = () =>
            {
                exception = null;
                try
                {
                    obj = new OutputMetadata(files);
                }
                catch (Exception e)
                {
                    exception = e;
                }
            };

            context["a web.config file exists"] = () =>
            {
                before = () => files = new List<string> {@"app\foo", @"app\bar", @"app\Web.Config"};

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

            context["an Procfile exists"] = () =>
            {
                string filename = null;
                before = () =>
                {
                    filename = Path.Combine(Path.GetTempPath(), "Procfile");
                    files = new List<string> { filename };
                };
                after = () => File.Delete(filename);

                context["with a `web` line"] = () =>
                {
                    before = () => File.WriteAllLines(filename,
                        new string[] {"worker2: issfdsi.exe", "web: billybob.exe fred jane jim", "worker1: isudf.exe"});

                    it["sets the Procfile as the start command (CF)"] = () =>
                    {
                        obj.DetectedStartCommand.Web.should_be(@"billybob.exe fred jane jim");
                    };

                    it["sets the Procfile as the start command (Diego)"] = () =>
                    {
                        obj.ExecutionMetadata.StartCommand.should_be(@"billybob.exe");
                        obj.ExecutionMetadata.StartCommandArgs.should_be(new string[] {"fred", "jane", "jim"});
                    };

                    context["and a web.config also exist"] = () =>
                    {
                        before = () => files = new List<string> { "Web.config", filename };

                        it["goes with the Procfile"] = () =>
                        {
                            obj.DetectedStartCommand.Web.should_be(@"billybob.exe fred jane jim");
                            obj.ExecutionMetadata.StartCommand.should_be(@"billybob.exe");
                            obj.ExecutionMetadata.StartCommandArgs.should_be(new string[] { "fred", "jane", "jim" });
                        };
                    };
                };

                context["without a 'web' line"] = () =>
                {
                    before = () => File.WriteAllLines(filename,
                        new string[] {"worker2: issfdsi.exe", "worker1: isudf.exe"});

                    it["throws an exception"] = () =>
                    {
                        exception.should_not_be_null();
                        exception.Message.should_be("Procfile didn't contain a web line");
                    };
                };

                context["Procfile has letters before procfile"] = () =>
                {
                    before = () =>
                    {
                        filename = Path.Combine(Path.GetTempPath(), "iausgdProcfile");
                        File.WriteAllLines(filename, new string[] { "worker2: issfdsi.exe", "web: stuff more stuff", "worker1: isudf.exe" });
                        files = new List<string> { filename };
                    };

                    it["throws an exception"] = () =>
                    {
                        exception.should_not_be_null();
                        exception.Message.should_be("No runnable application found.");
                    };
                };
            };

            context["two exe files exist"] = () =>
            {
                before = () => files = new List<string> { @"app\foo", @"app\jill.exe", @"app\bar", @"app\jane.exe" };

                it["throws an exception"] = () =>
                {
                    exception.should_not_be_null();
                    exception.Message.should_be("Directory contained more than 1 executable file.");
                };
            };

            context["neither web.config nor exe file exist"] = () =>
            {
                before = () => files = new List<string> { "foo" };

                it["throws an exception"] = () =>
                {
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



