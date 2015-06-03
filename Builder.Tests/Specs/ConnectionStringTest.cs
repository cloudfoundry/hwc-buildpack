using Builder.Models;
using NSpec;
using System;
using System.Collections.Generic;
using System.Xml;

namespace Builder.Tests
{
    class ConnectionStringTest : nspec
    {
        void describe_()
        {
            String webConfig = null;
            XmlDocument doc = null;
            Services services = null;
            Exception ex = null;

            act = () =>
            {
                doc = new XmlDocument();
                doc.LoadXml(webConfig);
                try
                {
                    Program.SetConnectionStrings(doc, services);
                }
                catch (Exception e)
                {
                    ex = e;
                }
            };

            context["with more than one service"] = () =>
            {
                before = () =>
                {
                    services = new Services()
                    {
                        UserProvided = new List<Service>()
                        {
                            new Service()
                            {
                                Name = "foo",
                                Credentials = new Dictionary<string, string>()
                                {
                                    {"name", "foo"},
                                    {"connectionString", "bar"},
                                    {"providerName", "baz"}
                                }
                            },
                            new Service()
                            {
                                Name = "invalidService",
                                Credentials = new Dictionary<string, string>()
                                {
                                    {"username", "username"},
                                    {"host", "localhost:3306"},
                                    {"database", "db_test"}
                                }
                            },
                            new Service()
                            {
                                Name = "anotherService",
                                Credentials = new Dictionary<string, string>()
                                {
                                    {"name", "foo2"},
                                    {"connectionString", "bar2"},
                                    {"providerName", "baz2"}
                                }
                            }
                        }
                    };
                    webConfig = "<configuration/>";
                };

                it["adds valid services only"] = () =>
                {
                    var elem = doc.SelectSingleNode("//configuration/connectionStrings");
                    ex.should_be_null();
                    elem.ChildNodes.Count.should_be(2);

                    elem.ChildNodes[0].Name.should_be("add");
                    elem.ChildNodes[0].Attributes["name"].Value.should_be("foo");
                    elem.ChildNodes[0].Attributes["connectionString"].Value.should_be("bar");
                    elem.ChildNodes[0].Attributes["providerName"].Value.should_be("baz");

                    elem.ChildNodes[1].Name.should_be("add");
                    elem.ChildNodes[1].Attributes["name"].Value.should_be("foo2");
                    elem.ChildNodes[1].Attributes["connectionString"].Value.should_be("bar2");
                    elem.ChildNodes[1].Attributes["providerName"].Value.should_be("baz2");
                };
            };

           context["with no services"] = () =>
            {
                before = () =>
                {
                    services = new Services();
                    webConfig = "<configuration><connectionStrings><add/><clear/></connectionStrings></configuration>";
                };

                it["leaves the connectionsStrings element"] = () =>
                {
                    var elem = doc.SelectSingleNode("//configuration/connectionStrings");
                    ex.should_be_null();
                    elem.ChildNodes.Count.should_be(2);
                    elem.ChildNodes[0].Name.should_be("add");
                    elem.ChildNodes[1].Name.should_be("clear");
                };
            };

            context["with one service"] = () =>
            {
                before = () =>
                {
                    services = new Services
                    {
                        UserProvided = new List<Service>()
                        {
                            new Service()
                            {
                                Name = "foo",
                                Credentials = new Dictionary<string, string>()
                                {
                                    {"name", "foo"},
                                    {"connectionString", "bar"},
                                    {"providerName", "baz"}
                                }
                            }
                        }
                    };
                };

                context["when the connectionStrings element does not exist"] = () =>
                {
                    before = () =>
                    {
                        webConfig = "<configuration />";
                    };

                    it["creates the connectionStrings node with the connection as its only child"] = () =>
                    {
                        var connectionString = doc.SelectSingleNode("//configuration/connectionStrings/add");
                        connectionString.should_not_be_null();
                        connectionString.Attributes["name"].Value.should_be("foo");
                        connectionString.Attributes["connectionString"].Value.should_be("bar");
                        connectionString.Attributes["providerName"].Value.should_be("baz");
                    };
                };

                context["when the connectionStrings element exists with some children"] = () =>
                {
                    before = () =>
                    {
                        webConfig = "<configuration><connectionStrings><add /><clear /></connectionStrings></configuration>";
                    };

                    it["deletes the children and adds one add child"] = () =>
                    {
                        doc.SelectSingleNode("//configuration/connectionStrings").ChildNodes.Count.should_be(1);
                    };
                };

                context["when an invalid document is passed"] = () =>
                {
                    before = () =>
                    {
                        webConfig = "<foo/>";
                    };

                    it["throws an exception"] = () =>
                    {
                        ex.should_not_be_null();
                        ex.Message.should_contain("invalid webconfig");
                    };
                };

                context["when there is an empty connectionStrings element"] = () =>
                {
                    before = () =>
                    {
                        webConfig = "<configuration><connectionStrings/></configuration>";
                    };

                    it["writes the connection string to the XmlDocument"] = () =>
                    {
                        var connectionString = doc.SelectSingleNode("//configuration/connectionStrings/add");
                        connectionString.should_not_be_null();
                        connectionString.Attributes["name"].Value.should_be("foo");
                        connectionString.Attributes["connectionString"].Value.should_be("bar");
                        connectionString.Attributes["providerName"].Value.should_be("baz");
                    };
                };
            };
        }
    }
}



