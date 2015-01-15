using Moq;
using NSpec;
using System.Threading;

namespace WebAppServer.Tests.Specs
{
    class ProgramTest : nspec
    {
        void describe_()
        {
            describe["RunWebServer"] = () =>
            {
                Mock<IWebServer> webServer = null;
                Mock<Logger> log = null;

                before = () =>
                {
                    log = new Mock<Logger>();
                };

                context["When WebServer does not have access to the port"] = () =>
                {
                    before = () =>
                    {
                        webServer = new Mock<IWebServer>();
                        webServer.Setup(x => x.Start())
                            .Throws(new WebCorePortException());
                    };

                    it["Gives the user a helpful error message"] = () =>
                    {
                        Program.RunWebServer(webServer.Object, log.Object, new ManualResetEvent(true));
                        log.Verify(
                            x => x.Error(
                                It.Is<string>(message => message.StartsWith("Please allow the user to access the port. eg. 'netsh http add urlacl url=http://*:9999/ user=")), 
                                It.IsAny<string>()
                            )
                        );
                    };
                };
            };
        }
    }
}
