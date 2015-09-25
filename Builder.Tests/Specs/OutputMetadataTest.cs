using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using NSpec;
using System.IO;

namespace Builder.Tests
{
    class OutputMetadataTest : nspec
    {
        private void describe_()
        {
            context["given the start command"] = () =>
            {
                OutputMetadata outputMetadata = null;

                before = () =>
                {
                    outputMetadata = new OutputMetadata()
                    {
                        ExecutionMetadata = new ExecutionMetadata() {
                        StartCommand = "foo",
                        StartCommandArgs = new string[] {},
                        }
                    };
                };

                context["and some arguments"] = () =>
                {
                    before =
                        () =>
                            outputMetadata.ExecutionMetadata.StartCommandArgs = new string[]
                            {
                                "bar", "baz"
                            };

                    it["sets the detected start command"] = () =>
                    {
                        outputMetadata.ProcessTypes.Web.should_be("foo bar baz");
                    };
                };

                context["and no arguments"] = () =>
                {
                    it["sets the detected start command"] = () =>
                    {
                        outputMetadata.ProcessTypes.Web.should_be("foo");
                    };
                };
            };
        }
    }
}



