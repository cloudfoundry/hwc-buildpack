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
                        StartCommand = "foo",
                    }
                };

                context["and some arguments"] = () =>
                {
                    before =
                        () =>
                            metaData.ExecutionMetadata.StartCommandArgs = new string[] {"bar", "baz"};

                    it["sets the detected start command"] = () =>
                    {
                        metaData.DetectedStartCommand.Web.should_be("foo bar baz");
                    };

                    it["returns execution_metadata as json string"] = () =>
                    {
                        metaData.execution_metadata.should_be("{\"start_command\":\"foo\",\"start_command_args\":[\"bar\",\"baz\"]}");
                    };
                };

                context["and no arguments"] = () =>
                {
                    it["sets the detected start command"] = () =>
                    {
                        metaData.DetectedStartCommand.Web.should_be("foo");
                    };

                    it["returns execution_metadata as json string"] = () =>
                    {
                        metaData.execution_metadata.should_be("{\"start_command\":\"foo\",\"start_command_args\":[]}");
                    };
                };
            };
        }
    }
}



