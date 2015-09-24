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
                OutputMetadata metaData = null;

                before = () => metaData = new OutputMetadata()
                {
                    ExecutionMetadata = new ExecutionMetadata()
                    {
                        ProcessTypes = new ProcessTypes() {
                            StartCommand = "foo",
                            StartCommandArgs = new string[] { },
                        },
                    }
                };

                context["and some arguments"] = () =>
                {
                    before =
                        () =>
                            metaData.ExecutionMetadata.ProcessTypes.StartCommandArgs = new string[]
                            {
                                "bar", "baz"
                            };

                    it["sets the detected start command"] = () =>
                    {
                        metaData.ExecutionMetadata.ProcessTypes.Web.should_be("foo bar baz");
                    };
                };

                context["and no arguments"] = () =>
                {
                    it["sets the detected start command"] = () =>
                    {
                        metaData.ExecutionMetadata.ProcessTypes.Web.should_be("foo");
                    };
                };
            };
        }
    }
}



